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
		t.Errorf("FindNameByKanji not found: %s", target)
	}
	target = name.Hiragana()
	if gimei.FindNameByHiragana(target) == nil {
		t.Errorf("FindNameByHiragana not found: %s", target)
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

func TestFindNameInvalidParams(t *testing.T) {
	target := "oneword"

	if gimei.FindNameByRomaji(target) != nil {
		t.Errorf("FindNameByRomaji should return nil: %s", target)
	}
}

func TestFindAddress(t *testing.T) {
	var target string
	addr := gimei.NewAddress()

	target = addr.Kanji()
	if gimei.FindAddressByKanji(target) == nil {
		t.Errorf("FindAddressByKanji not found: %s", target)
	}
	target = addr.Hiragana()
	if gimei.FindAddressByHiragana(target) == nil {
		t.Errorf("FindAddressByHiragana not found: %s", target)
	}
	target = addr.Katakana()
	if gimei.FindAddressByKatakana(target) == nil {
		t.Errorf("FindAddressByKatakana not found: %q", target)
	}
}

// Prefecture/City/Town.Romaji() should return empty string
func TestEmptyRomaji(t *testing.T) {
	if gimei.NewPrefecture().Romaji() != "" {
		t.Errorf("Prefecture.Romaji() should return empty string")
	}
	if gimei.NewCity().Romaji() != "" {
		t.Errorf("City.Romaji() should return empty string")
	}
	if gimei.NewTown().Romaji() != "" {
		t.Errorf("Town.Romaji() should return empty string")
	}
}

func TestCheckRaceCondition(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Run(fmt.Sprintf("TestCheckRaceCondition: %v", i), func(t *testing.T) {
			t.Parallel()
			gimei.NewName()
			gimei.NewDog()
			gimei.NewCat()
			gimei.NewMale()
			gimei.NewFemale()
			gimei.NewMaleDog()
			gimei.NewFemaleDog()
			gimei.NewMaleCat()
			gimei.NewFemaleCat()
			gimei.NewAddress()
			gimei.NewPrefecture()
			gimei.NewTown()
			gimei.NewCity()
			gimei.NewPostalCode()
			gimei.FindNameByKanji("小林 顕士")
			gimei.FindAddressByKanji("岡山県大島郡大和村稲木町")
			gimei.CountData()
		})
	}
}

func TestPostalCode(t *testing.T) {
	gimei.SetRandom(rand.New(rand.NewSource(42)))

	postal := gimei.NewPostalCode()
	if postal == nil {
		t.Fatal("NewPostalCode should not return nil")
	}

	if postal.Kanji() == "" {
		t.Fatal("PostalCode.Kanji() should not return empty string")
	}

	if postal.String() == "" {
		t.Fatal("PostalCode.String() should not return empty string")
	}
}
