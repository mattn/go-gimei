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

func TestFindName(t *testing.T) {
	var target string
	name := gimei.NewName()

	target = name.Kanji()
	if gimei.FindNameByKanji(target) == nil {
		t.Errorf("FindNameByKanji not found %s", target)
	}
	target = name.Hiragana()
	if gimei.FindNameByHiragana(target) == nil {
		t.Errorf("FindNameByHiragana not found %s", target)
	}
	target = name.Katakana()
	if gimei.FindNameByKatakana(target) == nil {
		t.Errorf("FindNameByKatakana not found: %s", target)
	}
	target = name.Romaji()
	if gimei.FindNameByRomaji(target) == nil {
		t.Errorf("FindNameByRomaji not found: %s", target)
	}
}

func TestFindAddress(t *testing.T) {
	var target string
	addr := gimei.NewAddress()

	target = addr.Kanji()
	if gimei.FindAddressByKanji(target) == nil {
		t.Errorf("FindAddressByKanji not found %s", target)
	}
	target = addr.Hiragana()
	if gimei.FindAddressByHiragana(target) == nil {
		t.Errorf("FindAddressByHiragana not found %s", target)
	}
	target = addr.Katakana()
	if gimei.FindAddressByKatakana(target) == nil {
		t.Errorf("FindAddressByKatakana not found: %q", target)
	}
}
