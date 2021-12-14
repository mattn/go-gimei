package gimei_test

import (
	"fmt"

	"github.com/mattn/go-gimei"
)

func ExampleGimeiName() {
	name := gimei.FindNameByKanji("小林 顕士")
	fmt.Println(name)
	fmt.Println(name.Kanji())
	fmt.Println(name.Hiragana())
	fmt.Println(name.Katakana())
	fmt.Println(name.Romaji())
	fmt.Println(name.Last.Kanji())
	fmt.Println(name.Last.Hiragana())
	fmt.Println(name.Last.Katakana())
	fmt.Println(name.Last.Romaji())
	fmt.Println(name.First.Kanji())
	fmt.Println(name.First.Hiragana())
	fmt.Println(name.First.Katakana())
	fmt.Println(name.First.Romaji())
	fmt.Println(name.Sex)
	// Output:
	// 小林 顕士
	// 小林 顕士
	// こばやし けんじ
	// コバヤシ ケンジ
	// Kenji Kobayashi
	// 小林
	// こばやし
	// コバヤシ
	// Kobayashi
	// 顕士
	// けんじ
	// ケンジ
	// Kenji
	// 男
}

func ExampleGimeiAddress() {
	address := gimei.FindAddressByKanji("岡山県大島郡大和村稲木町")
	fmt.Println(address)
	fmt.Println(address.Kanji())
	fmt.Println(address.Hiragana())
	fmt.Println(address.Katakana())
	// Output:
	// 岡山県大島郡大和村稲木町
	// 岡山県大島郡大和村稲木町
	// おかやまけんおおしまぐんやまとそんいなぎちょう
	// オカヤマケンオオシマグンヤマトソンイナギチョウ
}
