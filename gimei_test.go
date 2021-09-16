package gimei_test

import (
	"math/rand"
	"testing"

	"github.com/mattn/go-gimei"
)

func TestDeterministicRandom(t *testing.T) {
	var prev, curr [7]string

	gimei.SetRandom(rand.New(rand.NewSource(42)))
	storeResults(prev[:])

	gimei.SetRandom(rand.New(rand.NewSource(42)))
	storeResults(curr[:])

	// expect same result
	for i := 0; i < len(curr); i++ {
		if prev[i] != curr[i] {
			t.Errorf("curr[%d] == %q, want %q", i, curr[i], prev[i])
		}
	}
}

// store string results of gimei functions
func storeResults(results []string) {
	results[0] = gimei.NewName().String()
	results[1] = gimei.NewMale().String()
	results[2] = gimei.NewFemale().String()
	results[3] = gimei.NewAddress().String()
	results[4] = gimei.NewPrefecture().String()
	results[5] = gimei.NewCity().String()
	results[6] = gimei.NewTown().String()
}
