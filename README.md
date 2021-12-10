# go-gimei

[![Build Status](https://github.com/mattn/go-gimei/workflows/test/badge.svg?branch=master)](https://github.com/mattn/go-gimei/actions?query=workflow%3Atest)
[![Codecov](https://codecov.io/gh/mattn/go-gimei/branch/master/graph/badge.svg)](https://codecov.io/gh/mattn/go-gimei)
[![GoDoc](https://godoc.org/github.com/mattn/go-gimei?status.svg)](http://godoc.org/github.com/mattn/go-gimei)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattn/go-gimei)](https://goreportcard.com/report/github.com/mattn/go-gimei)

This project is a golang port of Ruby's [gimei](https://github.com/willnet/gimei).  Import
as library or use as CLI.

go-gimei generates fake data that people's name and address in Japanese and supports
furigana phonetic renderings of kanji.

The project name comes from Japanese '偽名' means a false name. 

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
	fmt.Println(name.Romaji())         // saito haruna
	fmt.Println(name.Last.Kanji())     // 斎藤
	fmt.Println(name.Last.Hiragana())  // さいとう
	fmt.Println(name.Last.Katakana())  // サイトウ
	fmt.Println(name.Last.Romaji())    // saito
	fmt.Println(name.First.Kanji())    // 陽菜
	fmt.Println(name.First.Hiragana()) // はるな
	fmt.Println(name.First.Katakana()) // ハルナ
	fmt.Println(name.First.Romaji())   // haruna
	fmt.Println(name.Sex)              // 女
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

### CLI Usage

```bash
$ gimei [OPTIONS] [ARGS]
```

#### OPTIONS

```
-type string
    display instead of a mixed-gender personal name: 'male', 'female' or 'address'.
    if ARGS is omitted, display it in kanji.
-sep separator
    field separator.
-h, -help
    display usage.
```    

#### ARGS

Arguments for a personal name:

```
to display full name:
    'kanji',
    'hiragana',
    'katakana'
to display last name:
    'last-kanji',
    'last-hiragana',
    'last-katakana'
to display first name:
    'first-kanji',
    'first-hiragana',
    'first-katakana'
to display which it is male/female:
    'is-male',
    'is-female'
```

Arguments for an address:

```
to display address:
    'kanji'
    'hiragana'
    'katakana'
to display prefecture:
    'prefecture-kanji'
    'prefecture-hiragana'
    'prefecture-katakana'
to display town:
    'town-kanji'
    'town-hiragana'
    'town-katakana'
to display city:
    'city-kanji'
    'city-hiragana'
    'city-katakana'
```

#### EXAMPLES

```bash
$ gimei
古賀 正浩
$ gimei kanji katakana
中村 紳一, ナカムラ シンイチ
$ gimei -type address -sep '/' prefecture-kanji town-kanji
滋賀県/田所町
```

## Requirements

golang

## Installation

Install the library.
```
$ go get github.com/mattn/go-gimei
```

### CLI installation

On Go version 1.16 or later, this command works:

```bash
$ go install github.com/mattn/go-gimei/cmd/gimei@latest
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
