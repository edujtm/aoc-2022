package rps

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// This challenge would probably be much simpler to solve
// using a 3x3 matrix with match outcomes, but I was already
// too deep into this solution to refactor everything.

type MatchResult int

const (
	DRAW MatchResult = iota
	DEFEAT
	WIN
)

func (mr MatchResult) ResultScore() int {
	switch mr {
	case WIN:
		return 6
	case DRAW:
		return 3
	case DEFEAT:
		return 0
	default:
		panic("Unrecognized match result")
	}
}

func (mr MatchResult) String() string {
	switch mr {
	case DRAW:
		return "DRAW"
	case DEFEAT:
		return "DEFEAT"
	case WIN:
		return "WIN"
	default:
		return fmt.Sprintf("%d", int(mr))
	}
}

var oppHands = map[string]HandPlay{
	"A": Rock{},
	"B": Paper{},
	"C": Scissors{},
}

var myHands = map[string]HandPlay{
	"X": Rock{},
	"Y": Paper{},
	"Z": Scissors{},
}

var expectedResult = map[string]MatchResult{
	"X": DEFEAT,
	"Y": DRAW,
	"Z": WIN,
}

type HandPlay interface {
	Compare(other HandPlay) MatchResult
	Predict(result MatchResult) HandPlay
}

func HandScore(hand HandPlay) int {
	switch hand.(type) {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	default:
		panic("Unrecognized hand played")
	}
}

type Rock struct{}

func (rock Rock) Compare(other HandPlay) MatchResult {
	switch other.(type) {
	case Rock:
		return DRAW
	case Paper:
		return DEFEAT
	case Scissors:
		return WIN
	default:
		panic("Unrecognized hand played.")
	}
}

func (rock Rock) Predict(result MatchResult) HandPlay {
	switch result {
	case WIN:
		return Paper{}
	case DRAW:
		return Rock{}
	case DEFEAT:
		return Scissors{}
	default:
		panic("Unrecognized match result")
	}
}

type Paper struct{}

func (paper Paper) Compare(other HandPlay) MatchResult {
	switch other.(type) {
	case Rock:
		return WIN
	case Paper:
		return DRAW
	case Scissors:
		return DEFEAT
	default:
		panic("Unrecognized hand played.")
	}
}

func (paper Paper) Predict(result MatchResult) HandPlay {
	switch result {
	case WIN:
		return Scissors{}
	case DRAW:
		return Paper{}
	case DEFEAT:
		return Rock{}
	default:
		panic("Unrecognized match result")
	}
}

type Scissors struct{}

func (scissors Scissors) Compare(other HandPlay) MatchResult {
	switch other.(type) {
	case Rock:
		return DEFEAT
	case Paper:
		return WIN
	case Scissors:
		return DRAW
	default:
		panic("Unrecognized hand played")
	}
}

func (scissors Scissors) Predict(result MatchResult) HandPlay {
	switch result {
	case WIN:
		return Rock{}
	case DRAW:
		return Scissors{}
	case DEFEAT:
		return Paper{}
	default:
		panic("Unrecognized match result")
	}
}

type RpsMatch struct {
	PlayerHand     HandPlay
	OpponentHand   HandPlay
	StrategyResult MatchResult // Part 2
}

func (r RpsMatch) Score() int {
	score := 0

	// The match score is the sum of the player's
	// hand score plus the score of the result of
	// the match given by the comparison of the hands
	score += HandScore(r.PlayerHand)

	matchResult := r.PlayerHand.Compare(r.OpponentHand)
	score += matchResult.ResultScore()

	return score
}

func (r RpsMatch) StrategyScore() int {
	score := 0

	// The strategy score is the sum of the result given
	// by the strategy plus the score for the hand
	// that must be played to achieve the strategy result.
	score += r.StrategyResult.ResultScore()

	handPlayed := r.OpponentHand.Predict(r.StrategyResult)

	score += HandScore(handPlayed)

	return score
}

func ParseMatch(matchtext string) (RpsMatch, error) {
	hands := strings.Split(matchtext, " ")
	if len(hands) < 2 {
		return RpsMatch{}, errors.New("invalid rock paper scissors match: missing hand information")
	}

	opphand := hands[0]
	myhand := hands[1]

	playerHand, ok := myHands[myhand]
	if !ok {
		return RpsMatch{}, errors.New("invalid rock paper scissor match: wrong hand representation for player")
	}

	oppHand, ok := oppHands[opphand]
	if !ok {
		return RpsMatch{}, errors.New("invalid rock paper scissors match: wrong hand representation for opponent")
	}

	strategyResult, ok := expectedResult[myhand]
	if !ok {
		return RpsMatch{}, errors.New("invalid strategy result")
	}

	return RpsMatch{
		PlayerHand:     playerHand,
		OpponentHand:   oppHand,
		StrategyResult: strategyResult,
	}, nil
}

type RpsStrategy struct {
	Matches []RpsMatch
}

func (rs RpsStrategy) TotalScore() int {
	total := 0

	for _, match := range rs.Matches {
		total += match.Score()
	}

	return total
}

func (rs RpsStrategy) TotalScoreFollowingStrategy() int {
	total := 0

	for _, match := range rs.Matches {
		total += match.StrategyScore()
	}

	return total
}

func ReadStrategy(reader io.Reader) (RpsStrategy, error) {
	scanner := bufio.NewScanner(reader)

	matches := []RpsMatch{}
	for scanner.Scan() {
		input := scanner.Text()

		match, err := ParseMatch(input)
		if err != nil {
			return RpsStrategy{}, fmt.Errorf("error reading strategy file: %w", err)
		}

		matches = append(matches, match)
	}

	return RpsStrategy{matches}, nil
}
