package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
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

	blameInfo, err := getBlameInfo(repo, commit)
	displayBlameInfo(blameInfo)
}

func getLatestCommit(repo *git.Repository) (*object.Commit, error) {
	head, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return repo.CommitObject(head.Hash())
}

func getBlameInfo(repo *git.Repository, commit *object.Commit) (blameInfo BlameInfo, err error) {
	to, err := commit.Tree()
	if err != nil {
		return BlameInfo{}, err
	}

	results := make(chan BlameInfo)
	numFiles := 0
	_ = to.Files().ForEach(func(f *object.File) error {
		go processFile(f, commit, results)
		numFiles++
		return nil
	})

	for i := 0; i < numFiles; i++ {
		result := <- results
		blameInfo = mergeBlameInfo(blameInfo, result)
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
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Author", "Lines Covered"})
	total := 0
	for author := range blameInfo {
		t.AppendRow(table.Row{
			author, blameInfo[author],
		})
		total += blameInfo[author]
	}
	t.AppendFooter(table.Row{"Total", total})
	t.Render()
}

// TODO "contents and commits have different length" in output
func processFile(f *object.File, commit *object.Commit, results chan<- BlameInfo) {
	blameInfo := make(map[string]int)

	blameResult, err := git.Blame(commit, f.Name)
	if err != nil {
		// TODO collect errors
		fmt.Println(err)
	}

	if blameResult == nil {
		results <- blameInfo
		return
	}

	for _, line := range blameResult.Lines {
		if _, ok := blameInfo[line.Author]; ok {
			blameInfo[line.Author]++
		} else {
			blameInfo[line.Author] = 1
		}
	}

	results <- blameInfo
}

type BlameInfo map[string]int
