package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-gimei"
)

func doName(name *gimei.Name, arg string) string {
	switch arg {
	case "kanji":
		return name.Kanji() // 斎藤 陽菜
	case "hiragana":
		return name.Hiragana() // さいとう はるな
	case "katakana":
		return name.Katakana() // サイトウ ハルナ
	case "romaji":
		return name.Romaji() // Haruna Saito
	case "last-name":
		return name.Last.String() // 斎藤
	case "last-kanji":
		return name.Last.Kanji() // 斎藤
	case "last-hiragana":
		return name.Last.Hiragana() // さいとう
	case "last-katakana":
		return name.Last.Katakana() // サイトウ
	case "last-romaji":
		return name.Last.Romaji() // Saito
	case "first-name":
		return name.First.String() // 陽菜
	case "first-kanji":
		return name.First.Kanji() // 陽菜
	case "first-hiragana":
		return name.First.Hiragana() // はるな
	case "first-katakana":
		return name.First.Katakana() // ハルナ
	case "first-romaji":
		return name.First.Romaji() // Haruna
	case "is-male":
		return fmt.Sprint(name.IsMale()) // false
	case "is-female":
		return fmt.Sprint(name.IsFemale()) // false
	default:
		return name.String() // 斎藤 陽菜
	}
}

func doAddress(address *gimei.Address, arg string) string {
	switch arg {
	case "kanji":
		return address.Kanji() // 岡山県大島郡大和村稲木町
	case "hiragana":
		return address.Hiragana() // おかやまけんおおしまぐんやまとそんいなぎちょう
	case "katakana":
		return address.Katakana() // オカヤマケンオオシマグンヤマトソンイナギチョウ
	case "prefecture-name":
		return address.Prefecture.String() // 岡山県
	case "prefecture-kanji":
		return address.Prefecture.Kanji() // 岡山県
	case "prefecture-hiragana":
		return address.Prefecture.Hiragana() // おかやまけん
	case "prefecture-katakana":
		return address.Prefecture.Katakana() // オカヤマケン
	case "town-name":
		return address.Town.String() // 大島郡大和村
	case "town-kanji":
		return address.Town.Kanji() // 大島郡大和村
	case "town-hiragana":
		return address.Town.Hiragana() // おおしまぐんやまとそん
	case "town-katakana":
		return address.Town.Katakana() // オオシマグンヤマトソン
	case "city-name":
		return address.City.String() // 稲木町
	case "city-kanji":
		return address.City.Kanji() // 稲木町
	case "city-hiragana":
		return address.City.Hiragana() // いなぎちょう
	case "city-katakana":
		return address.City.Katakana() // イナギチョウ
	default:
		return address.String() // 岡山県大島郡大和村稲木町
	}
}

func main() {
	var sep string
	var count bool
	var n int
	flag.IntVar(&n, "n", 1, "N records")
	flag.StringVar(&sep, "sep", ", ", "separator")
	flag.BoolVar(&count, "count", false, "")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of gimei:
  -sep separator
        field separator (default ",")
  -type string
        type of generator(name/male/female/address) (default "name")
  -count
        display count of data
  -h, -help
        display this usage

  Arguments for name/male/female:
    name
    kanji
    hiragana
    katakana
    romaji
    last-name
    last-kanji
    last-hiragana
    last-katakana
    last-romaji
    first-name
    first-kanji
    first-hiragana
    first-katakana
    first-romaji
    is-male
    is-female
  
  Arguments for address:
    name
    kanji
    hiragana
    katakana
    prefecture-name
    prefecture-kanji
    prefecture-hiragana
    prefecture-katakana
    town-name
    town-kanji
    town-hiragana
    town-katakana
    city-name
    city-kanji
    city-hiragana
    city-katakana
`)
	}
	flag.Parse()

	if count {
		fmt.Println(gimei.CountData())
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"name:name"}
	}

	for i := 0; i < n; i++ {
		gimeiName := gimei.NewName()
		gimeiMale := gimei.NewMale()
		gimeiFemale := gimei.NewFemale()
		gimeiAddress := gimei.NewAddress()
		gimeiDog := gimei.NewDog()
		gimeiCat := gimei.NewCat()
		for i, arg := range args {
			tokens := strings.SplitN(arg, ":", 2)
			if len(tokens) == 0 {
				flag.Usage()
				os.Exit(2)
			} else if len(tokens) == 1 {
				tokens = append(tokens, "name")
			}
			var result string
			switch tokens[0] {
			case "name":
				result = doName(gimeiName, tokens[1])
			case "male":
				result = doName(gimeiMale, tokens[1])
			case "female":
				result = doName(gimeiFemale, tokens[1])
			case "address":
				result = doAddress(gimeiAddress, tokens[1])
			case "dog":
				result = doName(gimeiDog, tokens[1])
			case "cat":
				result = doName(gimeiCat, tokens[1])
			default:
				flag.Usage()
				os.Exit(2)
			}
			if i > 0 {
				fmt.Print(sep)
			}
			fmt.Print(result)
		}
		fmt.Println()
	}
}
