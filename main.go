package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-18

A complete Golang package for phylogenomics. It deals with all the alignment and phylogenomics.

*/

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var (
	alignment  string
	title      string
	start      int
	end        int
	upstream   int
	downstream int
)

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "This is a complete alignmentGo package for the phylogenomics and whole genome alignment",
}

var unifiedCmd = &cobra.Command{
	Use:  "unified",
	Long: "This estimates the site proportion in your whole genome or gene specific alignment",
	Run:  unifiedRun,
}

var mergeCmd = &cobra.Command{
	Use:  "merge",
	Long: "This merges all the alignment and gives the alignment ID header you have specified",
	Run:  alignmentMerge,
}

var estimateCmd = &cobra.Command{
	Use:  "estimate",
	Long: "This estimates the site by site alignment estimates across the alignment and it takes the first genome as a reference for the estimation of the site variablitiy",
	Run:  alignmentEstimate,
}

var conservationCmd = &cobra.Command{
	Use:  "conservation",
	Long: "This estimates the conversation score across the alignment and then estimates the conversation score across all the sites",
	Run:  conservationEstimate,
}

var filterCmd = &cobra.Command{
	Use:  "filter",
	Long: "This is used to filter the sequences filter the sequences of a given alignment",
	Run:  removeFunc,
}

var specificCmd = &cobra.Command{
	Use:  "specifcblock",
	Long: "This estimates the site proportion in your whole genome or gene specific alignment",
	Run:  blockFunc,
}

var proportionCmd = &cobra.Command{
	Use:  "proportion",
	Long: "This esimates the proportion of the site conservation",
	Run:  proportionFunc,
}

var eDNACmd = &cobra.Command{
	Use:  "eDNA",
	Long: "This extacts the upstream and the downstream of the given alignment block with the block includes and also the block excluded",
	Run:  eDNAFunc,
}

