package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
	"sort"
)

var (
	commitHash = kingpin.Flag("commit", "optional commit hash to use").String()
	repoPath   = kingpin.Arg("repo", "path to git repository").Required().String()
)

func main() {
	kingpin.Parse()

	repo, err := git.PlainOpen(*repoPath)
	if err != nil {
		fmt.Printf("invalid repo: %v\n", err)
		os.Exit(1)
	}

	var commit *object.Commit
	if *commitHash == "" {
		commit, err = getLatestCommit(repo)
		if err != nil {
			fmt.Printf("could not get latest commit: %v\n", err)
			os.Exit(1)
		}
	} else {
		hash := plumbing.NewHash(*commitHash)
		commit, err = repo.CommitObject(hash)
		if err != nil {
			fmt.Printf("could not get commit %s: %v\n", *commitHash, err)
			os.Exit(1)
		}
	}

	blameInfo, err := getBlameInfo(commit)
	displayBlameInfo(blameInfo)
}

func getLatestCommit(repo *git.Repository) (*object.Commit, error) {
	head, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return repo.CommitObject(head.Hash())
}

func getBlameInfo(commit *object.Commit) (blameInfo BlameInfo, err error) {
	to, err := commit.Tree()
	if err != nil {
		return blameInfo, err
	}

	results := make(chan ResultWrapper)
	numFiles := 0
	_ = to.Files().ForEach(func(f *object.File) error {
		go processFile(f, commit, results)
		numFiles++
		return nil
	})

	for i := 0; i < numFiles; i++ {
		result := <- results
		if result.err != nil {
			fmt.Printf("error: %v\n", result.err)
			continue
		}
		blameInfo = mergeBlameInfo(blameInfo, result.blameInfo)
	}

	return blameInfo, err
}

func mergeBlameInfo(bi1, bi2 BlameInfo) BlameInfo {
	result := BlameInfo{}

	for key := range bi1 {
		result[key] = bi1[key]
	}

	for key := range bi2 {
		if _, ok := result[key]; ok {
			result[key] += bi2[key]
		} else {
			result[key] = bi2[key]
		}
	}

	return result
}

func displayBlameInfo(blameInfo BlameInfo) {
	totalLines := 0
	for author := range blameInfo {
		totalLines += blameInfo[author]
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Author", "Lines Covered", "Percent"})
	var rows []table.Row
	for author := range blameInfo {
		percentage := fmt.Sprintf("%7.2f%%", 100 * (float64(blameInfo[author]) / float64(totalLines)))
		rows = append(rows, table.Row{
			author, blameInfo[author], percentage,
		})
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i][1].(int) > rows[j][1].(int)
	})

	t.AppendRows(rows)
	t.AppendFooter(table.Row{"Total", totalLines, fmt.Sprintf("%7.2f%%", 100 * (float64(totalLines) / float64(totalLines)))})
	t.Render()
}

func processFile(f *object.File, commit *object.Commit, results chan<- ResultWrapper) {
	blameInfo := make(map[string]int)

	blameResult, err := git.Blame(commit, f.Name)
	if err != nil {
		results <- ResultWrapper{err: fmt.Errorf("%s: %v", f.Name, err)}
		return
	}

	if blameResult == nil {
		results <- ResultWrapper{}
		return
	}

	for _, line := range blameResult.Lines {
		if _, ok := blameInfo[line.Author]; ok {
			blameInfo[line.Author]++
		} else {
			blameInfo[line.Author] = 1
		}
	}

	results <- ResultWrapper{blameInfo: blameInfo}
}

type BlameInfo map[string]int

type ResultWrapper struct {
	blameInfo BlameInfo
	err error
}