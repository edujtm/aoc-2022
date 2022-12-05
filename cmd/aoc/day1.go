package aoc

import (
	"fmt"
	"os"

	"github.com/edujtm/aoc-2022/internal/elf"
	"github.com/spf13/cobra"
)

var part2 bool

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "gets the maximum calories from the collection of elfs",
	Args:  cobra.ExactArgs(1),
	Run:   FindMaxCalories,
}

func init() {
	day1Cmd.Flags().BoolVarP(&part2, "second", "s", false, "run part 2 of the challenge")
	rootCmd.AddCommand(day1Cmd)
}

func FindMaxCalories(cmd *cobra.Command, args []string) {
	filepath := args[0]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Some error ocurred while opening elf collection file: %s", err)
		os.Exit(1)
	}

	elfs := elf.NewCollectionFromReader(file)

	if !part2 {
		maxCal := elf.MaxCalories(elfs)

		fmt.Printf("Max calories: %d\n", maxCal)
	} else {
		maxCalories := elf.NthMaxCalories(elfs, 3)

		total := 0
		for idx, cal := range maxCalories {
			fmt.Printf("%d max calories: %d\n", idx, cal)
			total += cal
		}
		fmt.Printf("Total calories for %d biggest elfs: %d\n", 3, total)
	}
}
