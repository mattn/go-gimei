package gimei

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

//go:generate go-assets-builder -p gimei -o assets.go data

var (
	names       name
	addresses   address
	onceName    sync.Once
	onceAddress sync.Once
	r           *rand.Rand
)

// Item take three figure for japanese. Kanji/Hiragana/Katakana.
// It's not only the difference of sounds, it just letters.
type Item []string

// String implement Stringer
func (i Item) String() string {
	return i.Kanji()
}

// Kanji return string of Item as kanji.
func (i Item) Kanji() string {
	return i[0]
}

// Hiragana return string of Item as hiragana.
func (i Item) Hiragana() string {
	return i[1]
}

// Hiragana return string of Item as katakana.
func (i Item) Katakana() string {
	return i[2]
}

// Sex store Male ore Female.
type Sex int

// String implement Stringer.
func (s Sex) String() string {
	switch s {
	case Male:
		return "男"
	case Female:
		return "女"
	}
	return "？"
}

const (
	Male   Sex = iota + 1 // 男
	Female                // 女
)

// name store data sturecture just same as names.yml.
type name struct {
	FirstName struct {
		Male   []Item `yaml:"male"`
		Female []Item `yaml:"female"`
		Animal []Item `yaml:"animal"`
	} `yaml:"first_name"`
	LastName    []Item `yaml:"last_name"`
	LastNameDog []Item `yaml:"last_name_dog"`
	LastNameCat []Item `yaml:"last_name_cat"`
}

// Name store name and sex for a person.
type Name struct {
	First Item
	Last  Item
	Sex   Sex
}

func loadNames() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	if f, err := Assets.Open("/data/names.yml"); err == nil {
		if b, err := ioutil.ReadAll(f); err == nil {
			if err = yaml.Unmarshal(b, &names); err == nil {
				return
			}
		}
	}
	panic("failed to load names data")
}

// String implement Stringer.
func (n *Name) String() string {
	return n.Kanji()
}

// Kanji return string of Name as kanji.
func (n *Name) Kanji() string {
	return n.Last.Kanji() + " " + n.First.Kanji()
}

// Hiragana return string of Name as hiragana.
func (n *Name) Hiragana() string {
	return n.Last.Hiragana() + " " + n.First.Hiragana()
}

// Hiragana return string of Name as katakana.
func (n *Name) Katakana() string {
	return n.Last.Katakana() + " " + n.First.Katakana()
}

// IsMale return true if he is male.
func (n *Name) IsMale() bool {
	return n.Sex == Male
}

// IsMale return true if she is female.
func (n *Name) IsFemale() bool {
	return n.Sex == Female
}

// NewName return new instance of person.
func NewName() *Name {
	onceName.Do(loadNames)
	if r.Int()%2 == 0 {
		return NewMale()
	} else {
		return NewFemale()
	}
}

// NewName return new instance of person.
func NewDog() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Animal[r.Int()%len(names.FirstName.Animal)],
		Last:  names.LastNameDog[r.Int()%len(names.LastNameDog)],
	}
}

func NewCat() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Animal[r.Int()%len(names.FirstName.Animal)],
		Last:  names.LastNameCat[r.Int()%len(names.LastNameCat)],
	}
}

// NewMale return new instance of person that is male.
func NewMale() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Male[r.Int()%len(names.FirstName.Male)],
		Last:  names.LastName[r.Int()%len(names.LastName)],
		Sex:   Male,
	}
}

// NewFemale return new instance of person that is female.
func NewFemale() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Female[r.Int()%len(names.FirstName.Female)],
		Last:  names.LastName[r.Int()%len(names.LastName)],
		Sex:   Female,
	}
}

// NewMaleDog return new instance of person that is dog.
func NewMaleDog() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Male[r.Int()%len(names.FirstName.Male)],
		Last:  names.LastNameDog[r.Int()%len(names.LastNameDog)],
		Sex:   Male,
	}
}

// NewFemaleCat return new instance of person that is dog.
func NewFemaleDog() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Female[r.Int()%len(names.FirstName.Female)],
		Last:  names.LastNameDog[r.Int()%len(names.LastNameDog)],
		Sex:   Female,
	}
}

// NewMaleCat return new instance of person that is dog.
func NewMaleCat() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Male[r.Int()%len(names.FirstName.Male)],
		Last:  names.LastNameCat[r.Int()%len(names.LastNameCat)],
		Sex:   Male,
	}
}

// NewFemaleCat return new instance of person that is dog.
func NewFemaleCat() *Name {
	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Female[r.Int()%len(names.FirstName.Female)],
		Last:  names.LastNameCat[r.Int()%len(names.LastNameCat)],
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

// FindNameByKanji find Name from kanji.
func FindNameByKanji(kanji string) *Name {
	return findNameByIndex(kanji, 0)
}

// FindNameByHiragana find Name from hiragana.
func FindNameByHiragana(hiragana string) *Name {
	return findNameByIndex(hiragana, 1)
}

// FindNameByKanji find Name from katakana.
func FindNameByKatakana(katakana string) *Name {
	return findNameByIndex(katakana, 2)
}

// address store data sturecture just same as addresses.yml.
type address struct {
	Addresses struct {
		Prefecture []Item `yaml:"prefecture"`
		City       []Item `yaml:"city"`
		Town       []Item `yaml:"town"`
	} `yaml:"addresses"`
}

// Address store address that is pointed by prefecture/city/town.
type Address struct {
	Prefecture Item
	City       Item
	Town       Item
}

func loadAddresses() {
	if f, err := Assets.Open("/data/addresses.yml"); err == nil {
		if b, err := ioutil.ReadAll(f); err == nil {
			if err = yaml.Unmarshal(b, &addresses); err == nil {
				return
			}
		}
	}
	panic("failed to load names data")
}

// String implement Stringer.
func (a *Address) String() string {
	return a.Kanji()
}

// Kanji return string of Address as kanji.
func (a *Address) Kanji() string {
	return a.Prefecture.Kanji() + a.City.Kanji() + a.Town.Kanji()
}

// Hiragana return string of Address as hiragana.
func (a *Address) Hiragana() string {
	return a.Prefecture.Hiragana() + a.City.Hiragana() + a.Town.Hiragana()
}

// Hiragana return string of Address as katakana.
func (a *Address) Katakana() string {
	return a.Prefecture.Katakana() + a.City.Katakana() + a.Town.Katakana()
}

// NewAddress return new instance of address.
func NewAddress() *Address {
	return &Address{
		Prefecture: NewPrefecture(),
		City:       NewCity(),
		Town:       NewTown(),
	}
}

// NewPrefecture return new instance of prefecture.
func NewPrefecture() Item {
	onceAddress.Do(loadAddresses)
	return addresses.Addresses.Prefecture[r.Int()%len(addresses.Addresses.Prefecture)]
}

// NewTown return new instance of town.
func NewTown() Item {
	onceAddress.Do(loadAddresses)
	return addresses.Addresses.Town[r.Int()%len(addresses.Addresses.Town)]
}

// NewCity return new instance of city.
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

// FindAddressByKanji find Address from kanji.
func FindAddressByKanji(kanji string) *Address {
	return findAddressByIndex(kanji, 0)
}

// FindAddressByHiragana find Address from hiragana.
func FindAddressByHiragana(hiragana string) *Address {
	return findAddressByIndex(hiragana, 1)
}

// FindAddressByKanji find Address from katakana.
func FindAddressByKatakana(katakana string) *Address {
	return findAddressByIndex(katakana, 2)
}
