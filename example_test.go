package gimei_test

import (
	"fmt"

	"github.com/mattn/go-gimei"
)

func ExampleGimeiName() {
	defer gimei.LoadNames("")
	gimei.LoadNames("data/names_test.yml")
	name := gimei.NewMale()
	fmt.Println(name)
	fmt.Println(name.Kanji())
	fmt.Println(name.Hiragana())
	fmt.Println(name.Katakana())
	fmt.Println(name.Last.Kanji())
	fmt.Println(name.Last.Hiragana())
	fmt.Println(name.Last.Katakana())
	fmt.Println(name.First.Kanji())
	fmt.Println(name.First.Hiragana())
	fmt.Println(name.First.Katakana())
	// Output:
	// 小林 顕士
	// 小林 顕士
	// こばやし けんじ
	// コバヤシ ケンジ
	// 小林
	// こばやし
	// コバヤシ
	// 顕士
	// けんじ
	// ケンジ
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
