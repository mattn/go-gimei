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

### CLI

```bash
$ gimei [OPTION] [ARGS]
```

#### OPTIONS
```
  -type string
        type of generator(name/male/female/address) (default "name")
```
名前(name/male/female) と住所(address) のいずれかを type で指定できる。デフォルトは "name"。
指定したタイプごとに、出力するフィールドを指定できる。
名前(name/male/female) の場合
name/kanji/hiragana/katakana
last-name/last-kanji/last-hiragana/last-katakana
first-name/first-kanji/first-hiragana/first-katakana
is-male/is-female
住所(address)の場合
name/kanji/hiragana/katakana
prefecture-name/prefecture-kanji/prefecture-hiragana/prefecture-katakana
town-name/town-kanji/town-hiragana/town-katakana
city-name/city-kanji/city-hiragana/city-katakana

```
  -sep separator
        フィールドを区切る文字列を指定できる
```

#### EXAMPLES
```
$ gimei
古賀 正浩
$ gimei name katakana
中村 紳一, ナカムラ シンイチ
$ gimei -type address -sep '/' prefecture-kanji town-kanji
滋賀県/田所町
```

### Deterministic Random

go-gimei supports seeding of its pseudo-random number generator to provide
deterministic output of repeated method calls.

```go
import "math/rand"

gimei.SetRandom(rand.New(rand.NewSource(42)))
fmt.Println(gimei.NewName())    // 前川 永麻
fmt.Println(gimei.NewAddress()) // 佐賀県斜里郡斜里町浄法寺町樋口

gimei.SetRandom(rand.New(rand.NewSource(42)))
fmt.Println(gimei.NewName())    // 前川 永麻
fmt.Println(gimei.NewAddress()) // 佐賀県斜里郡斜里町浄法寺町樋口

```

## Requirements

golang

## Installation

```
$ go get github.com/mattn/go-gimei
```

## Running Tests

To run all the tests, do:

```bash
$ go test
```

## License

MIT

Dictionary YAML file is generated from [naist-jdic](https://ja.osdn.net/projects/naist-jdic/).

## Author

Yasuhiro Matsumoto (a.k.a mattn)
