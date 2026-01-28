package gimei

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var (
	//go:embed data/addresses.yml data/names.yml data/postalcodes.yml
	assets embed.FS

	names       name
	addresses   address
	postalCodes postalCode
	onceName    sync.Once
	onceAddress sync.Once
	oncePostal  sync.Once
	r           *rand.Rand
	mu          sync.Mutex
)

// Item take four figure for japanese. Kanji/Hiragana/Katakana/Romaji.
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

// Katakana return string of Item as katakana.
func (i Item) Katakana() string {
	return i[2]
}

// Romaji return string of Item as romaji.
func (i Item) Romaji() string {
	if len(i) <= 3 {
		return ""
	}
	return cases.Title(language.Und, cases.NoLower).String(i[3])
}

// Sex store Male or Female.
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

// list of sex
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

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// SetRandom set a pointer to rand.Rand that uses to generate random values.
func SetRandom(rnd *rand.Rand) {
	r = rnd
}

func loadNames() {
	if b, err := assets.ReadFile("data/names.yml"); err == nil {
		if err = yaml.Unmarshal(b, &names); err == nil {
			return
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

// Katakana return string of Name as katakana.
func (n *Name) Katakana() string {
	return n.Last.Katakana() + " " + n.First.Katakana()
}

// Romaji return string of Name as romaji.
func (n *Name) Romaji() string {
	return n.First.Romaji() + " " + n.Last.Romaji()
}

// IsMale return true if he is male.
func (n *Name) IsMale() bool {
	return n.Sex == Male
}

// IsFemale return true if she is female.
func (n *Name) IsFemale() bool {
	return n.Sex == Female
}

// NewName return new instance of person.
func NewName() *Name {
	mu.Lock()
	defer mu.Unlock()

	onceName.Do(loadNames)
	if r.Int()%2 == 0 {
		return &Name{
			First: names.FirstName.Male[r.Int()%len(names.FirstName.Male)],
			Last:  names.LastName[r.Int()%len(names.LastName)],
			Sex:   Male,
		}
	} else {
		return &Name{
			First: names.FirstName.Female[r.Int()%len(names.FirstName.Female)],
			Last:  names.LastName[r.Int()%len(names.LastName)],
			Sex:   Female,
		}
	}
}

// NewDog return new instance of person whose last name begins "inu".
func NewDog() *Name {
	mu.Lock()
	defer mu.Unlock()

	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Animal[r.Int()%len(names.FirstName.Animal)],
		Last:  names.LastNameDog[r.Int()%len(names.LastNameDog)],
	}
}

// NewCat return new instance of person whose last name begins "neko".
func NewCat() *Name {
	mu.Lock()
	defer mu.Unlock()

	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Animal[r.Int()%len(names.FirstName.Animal)],
		Last:  names.LastNameCat[r.Int()%len(names.LastNameCat)],
	}
}

// NewMale return new instance of person that is male.
func NewMale() *Name {
	mu.Lock()
	defer mu.Unlock()

	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Male[r.Int()%len(names.FirstName.Male)],
		Last:  names.LastName[r.Int()%len(names.LastName)],
		Sex:   Male,
	}
}

// NewFemale return new instance of person that is female.
func NewFemale() *Name {
	mu.Lock()
	defer mu.Unlock()

	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Female[r.Int()%len(names.FirstName.Female)],
		Last:  names.LastName[r.Int()%len(names.LastName)],
		Sex:   Female,
	}
}

// NewMaleDog return new instance of male person whose last name begins "inu".
func NewMaleDog() *Name {
	mu.Lock()
	defer mu.Unlock()

	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Male[r.Int()%len(names.FirstName.Male)],
		Last:  names.LastNameDog[r.Int()%len(names.LastNameDog)],
		Sex:   Male,
	}
}

// NewFemaleDog return new instance of female person whose last name begins "inu".
func NewFemaleDog() *Name {
	mu.Lock()
	defer mu.Unlock()

	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Female[r.Int()%len(names.FirstName.Female)],
		Last:  names.LastNameDog[r.Int()%len(names.LastNameDog)],
		Sex:   Female,
	}
}

// NewMaleCat return new instance of male person whose last name begins "neko".
func NewMaleCat() *Name {
	mu.Lock()
	defer mu.Unlock()

	onceName.Do(loadNames)
	return &Name{
		First: names.FirstName.Male[r.Int()%len(names.FirstName.Male)],
		Last:  names.LastNameCat[r.Int()%len(names.LastNameCat)],
		Sex:   Male,
	}
}

