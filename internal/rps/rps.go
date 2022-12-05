package rps

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type MatchResult int

const (
	DRAW MatchResult = iota
	DEFEAT
	WIN
)

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

type HandPlay interface {
	Compare(other HandPlay) MatchResult
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

type RpsMatch struct {
	PlayerHand   HandPlay
	OpponentHand HandPlay
}

func (r RpsMatch) Score() int {
	score := 0

	switch r.PlayerHand.(type) {
	case Rock:
		score += 1
	case Paper:
		score += 2
	case Scissors:
		score += 3
	default:
		panic("Unrecognized hand played")
	}

	switch r.PlayerHand.Compare(r.OpponentHand) {
	case WIN:
		score += 6
	case DRAW:
		score += 3
	case DEFEAT:
		score += 0
	}

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

	return RpsMatch{
		PlayerHand:   playerHand,
		OpponentHand: oppHand,
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
