# go-gimei

golang port of [gimei](https://github.com/willnet/gimei)

## Usage

```go
package main

import (
	"fmt"

	"github.com/mattn/go-gimei"
)

func main() {
	name := gimei.NewName()
	fmt.Println(name)                  // 斎藤 陽菜
	fmt.Println(name.Kanji())          // 斎藤 陽菜
	fmt.Println(name.Hiragana())       // さいとう はるな
	fmt.Println(name.Katakana())       // サイトウ ハルナ
	fmt.Println(name.Last.Kanji())     // 斎藤
	fmt.Println(name.Last.Hiragana())  // さいとう
	fmt.Println(name.Last.Katakana())  // サイトウ
	fmt.Println(name.First.Kanji())    // 陽菜
	fmt.Println(name.First.Hiragana()) // はるな
	fmt.Println(name.First.Katakana()) // ハルナ
	fmt.Println(name.IsMale())         // false

	male := gimei.NewMale()
	fmt.Println(male)            // 小林 顕士
	fmt.Println(male.IsMale())   // true
	fmt.Println(male.IsFemale()) // false

	address := gimei.NewAddress()
	fmt.Println(address)                       // 岡山県大島郡大和村稲木町
	fmt.Println(address.Kanji())               // 岡山県大島郡大和村稲木町
	fmt.Println(address.Hiragana())            // おかやまけんおおしまぐんやまとそんいなぎちょう
	fmt.Println(address.Katakana())            // オカヤマケンオオシマグンヤマトソンイナギチョウ
	fmt.Println(address.Prefecture)            // 岡山県
	fmt.Println(address.Prefecture.Kanji())    // 岡山県
	fmt.Println(address.Prefecture.Hiragana()) // おかやまけん
	fmt.Println(address.Prefecture.Katakana()) // オカヤマケン
	fmt.Println(address.Town)                  // 大島郡大和村
	fmt.Println(address.Town.Kanji())          // 大島郡大和村
	fmt.Println(address.Town.Hiragana())       // おおしまぐんやまとそん
	fmt.Println(address.Town.Katakana())       // オオシマグンヤマトソン
	fmt.Println(address.City)                  // 稲木町
	fmt.Println(address.City.Kanji())          // 稲木町
	fmt.Println(address.City.Hiragana())       // いなぎちょう
	fmt.Println(address.City.Katakana())       // イナギチョウ

	prefecture := gimei.NewPrefecture()
	fmt.Println(prefecture) // 青森県
}
```

## Requirements

golang

## Installation

```
$ go get github.com/mattn/go-gimei
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