// NewFemaleCat return new instance of female person whose last name begins "neko".
func NewFemaleCat() *Name {
	mu.Lock()
	defer mu.Unlock()

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
	if i == 3 { // by romaji
		token[0], token[1] = strings.ToLower(token[1]), strings.ToLower(token[0])
	}
	for _, last := range names.LastName {
		if last[i] != token[0] {
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

// FindNameByKanji find Name by kanji.
func FindNameByKanji(kanji string) *Name {
	return findNameByIndex(kanji, 0)
}

// FindNameByHiragana find Name by hiragana.
func FindNameByHiragana(hiragana string) *Name {
	return findNameByIndex(hiragana, 1)
}

// FindNameByKatakana find Name by katakana.
func FindNameByKatakana(katakana string) *Name {
	return findNameByIndex(katakana, 2)
}

// FindNameByRomaji find Name by romaji.
func FindNameByRomaji(romaji string) *Name {
	return findNameByIndex(romaji, 3)
}

// address store data sturecture just same as addresses.yml.
type address struct {
	Addresses struct {
		Prefecture []Item `yaml:"prefecture"`
		City       []Item `yaml:"city"`
		Town       []Item `yaml:"town"`
	} `yaml:"addresses"`
}

// postalCode store data structure for postal codes
type postalCode struct {
	PostalCodes []Item `yaml:"postal_codes"`
}

// Address store address that is pointed by prefecture/city/town.
type Address struct {
	Prefecture Item
	City       Item
	Town       Item
}

func loadAddresses() {
	if b, err := assets.ReadFile("data/addresses.yml"); err == nil {
		if err = yaml.Unmarshal(b, &addresses); err == nil {
			return
		}
	}
	panic("failed to load addresses data")
}

func loadPostalCodes() {
	if b, err := assets.ReadFile("data/postalcodes.yml"); err == nil {
		if err = yaml.Unmarshal(b, &postalCodes); err == nil {
			return
		}
	}
	panic("failed to load postal codes data")
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

// Katakana return string of Address as katakana.
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
	mu.Lock()
	defer mu.Unlock()

	onceAddress.Do(loadAddresses)
	return addresses.Addresses.Prefecture[r.Int()%len(addresses.Addresses.Prefecture)]
}

// NewTown return new instance of town.
func NewTown() Item {
	mu.Lock()
	defer mu.Unlock()

	onceAddress.Do(loadAddresses)
	return addresses.Addresses.Town[r.Int()%len(addresses.Addresses.Town)]
}

// NewCity return new instance of city.
func NewCity() Item {
	mu.Lock()
	defer mu.Unlock()

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

// FindAddressByKanji find Address by kanji.
func FindAddressByKanji(kanji string) *Address {
	return findAddressByIndex(kanji, 0)
}

// FindAddressByHiragana find Address by hiragana.
func FindAddressByHiragana(hiragana string) *Address {
	return findAddressByIndex(hiragana, 1)
}

// FindAddressByKatakana find Address by katakana.
func FindAddressByKatakana(katakana string) *Address {
	return findAddressByIndex(katakana, 2)
}

// PostalCode store postal code
type PostalCode struct {
	Code Item
}

// String implement Stringer.
func (p *PostalCode) String() string {
	return p.Kanji()
}

// Kanji return string of PostalCode as kanji.
func (p *PostalCode) Kanji() string {
	return p.Code.Kanji()
}

// NewPostalCode return new instance of postal code.
func NewPostalCode() *PostalCode {
	mu.Lock()
	defer mu.Unlock()

	oncePostal.Do(loadPostalCodes)
	return &PostalCode{
		Code: postalCodes.PostalCodes[r.Int()%len(postalCodes.PostalCodes)],
	}
}

func CountData() string {
	onceName.Do(loadNames)
	onceAddress.Do(loadAddresses)
	oncePostal.Do(loadPostalCodes)

	var addr = &addresses.Addresses
	return fmt.Sprintf(`FirstName:
  Male:       %5d 
  Female:     %5d
LastName:     %5d
Adresses:
  Prefecture: %5d
  City:       %5d
  Town:       %5d
PostalCodes: %5d`,
		len(names.FirstName.Male), len(names.FirstName.Female), len(names.LastName),
		len(addr.Prefecture), len(addr.City), len(addr.Town),
		len(postalCodes.PostalCodes))
}
