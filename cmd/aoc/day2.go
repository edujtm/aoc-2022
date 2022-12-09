package aoc

import (
	"fmt"
	"os"

	"github.com/edujtm/aoc-2022/internal/rps"
	"github.com/spf13/cobra"
)

var day2part2 bool

var day2Cmd = &cobra.Command{
	Use:   "day2 [filepath of strategy]",
	Short: "calculates the total score for the given strategy",
	Args:  cobra.ExactArgs(1),
	Run:   CalculateStrategyScore,
}

func init() {
	day2Cmd.Flags().BoolVarP(&day2part2, "second", "s", false, "run part 2 of the challenge")
	rootCmd.AddCommand(day2Cmd)
}

func CalculateStrategyScore(cmd *cobra.Command, args []string) {
	filepath := args[0]

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening strategy file: %s", err)
		os.Exit(1)
	}

	strategy, err := rps.ReadStrategy(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing strategy file: %s", err)
		os.Exit(1)
	}

	if !day2part2 {
		fmt.Printf("Total score: %d\n", strategy.TotalScore())
	} else {
		fmt.Printf("Total strategy score: %d\n", strategy.TotalScoreFollowingStrategy())
	}
}
