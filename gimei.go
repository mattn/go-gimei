package gimei

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
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
	LastName []Item `yaml:"last_name"`
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

func findNameByIndex(n string, i int) *Name {
	onceName.Do(loadNames)
	token := strings.SplitN(n, " ", 2)
	if len(token) != 2 {
		return nil
	}
	for _, last := range names.LastName {
		if last.Kanji() != token[0] {
			continue
		}
		for _, first := range names.FirstName.Male {
			if first[i] != token[1] {
				continue
			}
			return &Name{
				First: first,
				Last:  last,
				Sex:   Male,
			}
		}
		for _, first := range names.FirstName.Female {
			if first[i] != token[1] {
				continue
			}
			return &Name{
				First: first,
				Last:  last,
				Sex:   Female,
			}
		}
	}
	return nil
}

func FindNameByKanji(name string) *Name {
	return findNameByIndex(name, 0)
}

func FindNameByHiragana(name string) *Name {
	return findNameByIndex(name, 1)
}

func FindNameByKatakana(name string) *Name {
	return findNameByIndex(name, 2)
}

type address struct {
	Addresses struct {
		Prefecture []Item `yaml:"prefecture"`
		City       []Item `yaml:"city"`
		Town       []Item `yaml:"town"`
	} `yaml:"addresses"`
}

type Address struct {
	Prefecture Item
	City       Item
	Town       Item
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
	return a.Prefecture.Kanji() + a.City.Kanji() + a.Town.Kanji()
}

func (a *Address) Hiragana() string {
	return a.Prefecture.Hiragana() + a.City.Hiragana() + a.Town.Hiragana()
}

func (a *Address) Katakana() string {
	return a.Prefecture.Katakana() + a.City.Katakana() + a.Town.Katakana()
}

func NewAddress() *Address {
	return &Address{
		Prefecture: NewPrefecture(),
		City:       NewCity(),
		Town:       NewTown(),
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

func findAddressByIndex(a string, i int) *Address {
	onceAddress.Do(loadAddresses)
	for _, prefecture := range addresses.Addresses.Prefecture {
		if !strings.HasPrefix(a, prefecture[i]) {
			continue
		}
		for _, city := range addresses.Addresses.City {
			if !strings.HasPrefix(a, prefecture[i]+city[i]) {
				continue
			}
			for _, town := range addresses.Addresses.Town {
				if a != prefecture[i]+city[i]+town[i] {
					continue
				}
				println((&Address{
					Prefecture: prefecture,
					City:       city,
					Town:       town,
				}).String())
				return &Address{
					Prefecture: prefecture,
					City:       city,
					Town:       town,
				}
			}
		}
	}
	return nil
}

func FindAddressByKanji(address string) *Address {
	return findAddressByIndex(address, 0)
}

func FindAddressByHiragana(address string) *Address {
	return findAddressByIndex(address, 1)
}

func FindAddressByKatakana(address string) *Address {
	return findAddressByIndex(address, 2)
}