func init() {
	unifiedCmd.Flags().StringVarP(&alignment, "alignmentfile", "a", "align", "a alignment file")
	mergeCmd.Flags().StringVarP(&alignment, "alignmentfile", "A", "align", "a alignment file")
	mergeCmd.Flags().StringVarP(&title, "title", "T", "title for the alignment", "alignment title")
	estimateCmd.Flags().
		StringVarP(&alignment, "title", "a", "alignment file", "alignment file for the estimate")
	conservationCmd.Flags().
		StringVarP(&alignment, "alignmentfile", "a", "alignment file convservation score", "alignment to be analyzed")
	filterCmd.Flags().
		StringVarP(&alignment, "alignmentfile", "a", "alignment file for filter", "alignment to be filtered")
	specificCmd.Flags().StringVarP(&alignment, "alignmentfile", "a", "align", "a alignment file")
	specificCmd.Flags().IntVarP(&start, "startcoordinate", "s", 1, "start of the alignment block")
	specificCmd.Flags().IntVarP(&end, "endcoordinate", "e", 40, "end of the alignment block")
	proportionCmd.Flags().StringVarP(&alignment, "alignmentfile", "a", "align", "a alignment file")
	eDNACmd.Flags().StringVarP(&alignment, "alignmentfile", "a", "align", "a alignment file")
	eDNACmd.Flags().IntVarP(&start, "startcoordinate", "s", 1, "start of the alignment block")
	eDNACmd.Flags().IntVarP(&end, "endcoordinate", "e", 40, "end of the alignment block")
	eDNACmd.Flags().IntVarP(&upstream, "upstream-alignment", "u", 4, "upstream of the alignment")
	eDNACmd.Flags().
		IntVarP(&downstream, "downstream-alignment", "d", 4, "downstream of the alignment")

	rootCmd.AddCommand(unifiedCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.AddCommand(estimateCmd)
	rootCmd.AddCommand(conservationCmd)
	rootCmd.AddCommand(filterCmd)
	rootCmd.AddCommand(specificCmd)
	rootCmd.AddCommand(proportionCmd)
	rootCmd.AddCommand(eDNACmd)
}

func unifiedRun(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignIDcapture := []alignmentID{}
	alignSeqcapture := []alignmentSeq{}
	sequenceCap := []string{}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignIDcapture = append(alignIDcapture, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignSeqcapture = append(alignSeqcapture, alignmentSeq{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceCap = append(sequenceCap, string(line))
		}
	}

	counterA := 0
	counterT := 0
	counterG := 0
	counterC := 0

	counterAsite := []int{}
	counterTsite := []int{}
	counterGsite := []int{}
	counterCsite := []int{}

	for i := 0; i < len(sequenceCap)-1; i++ {
		for j := 0; j < len(sequenceCap[0]); j++ {
			if string(sequenceCap[i][j]) == "A" && string(sequenceCap[i+1][j]) == "A" {
				counterA++
				counterAsite = append(counterAsite, j)
			}
			if string(sequenceCap[i][j]) == "T" && string(sequenceCap[i+1][j]) == "T" {
				counterT++
				counterTsite = append(counterTsite, j)
			}
			if string(sequenceCap[i][j]) == "G" && string(sequenceCap[i+1][j]) == "G" {
				counterG++
				counterGsite = append(counterGsite, j)
			}
			if string(sequenceCap[i][j]) == "C" && string(sequenceCap[i+1][j]) == "C" {
				counterC++
				counterCsite = append(counterCsite, j)
			}

		}
	}
	fmt.Printf(
		"The alignment counts for the A unified bases to the rest of the same sites in the block are: %d",
		counterA,
	)
	fmt.Printf(
		"The alignment counts for the T unified bases to the rest of the same sites in the block are: %d",
		counterT,
	)
	fmt.Printf(
		"The alignment counts for the G unified bases to the rest of the same sites in the block are: %d",
		counterG,
	)
	fmt.Printf(
		"The alignment counts for the C unified bases to the rest of the same sites in the block are: %d",
		counterC,
	)

	for i := range counterAsite {
		fmt.Printf("The A sites are %d\n", counterAsite[i])
	}

	for i := range counterTsite {
		fmt.Printf("The T sites are %d", counterTsite[i])
	}

	for i := range counterGsite {
		fmt.Printf("The G sites are %d", counterGsite[i])
	}
	for i := range counterCsite {
		fmt.Printf("The C sites are %d", counterCsite[i])
	}
}

func alignmentMerge(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignmentFilterID struct {
		id string
	}

	type alignmentFilterSeq struct {
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	fRead := bufio.NewScanner(fOpen)
	alignmentReadID := []alignmentID{}
	alignmentReadSeq := []alignmentSeq{}
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentReadID = append(alignmentReadID, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentReadSeq = append(alignmentReadSeq, alignmentSeq{
				seq: string(line),
			})
		}
	}

	var seqKeep string
	for b := range alignmentReadSeq {
		for j := range alignmentReadSeq[b].seq {
			seqKeep += string(alignmentReadSeq[b].seq[j])
		}
	}

	fmt.Println(">", title, "\n", seqKeep)
}

func alignmentEstimate(cmd *cobra.Command, args []string) {
	type alignmentIDStore struct {
		id string
	}

	type alignmentSeqStore struct {
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignmentID := []alignmentIDStore{}
	alignmentSeq := []alignmentSeqStore{}
	sequenceSpec := []string{}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentID = append(alignmentID, alignmentIDStore{
				id: strings.Replace(string(line), ">", "", -1),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentSeq = append(alignmentSeq, alignmentSeqStore{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceSpec = append(sequenceSpec, string(line))
		}
	}

	counterAT := 0
	counterAG := 0
	counterAC := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "A" && string(sequenceSpec[i+1][j]) == "T" {
				counterAT++
			}
			if string(sequenceSpec[i][j]) == "A" && string(sequenceSpec[i+1][j]) == "C" {
				counterAG++
			}
			if string(sequenceSpec[i][j]) == "A" && string(sequenceSpec[i+1][j]) == "G" {
				counterAC++
			}
		}
	}

	counterTG := 0
	counterTC := 0
	counterTA := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "T" && string(sequenceSpec[i+1][j]) == "G" {
				counterTA++
			}
			if string(sequenceSpec[i][j]) == "T" && string(sequenceSpec[i+1][j]) == "C" {
				counterTC++
			}
			if string(sequenceSpec[i][j]) == "T" && string(sequenceSpec[i+1][j]) == "A" {
				counterTA++
			}
		}
	}

	counterGC := 0
	counterGA := 0
	counterGT := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "G" && string(sequenceSpec[i+1][j]) == "C" {
				counterGC++
			}
			if string(sequenceSpec[i][j]) == "G" && string(sequenceSpec[i+1][j]) == "A" {
				counterGA++
			}
			if string(sequenceSpec[i][j]) == "G" && string(sequenceSpec[i+1][j]) == "T" {
				counterGT++
			}
		}
	}

	counterCA := 0
	counterCT := 0
	counterCG := 0

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if string(sequenceSpec[i][j]) == "C" && string(sequenceSpec[i+1][j]) == "A" {
				counterCA++
			}
			if string(sequenceSpec[i][j]) == "C" && string(sequenceSpec[i+1][j]) == "T" {
				counterCT++
			}
			if string(sequenceSpec[i][j]) == "C" && string(sequenceSpec[i+1][j]) == "G" {
				counterCG++
			}
		}
	}

	fmt.Printf(
		"The collinearity block for A as a base pattern and T as a mismatch is %d\n", counterAT)
	fmt.Printf("The collinearity block for A as a base pattern G as a mismatch is %d", counterAG)
	fmt.Printf(
		"The collinearity block for A as a base pattern and C as a mismatch is %d",
		counterAC,
	)

	fmt.Printf(
		"The collinearity block for T as a base pattern and G as a mismatch is %d",
		counterTG,
	)
	fmt.Printf("The collinearity block for T as a base pattern C as a mismatch is  %d", counterTC)
	fmt.Printf(
		"The collinearity block for T as a base pattern and A as a mismatch is %d",
		counterTA,
	)

	fmt.Printf(
		"The collinearity block for G as a base pattern and C as a mismatch is %d",
		counterGC,
	)
	fmt.Printf("The collinearity block for G as a base pattern A as a mismatch is  %d", counterGA)
	fmt.Printf(
		"The collinearity block for G as a base pattern and T as a mismatch is %d",
		counterGT,
	)

	fmt.Printf(
		"The collinearity block for C as a base pattern and A as a mismatch is %d",
		counterCA,
	)
	fmt.Printf("The collinearity block for C as a base pattern T as a mismatch is  %d", counterCT)
	fmt.Printf(
		"The collinearity block for C as a base pattern and G as a mismatch is %d",
		counterCG,
	)
}

