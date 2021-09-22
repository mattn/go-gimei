package gimei_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mattn/go-gimei"
)

func TestDeterministicRandom(t *testing.T) {
	gimei.SetRandom(rand.New(rand.NewSource(42)))
	prev := collectNewResults()

	gimei.SetRandom(rand.New(rand.NewSource(42)))
	curr := collectNewResults()

	// expect same result
	for i := 0; i < len(curr); i++ {
		if prev[i].String() != curr[i].String() {
			t.Errorf("curr[%d] == %q, want %q", i, curr[i], prev[i])
		}
	}
}

// returns slice of fmt.Stringer which return value of gimei 'New' functions
func collectNewResults() []fmt.Stringer {
	var s []fmt.Stringer

	s = append(s, gimei.NewName())
	s = append(s, gimei.NewMale())
	s = append(s, gimei.NewFemale())
	s = append(s, gimei.NewAddress())
	s = append(s, gimei.NewPrefecture())
	s = append(s, gimei.NewCity())
	s = append(s, gimei.NewTown())

	return s
}
