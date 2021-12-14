package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mattn/go-gimei"
)

func doName(name *gimei.Name, args []string, sep string) {
	if len(args) == 0 {
		args = []string{"name"}
	}
	for i, arg := range args {
		var ret string
		switch arg {
		case "name":
			ret = name.String() // 斎藤 陽菜
		case "kanji":
			ret = name.Kanji() // 斎藤 陽菜
		case "hiragana":
			ret = name.Hiragana() // さいとう はるな
		case "katakana":
			ret = name.Katakana() // サイトウ ハルナ
		case "romaji":
			ret = name.Romaji() // saito haruna
		case "last-name":
			ret = name.Last.String() // 斎藤
		case "last-kanji":
			ret = name.Last.Kanji() // 斎藤
		case "last-hiragana":
			ret = name.Last.Hiragana() // さいとう
		case "last-katakana":
			ret = name.Last.Katakana() // サイトウ
		case "last-romaji":
			ret = name.Last.Romaji() // saito
		case "first-name":
			ret = name.First.String() // 陽菜
		case "first-kanji":
			ret = name.First.Kanji() // 陽菜
		case "first-hiragana":
			ret = name.First.Hiragana() // はるな
		case "first-katakana":
			ret = name.First.Katakana() // ハルナ
		case "first-romaji":
			ret = name.First.Romaji() // haruna
		case "is-male":
			ret = fmt.Sprint(name.IsMale()) // false
		case "is-female":
			ret = fmt.Sprint(name.IsFemale()) // false
		}
		if ret != "" {
			if i > 0 {
				fmt.Print(sep)
			}
			fmt.Print(ret)
		}
	}
}

func doAddress(address *gimei.Address, args []string, sep string) {
	if len(args) == 0 {
		args = []string{"name"}
	}
	for i, arg := range args {
		var ret string
		switch arg {
		case "name":
			ret = address.String() // 岡山県大島郡大和村稲木町
		case "kanji":
			ret = address.Kanji() // 岡山県大島郡大和村稲木町
		case "hiragana":
			ret = address.Hiragana() // おかやまけんおおしまぐんやまとそんいなぎちょう
		case "katakana":
			ret = address.Katakana() // オカヤマケンオオシマグンヤマトソンイナギチョウ
		case "prefecture-name":
			ret = address.Prefecture.String() // 岡山県
		case "prefecture-kanji":
			ret = address.Prefecture.Kanji() // 岡山県
		case "prefecture-hiragana":
			ret = address.Prefecture.Hiragana() // おかやまけん
		case "prefecture-katakana":
			ret = address.Prefecture.Katakana() // オカヤマケン
		case "town-name":
			ret = address.Town.String() // 大島郡大和村
		case "town-kanji":
			ret = address.Town.Kanji() // 大島郡大和村
		case "town-hiragana":
			ret = address.Town.Hiragana() // おおしまぐんやまとそん
		case "town-katakana":
			ret = address.Town.Katakana() // オオシマグンヤマトソン
		case "city-name":
			ret = address.City.String() // 稲木町
		case "city-kanji":
			ret = address.City.Kanji() // 稲木町
		case "city-hiragana":
			ret = address.City.Hiragana() // いなぎちょう
		case "city-katakana":
			ret = address.City.Katakana() // イナギチョウ
		}
		if ret != "" {
			if i > 0 {
				fmt.Print(sep)
			}
			fmt.Print(ret)
		}
	}
}

func main() {
	var t string
	var sep string
	flag.StringVar(&t, "type", "name", "type of generator(name/male/female/address)")
	flag.StringVar(&sep, "sep", ", ", "separator")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of gimei:
  -sep separator
        field separator (default ",")
  -type string
        type of generator(name/male/female/address) (default "name")
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

	switch t {
	case "name":
		doName(gimei.NewName(), flag.Args(), sep)
	case "male":
		doName(gimei.NewMale(), flag.Args(), sep)
	case "female":
		doName(gimei.NewFemale(), flag.Args(), sep)
	case "address":
		doAddress(gimei.NewAddress(), flag.Args(), sep)
	case "dog":
		doName(gimei.NewDog(), flag.Args(), sep)
	case "cat":
		doName(gimei.NewCat(), flag.Args(), sep)
	default:
		flag.Usage()
		os.Exit(2)
	}
	fmt.Println()
}
