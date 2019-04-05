who-to-blame
============

`git blame` coverage. A project for fun to play around with go modules, channels, and goroutines.

### Usage

```
who-to-blame --help
usage: who-to-blame [<flags>] <repo>

Flags:
  --help           Show context-sensitive help (also try --help-long and --help-man).
  --commit=COMMIT  optional commit hash to use
  --version        Show application version.

Args:
  <repo>  path to git repository

```

### Example

```
$ git clone git@github.com:pallets/click.git
$ who-to-blame click
error: click/_textwrap.py: contents and commits have different length
error: docs/index.rst: contents and commits have different length
+-----------------------------------------------+---------------+----------+
| AUTHOR                                        | LINES COVERED | PERCENT  |
+-----------------------------------------------+---------------+----------+
| armin.ronacher@active-4.com                   |         13080 |   70.41% |
| davidism@gmail.com                            |           833 |    4.48% |
| nwiles@google.com                             |           739 |    3.98% |
| me@kevinyap.ca                                |           639 |    3.44% |
| markus@unterwaditzer.net                      |           393 |    2.12% |
| jacrotts@cisco.com                            |           332 |    1.79% |
| jbischko@barracuda.com                        |           180 |    0.97% |
| sirosen@uchicago.edu                          |           156 |    0.84% |
| sirosen@globus.org                            |           105 |    0.57% |
| dharris@truthinitiative.org                   |           104 |    0.56% |
| SWu@users.noreply.github.com                  |           100 |    0.54% |
| thecubic@thecubic.net                         |            99 |    0.53% |
| santiago.basulto@gmail.com                    |            71 |    0.38% |
| s.neubauer@syseleven.de                       |            69 |    0.37% |
| ryanjsiemens@gmail.com                        |            66 |    0.36% |
| kissgyorgy@me.com                             |            65 |    0.35% |
| rpooley@uber.com                              |            64 |    0.34% |
| simon@simonjagoe.com                          |            58 |    0.31% |
| skozlovf@gmail.com                            |            57 |    0.31% |
| dawran6@gmail.com                             |            57 |    0.31% |
| markbt@efaref.net                             |            55 |    0.30% |
| jh@web.de                                     |            50 |    0.27% |
| slafs@op.pl                                   |            48 |    0.26% |
| dharris@gmail.com                             |            46 |    0.25% |
| julenx@gmail.com                              |            45 |    0.24% |
| joe@quantopian.com                            |            45 |    0.24% |
| mail@torf.cc                                  |            44 |    0.24% |
| segev208@gmail.com                            |            43 |    0.23% |
| ja@matejcik.cz                                |            43 |    0.23% |
| mabidi6@gmail.com                             |            43 |    0.23% |
| dsully@linkedin.com                           |            36 |    0.19% |
| django@bittner.it                             |            35 |    0.19% |
| claudio.bandera@kit.edu                       |            35 |    0.19% |
| kporangehat@gmail.com                         |            35 |    0.19% |
| ckin2001+github@gmail.com                     |            31 |    0.17% |
| eric.frederich@gmail.com                      |            30 |    0.16% |
| opensource@ronnypfannschmidt.de               |            29 |    0.16% |
| frankie@robertson.name                        |            28 |    0.15% |
| rovanion.luckey@gmail.com                     |            24 |    0.13% |
| aspyatkin@users.noreply.github.com            |            23 |    0.12% |
| rpgraham84@gmail.com                          |            23 |    0.12% |
| p.f.moore@gmail.com                           |            21 |    0.11% |
| josiahwhite@gmail.com                         |            21 |    0.11% |
| fabio.menegazzo@axado.com.br                  |            21 |    0.11% |
| patrick@mezard.eu                             |            20 |    0.11% |
| pietro@pietroalbini.io                        |            19 |    0.10% |
| i@julin.me                                    |            18 |    0.10% |
| frho@mail.de                                  |            18 |    0.10% |
| wpifer3@gatech.edu                            |            17 |    0.09% |
| ulo@ulo.pe                                    |            15 |    0.08% |
| russ@rw.id.au                                 |            14 |    0.08% |
| peter@knewton.com                             |            14 |    0.08% |
| georgeyk.dev@gmail.com                        |            13 |    0.07% |
| juice500ml@gmail.com                          |            12 |    0.06% |
| vmalloc@gmail.com                             |            12 |    0.06% |
| quodlibetor@gmail.com                         |            11 |    0.06% |
| remy.greinhofer@gmail.com                     |            11 |    0.06% |
| ath.daglis@gmail.com                          |            11 |    0.06% |
| jon.dufresne@gmail.com                        |            11 |    0.06% |
| quobit@users.noreply.github.com               |            10 |    0.05% |
| oleksandr.shulgin@zalando.de                  |            10 |    0.05% |
| her0e1c1@gmail.com                            |             9 |    0.05% |
| tonjayin@gmail.com                            |             9 |    0.05% |
| kkcocogogo@gmail.com                          |             9 |    0.05% |
| rhgrant10@gmail.com                           |             9 |    0.05% |
| naoya.tozuka@gmail.com                        |             8 |    0.04% |
| jhwgh1968@noreply.github.com                  |             7 |    0.04% |
| cocodrips@gmail.com                           |             7 |    0.04% |
| nicholas.chammas@gmail.com                    |             6 |    0.03% |
| pombredanne@gmail.com                         |             6 |    0.03% |
| obonilla@linkedin.com                         |             6 |    0.03% |
| stefan@sofa-rockers.org                       |             6 |    0.03% |
| nwiles@nwiles-macbookpro.roam.corp.google.com |             5 |    0.03% |
| bast@users.noreply.github.com                 |             5 |    0.03% |
| lvt.stx@gmail.com                             |             5 |    0.03% |
| d_hanchar@wargaming.net                       |             5 |    0.03% |
| swillison@gmail.com                           |             5 |    0.03% |
| zyegfryed@gmail.com                           |             4 |    0.02% |
| kaushik.ghose@sbgenomics.com                  |             4 |    0.02% |
| bbethke@airware.com                           |             4 |    0.02% |
| matthias@urlichs.de                           |             4 |    0.02% |
| jon@jonafato.com                              |             3 |    0.02% |
| me@cooperlees.com                             |             3 |    0.02% |
| zbir@zacbir.net                               |             3 |    0.02% |
| panktist@gmail.com                            |             3 |    0.02% |
| bionikspoon@gmail.com                         |             3 |    0.02% |
| ant.bertin@gmail.com                          |             3 |    0.02% |
| gordon.fierce@gmail.com                       |             3 |    0.02% |
| lasse@bigum.org                               |             3 |    0.02% |
| github@thequod.de                             |             3 |    0.02% |
| carriere.denis@gmail.com                      |             3 |    0.02% |
| lamby@debian.org                              |             3 |    0.02% |
| renne@rennerocha.com                          |             3 |    0.02% |
| bblackifyme@gmail.com                         |             3 |    0.02% |
| vishalsodani@rediffmail.com                   |             3 |    0.02% |
| ax2@princeton.edu                             |             3 |    0.02% |
| me@lepture.com                                |             2 |    0.01% |
| yuval.langer@gmail.com                        |             2 |    0.01% |
| 39267185+Atausch@users.noreply.github.com     |             2 |    0.01% |
| dirn@dirnonline.com                           |             2 |    0.01% |
| grinch@grinchcentral.com                      |             2 |    0.01% |
| g@rre.tt                                      |             2 |    0.01% |
| apetrovic@atlassian.com                       |             1 |    0.01% |
| jamestwebber@users.noreply.github.com         |             1 |    0.01% |
| mrtact@gmail.com                              |             1 |    0.01% |
| fgimiansoftware@gmail.com                     |             1 |    0.01% |
| jpzapanta22@gmail.com                         |             1 |    0.01% |
| rocky@fusionbox.com                           |             1 |    0.01% |
| ryan@ryanhiebert.com                          |             1 |    0.01% |
| amjith.r@gmail.com                            |             1 |    0.01% |
| ahmed.ettanany@gmail.com                      |             1 |    0.01% |
| holmars@gmail.com                             |             1 |    0.01% |
| scottabelden+github.sb@gmail.com              |             1 |    0.01% |
| eli.boyarski@gmail.com                        |             1 |    0.01% |
| tedmiston@gmail.com                           |             1 |    0.01% |
| minaasha@amazon.com                           |             1 |    0.01% |
| public@enkore.de                              |             1 |    0.01% |
| contact@ionelmc.ro                            |             1 |    0.01% |
| vincent@vincent-jacques.net                   |             1 |    0.01% |
| flo@chaoflow.net                              |             1 |    0.01% |
| lord63.j@gmail.com                            |             1 |    0.01% |
| rafael@rafaelmartins.eng.br                   |             1 |    0.01% |
| mholtz@protonmail.com                         |             1 |    0.01% |
| moul@moul.re                                  |             1 |    0.01% |
| withlihui@gmail.com                           |             1 |    0.01% |
| anonfunc@gmail.com                            |             1 |    0.01% |
| kevin@skytruth.org                            |             1 |    0.01% |
| sergey.kozub@p9ft.com                         |             1 |    0.01% |
| kislyuk@gmail.com                             |             1 |    0.01% |
| jmorgancampbell@gmail.com                     |             1 |    0.01% |
| brandon.milton21@gmail.com                    |             1 |    0.01% |
| sylvain.fankhauser@gmail.com                  |             1 |    0.01% |
| jeff@jeffwidman.com                           |             1 |    0.01% |
| luw2007@gmail.com                             |             1 |    0.01% |
| Brian.M.Bruggeman@gmail.com                   |             1 |    0.01% |
| accounts@sheckel.net                          |             1 |    0.01% |
| carolcode@willingconsulting.com               |             1 |    0.01% |
| wangwenpei@nextoa.com                         |             1 |    0.01% |
| me@mause.me                                   |             1 |    0.01% |
| ralt@users.noreply.github.com                 |             1 |    0.01% |
| james.martin7@gmail.com                       |             1 |    0.01% |
| razzi53@gmail.com                             |             1 |    0.01% |
| christophe31@gmail.com                        |             1 |    0.01% |
+-----------------------------------------------+---------------+----------+
| TOTAL                                         |         18578 |  100.00% |
+-----------------------------------------------+---------------+----------+
```
