package gimei

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	names       name
	addresses   address
	onceName    sync.Once
	onceAddress sync.Once
	r           *rand.Rand
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type Item []string

func (i Item) String() string {
	return i.Kanji()
}

func (i Item) Kanji() string {
	return i[0]
}

func (i Item) Hiragana() string {
	return i[1]
}

func (i Item) Katakana() string {
	return i[2]
}

type Sex int

const (
	Male Sex = iota
	Female
)

type name struct {
	FirstName struct {
		Male   []Item `yaml:"male"`
		Female []Item `yaml:"female"`
	} `yaml:"first_name"`
	LastName [][]string `yaml:"last_name"`
}

type Name struct {
	First Item
	Last  Item
	Sex   Sex
}

func loadNames() {
	rp := "src/github.com/mattn/go-gimei/data/names.yml"
	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		f := filepath.Join(p, rp)
		if _, err := os.Stat(f); err == nil {
			if b, err := ioutil.ReadFile(f); err == nil {
				if err = yaml.Unmarshal(b, &names); err == nil {
					return
				}
			}
		}
	}
	panic("failed to load names data")
}

func (n *Name) String() string {
	return n.Kanji()
}

func (n *Name) Kanji() string {
	return n.Last.Kanji() + " " + n.First.Kanji()
}

func (n *Name) Hiragana() string {
	return n.Last.Hiragana() + " " + n.First.Hiragana()
}

func (n *Name) Katakana() string {
	return n.Last.Katakana() + " " + n.First.Katakana()
}

func (n *Name) IsMale() bool {
	return n.Sex == Male
}

func (n *Name) IsFemale() bool {
	return n.Sex == Female
}

func NewName() *Name {
	if rand.Int()%2 == 0 {
		return NewMale()
	} else {
		return NewFemale()
	}
}

func NewMale() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Male[r.Int()%len(names.FirstName.Male)],
		Last:  names.LastName[r.Int()%len(names.LastName)],
		Sex:   Male,
	}
}

func NewFemale() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Female[r.Int()%len(names.FirstName.Female)],
		Last:  names.LastName[r.Int()%len(names.LastName)],
		Sex:   Female,
	}
}

type address struct {
	Addresses struct {
		Prefecture []Item `yaml:"prefecture"`
		Town       []Item `yaml:"town"`
		City       []Item `yaml:"city"`
	} `yaml:"addresses"`
}

type Address struct {
	Prefecture Item
	Town       Item
	City       Item
}

func loadAddresses() {
	rp := "src/github.com/mattn/go-gimei/data/addresses.yml"
	for _, p := range filepath.SplitList(os.Getenv("GOPATH")) {
		f := filepath.Join(p, rp)
		if _, err := os.Stat(f); err == nil {
			if b, err := ioutil.ReadFile(f); err == nil {
				if err = yaml.Unmarshal(b, &addresses); err == nil {
					return
				}
			}
		}
	}
	panic("failed to load names data")
}

func (a *Address) String() string {
	return a.Kanji()
}

func (a *Address) Kanji() string {
	return a.Prefecture.Kanji() + a.Town.Kanji() + a.City.Kanji()
}

func (a *Address) Hiragana() string {
	return a.Prefecture.Hiragana() + a.Town.Hiragana() + a.City.Hiragana()
}

func (a *Address) Katakana() string {
	return a.Prefecture.Katakana() + a.Town.Katakana() + a.City.Katakana()
}

func NewAddress() *Address {
	return &Address{
		Prefecture: NewPrefecture(),
		Town:       NewTown(),
		City:       NewCity(),
	}
}

func NewPrefecture() Item {
	onceAddress.Do(loadAddresses)
	return addresses.Addresses.Prefecture[r.Int()%len(addresses.Addresses.Prefecture)]
}

func NewTown() Item {
	onceAddress.Do(loadAddresses)
	return addresses.Addresses.Town[r.Int()%len(addresses.Addresses.Town)]
}

func NewCity() Item {
	onceAddress.Do(loadAddresses)
	return addresses.Addresses.City[r.Int()%len(addresses.Addresses.City)]
}
