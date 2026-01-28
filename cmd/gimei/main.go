package main

import (
	"encoding/json"
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
	case "city-name":
		return address.City.String() // 大島郡大和村
	case "city-kanji":
		return address.City.Kanji() // 大島郡大和村
	case "city-hiragana":
		return address.City.Hiragana() // おおしまぐんやまとそん
	case "city-katakana":
		return address.City.Katakana() // オオシマグンヤマトソン
	case "town-name":
		return address.Town.String() // 稲木町
	case "town-kanji":
		return address.Town.Kanji() // 稲木町
	case "town-hiragana":
		return address.Town.Hiragana() // いなぎちょう
	case "town-katakana":
		return address.Town.Katakana() // イナギチョウ
	default:
		return address.String() // 岡山県大島郡大和村稲木町
	}
}

func doPostalCode(postalCode *gimei.PostalCode, arg string) string {
	switch arg {
	case "kanji":
		return postalCode.Kanji() // 060-0001
	default:
		return postalCode.String() // 060-0001
	}
}

const (
	name    = "gimei"
	version = "0.0.0"
)

var revision = "HEAD"

func main() {
	var sep string
	var count bool
	var jsonOutput bool
	var showVersion bool
	var n int
	flag.IntVar(&n, "n", 1, "N records")
	flag.StringVar(&sep, "sep", ", ", "separator")
	flag.BoolVar(&count, "count", false, "")
	flag.BoolVar(&jsonOutput, "json", false, "output as JSON array")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: gimei [OPTIONS] [ARGS]

  -sep string
        specify string used to separate fields(default: ", ")
  -n number
        display number record(s)
  -count
        display records read from embedded yaml files and exit
  -json
        output as JSON array
  -h, -help
        display this usage and exit
  -v
        show version

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
    city-name
    city-kanji
    city-hiragana
    city-katakana
    town-name
    town-kanji
    town-hiragana
    town-katakana

  Arguments for postal code:
    name
    kanji

  Example:
    $ gimei -n 3 name:name name:hiragana address:name postal:name
    鈴木 真里緒, すずき まりお, 山口県新居浜市森川町, 060-0001
    宮原 秋南, みやはら あきな, 大阪府枝幸郡浜頓別町豊郷, 100-0005
    大内 亮佳, おおうち あきか, 福島県磯城郡田原本町高清水上佐野, 150-0008
`)
	}
	flag.Parse()

	if showVersion {
		fmt.Println(version)
		return
	}

	if count {
		fmt.Println(gimei.CountData())
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"name:name"}
	}

	if jsonOutput {
		var allRecords []map[string]string
		for i := 0; i < n; i++ {
			var (
				gimeiName, gimeiMale, gimeiFemale *gimei.Name       = nil, nil, nil
				gimeiAddress                      *gimei.Address    = nil
				gimeiPostalCode                   *gimei.PostalCode = nil
				gimeiDog, gimeiCat                *gimei.Name       = nil, nil
			)
			record := make(map[string]string)
			for _, arg := range args {
				tokens := strings.SplitN(arg, ":", 2)
				if len(tokens) == 0 {
					flag.Usage()
					os.Exit(2)
				} else if len(tokens) == 1 {
					tokens = append(tokens, "name")
				}
				var result string
				var fieldName string
				switch tokens[0] {
				case "name":
					if gimeiName == nil {
						gimeiName = gimei.NewName()
					}
					result = doName(gimeiName, tokens[1])
					fieldName = "name"
				case "male":
					if gimeiMale == nil {
						gimeiMale = gimei.NewMale()
					}
					result = doName(gimeiMale, tokens[1])
					fieldName = "male"
				case "female":
					if gimeiFemale == nil {
						gimeiFemale = gimei.NewFemale()
					}
					result = doName(gimeiFemale, tokens[1])
					fieldName = "female"
				case "address":
					if gimeiAddress == nil {
						gimeiAddress = gimei.NewAddress()
					}
					result = doAddress(gimeiAddress, tokens[1])
					fieldName = "address"
				case "postal":
					if gimeiPostalCode == nil {
						gimeiPostalCode = gimei.NewPostalCode()
					}
					result = doPostalCode(gimeiPostalCode, tokens[1])
					fieldName = "postal"
				case "dog":
					if gimeiDog == nil {
						gimeiDog = gimei.NewDog()
					}
					result = doName(gimeiDog, tokens[1])
					fieldName = "dog"
				case "cat":
					if gimeiCat == nil {
						gimeiCat = gimei.NewCat()
					}
					result = doName(gimeiCat, tokens[1])
					fieldName = "cat"
				default:
					flag.Usage()
					os.Exit(2)
				}
				record[fieldName] = result
			}
			allRecords = append(allRecords, record)
		}
		jsonData, err := json.Marshal(allRecords)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	} else {
		for i := 0; i < n; i++ {
			var (
				gimeiName, gimeiMale, gimeiFemale *gimei.Name       = nil, nil, nil
				gimeiAddress                      *gimei.Address    = nil
				gimeiPostalCode                   *gimei.PostalCode = nil
				gimeiDog, gimeiCat                *gimei.Name       = nil, nil
			)
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
					if gimeiName == nil {
						gimeiName = gimei.NewName()
					}
					result = doName(gimeiName, tokens[1])
				case "male":
					if gimeiMale == nil {
						gimeiMale = gimei.NewMale()
					}
					result = doName(gimeiMale, tokens[1])
				case "female":
					if gimeiFemale == nil {
						gimeiFemale = gimei.NewFemale()
					}
					result = doName(gimeiFemale, tokens[1])
				case "address":
					if gimeiAddress == nil {
						gimeiAddress = gimei.NewAddress()
					}
					result = doAddress(gimeiAddress, tokens[1])
				case "postal":
					if gimeiPostalCode == nil {
						gimeiPostalCode = gimei.NewPostalCode()
					}
					result = doPostalCode(gimeiPostalCode, tokens[1])
				case "dog":
					if gimeiDog == nil {
						gimeiDog = gimei.NewDog()
					}
					result = doName(gimeiDog, tokens[1])
				case "cat":
					if gimeiCat == nil {
						gimeiCat = gimei.NewCat()
					}
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
}