func conservationEstimate(cmd *cobra.Command, args []string) {
	type alignmentIDStore struct {
		id string
	}

	type alignmentSeqStore struct {
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignmentID := []alignmentIDStore{}
	alignmentSeq := []alignmentSeqStore{}
	sequenceSpec := []string{}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentID = append(alignmentID, alignmentIDStore{
				id: strings.Replace(string(line), ">", "", -1),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentSeq = append(alignmentSeq, alignmentSeqStore{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceSpec = append(sequenceSpec, string(line))
		}
	}

	alignmentMatrixMatch := []string{}
	alignmentMatrixMisMatch := []string{}

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if sequenceSpec[i][j] != sequenceSpec[i+1][j] {
				alignmentMatrixMatch = append(alignmentMatrixMatch, string(sequenceSpec[i][j]))
			}
		}
	}

	for i := 0; i < len(sequenceSpec)-1; i++ {
		for j := 0; j < len(sequenceSpec[0]); j++ {
			if sequenceSpec[i][j] == sequenceSpec[i+1][j] {
				alignmentMatrixMisMatch = append(
					alignmentMatrixMisMatch,
					string(sequenceSpec[i][j]),
				)
			}
		}
	}

	fmt.Printf(
		"The number of the matches across the alignments blocks are: %d",
		len(alignmentMatrixMatch),
	)

	fmt.Printf(
		"The number of the mismatches across the alignments blocks are: %d",
		len(alignmentMatrixMisMatch),
	)

	add := len(sequenceSpec[0]) / len(alignmentMatrixMatch)

	fmt.Println(add)

	fmt.Printf("The sequence conservation score for the matrix alignment is: %d", add)
}

func removeFunc(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignmentFilterID struct {
		id string
	}

	type alignmentFilterSeq struct {
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	fRead := bufio.NewScanner(fOpen)
	alignmentReadID := []alignmentID{}
	alignmentReadSeq := []alignmentSeq{}
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentReadID = append(alignmentReadID, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentReadSeq = append(alignmentReadSeq, alignmentSeq{
				seq: string(line),
			})
		}
	}

	var seqFilter string

	for i := range alignmentReadSeq {
		for j := range alignmentReadSeq[i].seq {
			if string(alignmentReadSeq[i].seq[j]) == "-" {
				continue
			} else {
				seqFilter += string(alignmentReadSeq[i].seq[j])
			}
		}
	}

	fmt.Println(">", title, "\n", seqFilter)
}

func blockFunc(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignBlock struct {
		id  string
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignIDcapture := []alignmentID{}
	alignSeqcapture := []alignmentSeq{}
	sequenceCap := []string{}
	sequenceID := []string{}
	alignmentBlock := []alignBlock{}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignIDcapture = append(alignIDcapture, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignSeqcapture = append(alignSeqcapture, alignmentSeq{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceCap = append(sequenceCap, string(line))
		}
		if strings.HasPrefix(string(line), ">") {
			sequenceID = append(sequenceID, string(line))
		}
	}

	for i := 0; i < len(sequenceID); i++ {
		alignmentBlock = append(alignmentBlock, alignBlock{
			id:  string((sequenceID[i])),
			seq: string((sequenceCap[i][start:end])),
		})
	}

	for i := range alignmentBlock {
		fmt.Println(alignmentBlock[i].id, "\t", alignmentBlock[i].seq)
	}
}

func proportionFunc(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignIDcapture := []alignmentID{}
	alignSeqcapture := []alignmentSeq{}
	sequenceCap := []string{}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignIDcapture = append(alignIDcapture, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignSeqcapture = append(alignSeqcapture, alignmentSeq{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceCap = append(sequenceCap, string(line))
		}
	}

	counterA := 0
	counterT := 0
	counterG := 0
	counterC := 0

	counterAsite := []int{}
	counterTsite := []int{}
	counterGsite := []int{}
	counterCsite := []int{}

	for i := 0; i < len(sequenceCap)-1; i++ {
		for j := 0; j < len(sequenceCap[0]); j++ {
			if string(sequenceCap[i][j]) == "A" && string(sequenceCap[i+1][j]) == "A" {
				counterA++
				counterAsite = append(counterAsite, j)
			}
			if string(sequenceCap[i][j]) == "T" && string(sequenceCap[i+1][j]) == "T" {
				counterT++
				counterTsite = append(counterTsite, j)
			}
			if string(sequenceCap[i][j]) == "G" && string(sequenceCap[i+1][j]) == "G" {
				counterG++
				counterGsite = append(counterGsite, j)
			}
			if string(sequenceCap[i][j]) == "C" && string(sequenceCap[i+1][j]) == "C" {
				counterC++
				counterCsite = append(counterCsite, j)
			}

		}
	}
	fmt.Printf(
		"The alignment counts for the A unified bases to the rest of the same sites in the block are: %d",
		counterA,
	)
	fmt.Printf(
		"The alignment counts for the T unified bases to the rest of the same sites in the block are: %d",
		counterT,
	)
	fmt.Printf(
		"The alignment counts for the G unified bases to the rest of the same sites in the block are: %d",
		counterG,
	)
	fmt.Printf(
		"The alignment counts for the C unified bases to the rest of the same sites in the block are: %d",
		counterC,
	)

	for i := range counterAsite {
		fmt.Printf("The A sites are %d", counterAsite[i])
	}

	for i := range counterTsite {
		fmt.Printf("The T sites are %d", counterTsite[i])
	}

	for i := range counterGsite {
		fmt.Printf("The G sites are %d", counterGsite[i])
	}
	for i := range counterCsite {
		fmt.Printf("The C sites are %d", counterCsite[i])
	}
}

func eDNAFunc(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignBlock struct {
		id  string
		seq string
	}

	type updownStream struct {
		id  string
		seq string
	}

	type upstreamStart struct {
		id  string
		seq string
	}

	type downstreamStart struct {
		id  string
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignIDcapture := []alignmentID{}
	alignSeqcapture := []alignmentSeq{}
	sequenceCap := []string{}
	sequenceID := []string{}
	alignmentBlock := []alignBlock{}
	upstreamBlock := []updownStream{}
	upstreamfinal := start - upstream
	downstreamfinal := end + downstream
	upstreamS := []upstreamStart{}
	downstreamS := []downstreamStart{}

	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignIDcapture = append(alignIDcapture, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignSeqcapture = append(alignSeqcapture, alignmentSeq{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceCap = append(sequenceCap, string(line))
		}
		if strings.HasPrefix(string(line), ">") {
			sequenceID = append(sequenceID, string(line))
		}
	}

	for i := 0; i < len(sequenceID); i++ {
		alignmentBlock = append(alignmentBlock, alignBlock{
			id:  string(sequenceID[i]),
			seq: string(sequenceCap[i][start:end]),
		})
	}

	for i := range alignmentBlock {
		fmt.Println("This is the alignment block that has been extracted")
		fmt.Println(alignmentBlock[i].id, "\t", alignmentBlock[i].seq)
	}

	for i := 0; i < len(sequenceID); i++ {
		upstreamBlock = append(upstreamBlock, updownStream{
			id:  string(sequenceID[i]),
			seq: string(sequenceCap[i][upstreamfinal:downstreamfinal]),
		})
	}
	fmt.Println(
		"These are the upstream and the downstream blocks for the chosen block including the block",
	)
	for i := range upstreamBlock {
		fmt.Println(upstreamBlock[i].id, "\t", upstreamBlock[i].seq)
	}

	for i := 0; i < len(sequenceID); i++ {
		upstreamS = append(upstreamS, upstreamStart{
			id:  string(sequenceID[i]),
			seq: string(sequenceCap[i][upstreamfinal:start]),
		})
	}
	for i := 0; i < len(sequenceID); i++ {
		downstreamS = append(downstreamS, downstreamStart{
			id:  string(sequenceID[i]),
			seq: string(sequenceCap[i][end:downstreamfinal]),
		})
	}

	fmt.Println("The upstream from the given position till the start is given below:")
	for i := range upstreamS {
		fmt.Println(upstreamS[i].id, "\t", upstreamS[i].seq)
	}
	fmt.Println(
		"The downstream from the end to the given downstream coordinate is given below:",
	)
	for i := range downstreamS {
		fmt.Println(downstreamS[i].id, "\t", downstreamS[i].seq)
	}
}
