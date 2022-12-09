package rps_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/edujtm/aoc-2022/internal/rps"
)

func TestHandComparison(t *testing.T) {
	testCases := []struct {
		PlayerHand      rps.HandPlay
		OpponentHand    rps.HandPlay
		ExpectedOutcome rps.MatchResult
	}{
		{rps.Rock{}, rps.Paper{}, rps.DEFEAT},
		{rps.Paper{}, rps.Rock{}, rps.WIN},
		{rps.Scissors{}, rps.Paper{}, rps.WIN},
		{rps.Paper{}, rps.Scissors{}, rps.DEFEAT},
		{rps.Rock{}, rps.Scissors{}, rps.WIN},
		{rps.Scissors{}, rps.Rock{}, rps.DEFEAT},
		{rps.Paper{}, rps.Paper{}, rps.DRAW},
		{rps.Scissors{}, rps.Scissors{}, rps.DRAW},
		{rps.Rock{}, rps.Rock{}, rps.DRAW},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Testing match: player - %#v vs %#v", tc.PlayerHand, tc.OpponentHand), func(t *testing.T) {

			matchResult := tc.PlayerHand.Compare(tc.OpponentHand)

			if matchResult != tc.ExpectedOutcome {
				t.Errorf("Expected rock paper scissors match to result in %v but got %v", rps.MatchResult(tc.ExpectedOutcome), matchResult)
			}
		})
	}
}

func TestParseMatch(t *testing.T) {
	testCases := []struct {
		SerializedMatch string
		ExpectedMatch   rps.RpsMatch
	}{
		{"A Y", rps.RpsMatch{rps.Paper{}, rps.Rock{}, rps.DRAW}},
		{"B X", rps.RpsMatch{rps.Rock{}, rps.Paper{}, rps.DEFEAT}},
		{"C Z", rps.RpsMatch{rps.Scissors{}, rps.Scissors{}, rps.WIN}},
	}

	for _, tc := range testCases {
		t.Run("parsing match: %s", func(t *testing.T) {
			match, err := rps.ParseMatch(tc.SerializedMatch)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(match, tc.ExpectedMatch) {
				t.Errorf("expected parsed match to be %#v but got %#v", tc.ExpectedMatch, match)
			}
		})
	}
}

func TestStrategyScore(t *testing.T) {
	testCases := []struct {
		Match         rps.RpsMatch
		ExpectedScore int
	}{
		{rps.RpsMatch{rps.Paper{}, rps.Rock{}, rps.DRAW}, 4},
		{rps.RpsMatch{rps.Rock{}, rps.Paper{}, rps.DEFEAT}, 1},
		{rps.RpsMatch{rps.Scissors{}, rps.Scissors{}, rps.WIN}, 7},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("calculates strategy score for match: %+v", tc.Match), func(t *testing.T) {
			sscore := tc.Match.StrategyScore()

			if sscore != tc.ExpectedScore {
				t.Errorf("Expected strategy score to be %d but got %d", tc.ExpectedScore, sscore)
			}
		})
	}
}
