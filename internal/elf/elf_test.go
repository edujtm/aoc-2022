package elf_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/edujtm/aoc-2022/internal/elf"
)

func TestMaxCalories(t *testing.T) {
	data, err := os.Open("../../data/day01-sample.txt")
	if err != nil {
		t.Fatalf("Could not read data file: %s", err)
	}

	elfs := elf.NewCollectionFromReader(data)
	got := elf.MaxCalories(elfs)
	expected := 24000

	if got != expected {
		t.Errorf("Expected max calores to be %v but got %v", expected, got)
	}
}

func TestMultipleMaxCalories(t *testing.T) {
	data, err := os.Open("../../data/day01-sample.txt")
	if err != nil {
		t.Fatalf("Could not read data file: %s", err)
	}

	elfs := elf.NewCollectionFromReader(data)

	maxCalories := elf.NthMaxCalories(elfs, 2)
	expected := []int{24000, 11000}

	if !reflect.DeepEqual(expected, maxCalories) {
		t.Errorf("expected 2 biggest max calores to be %v but got %v", expected, maxCalories)
	}
}
