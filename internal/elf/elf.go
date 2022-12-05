package elf

import (
	"bufio"
	"io"
	"sort"
	"strconv"

	"github.com/edujtm/aoc-2022/internal/strext"
)

type elf struct {
	TotalCalories int
}

func NewCollectionFromReader(reader io.Reader) []elf {
	scanner := bufio.NewScanner(reader)

	elfs := []elf{}

	e := elf{}
	for scanner.Scan() {
		input := scanner.Text()

		if strext.IsBlank(input) {
			elfs = append(elfs, e)
			e = elf{}
		} else {
			cal, _ := strconv.Atoi(input)
			e.TotalCalories += cal
		}
	}
	return elfs
}

func MaxCalories(elfs []elf) int {
	maxCal := 0
	for _, elf := range elfs {
		if elf.TotalCalories > maxCal {
			maxCal = elf.TotalCalories
		}
	}

	return maxCal
}

func NthMaxCalories(elfs []elf, nth int) []int {
	copies := make([]elf, len(elfs))
	copy(copies, elfs)

	sort.Slice(copies, func(i, j int) bool {
		return copies[i].TotalCalories > copies[j].TotalCalories
	})

	result := []int{}
	for _, elf := range copies[:nth] {
		result = append(result, elf.TotalCalories)
	}
	return result
}
