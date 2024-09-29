# alignmentGO 
- a golang package for phylogenomics
- **This single application bring together all the functions i coded previously in a single package. By tomorrow, i will compile the final package for the release**
- All single package will be still available from my github repository and this complete package will be available from GO packages. 
  --https://github.com/codecreatede/golang-reference-alignment-estimate \
  --https://github.com/codecreatede/golang-alignment-block-genotyper-eDNA \
  --https://github.com/codecreatede/goroutines-phylogenomics-filter \
  --https://github.com/codecreatede/golang-alignment-block-getter \
  --https://github.com/codecreatede/go-alignment-estimate \
  --https://github.com/codecreatede/go-phylogenomics-tab \
  --https://github.com/codecreatede/goroutines-alignment-merger \
  --https://github.com/codecreatede/go-alignment-block-estimate \
  --https://github.com/codecreatede/go-alignment-proportion \
- it deals with the complete phylogenomics and perfom all the work of the phylogenomics
- from alignment to the merging, filtering and generating the bootstrap phylogenomes.
  -- filter alignment
  -- removes block of alignment
  -- genomescape alignment merge
  -- serves the alignment to a http server
  -- calculate the identity percentage of the alignment.

```
➜  alignmentGO git:(main) ✗ go run main.go
This is a complete alignmentGo package for the phylogenomics and whole genome alignment

Usage:
  flags [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  estimate
  help        Help about any command
  merge
  unified

Flags:
  -h, --help   help for flags

Use "flags [command] --help" for more information about a command.
➜  alignmentGO git:(main) ✗ go run main.go estimate -h
This estimates the site by site alignment estimates across the alignment and it takes the first genome as a reference for the estimation of the site variablitiy

Usage:
  flags estimate [flags]

Flags:
  -h, --help           help for estimate
  -a, --title string   alignment file for the estimate (default "alignment file")
➜  alignmentGO git:(main) ✗ go run main.go merge -h
This merges all the alignment and gives the alignment ID header you have specified

Usage:
  flags merge [flags]

Flags:
  -A, --alignmentfile string   a alignment file (default "align")
  -h, --help                   help for merge
  -T, --title string           alignment title (default "title for the alignment")
➜  alignmentGO git:(main) ✗ go run main.go unified -h
This estimates the site proportion in your whole genome or gene specific alignment

Usage:
  flags unified [flags]

Flags:
  -a, --alignmentfile string   a alignment file (default "align")
  -h, --help                   help for unified
```
=======
  - filter alignment
  - removes block of alignment
  - genomescape alignment merge
  - serves the alignment to a http server
  - calculate the identity percentage of the alignment.
- all the alignment utilites combined with a http server in a single genome scale package.
>>>>>>> origin/main

Gaurav Sablok
