// tf2-sku - Golang parser for Team Fortress 2 Stock Keeping Units (SKU)

// Package tf2sku provides simple Team Fortress 2 SKU parsing. 
// Helpful when SKU is required for some endpoints, and when
// you need to retrieve most important information out of a response.
package tf2sku

import (
	"fmt"
	"strconv"
	"strings"
)

const QualityStrange = 11

// Use NewItem function to get this.
type Item struct {
	Defindex		int
	Quality			int
	Craftable		bool
	Tradable		bool
	Australium		bool
	Festivized		bool
	Quality2		int
	Particle		int
	Skin			int
	KillstreakTier	int
	Wear			int
	TargetDefindex 	int
	OutputDefindex	int
	OutputQuality	int
	CraftNumber		int
	CrateNumber		int
	Paint			int
	Sheen			int
	Killstreaker	int
}

// Returns new *Item struct
//
// Values < 0 means not set.
//
// Sets the default values:
//	- item.Defindex			= 0
//	- item.Quality			= 0
//	- item.Craftable		= true
//	- item.Tradable 		= true
//	- item.Australium 		= false
//	- item.Festivized 		= false
//  - item.Quality2			= -1
//	- item.Particle			= -1
//	- item.Skin				= -1
//	- item.KillstreakTier	= 0 
//	- item.Wear				= -1
//	- item.TargetDefindex 	= -1
//	- item.OutputDefindex	= -1
//	- item.OutputQuality	= -1
//	- item.CraftNumber		= -1
//	- item.CrateNumber		= -1
//	- item.Paint			= -1
//	- item.Sheen			= -1
//	- item.Killstreaker		= -1
func NewItem() *Item {
	item := &Item{}

	item.Defindex		= 0
	item.Quality		= 0
	item.Craftable		= true
	item.Tradable 		= true
	item.Australium 	= false
	item.Festivized 	= false
	item.Quality2		= -1
	item.Particle		= -1
	item.Skin			= -1
	item.KillstreakTier	= 0 
	item.Wear			= -1
	item.TargetDefindex = -1
	item.OutputDefindex	= -1
	item.OutputQuality	= -1
	item.CraftNumber	= -1
	item.CrateNumber	= -1
	item.Paint			= -1
	item.Sheen			= -1
	item.Killstreaker	= -1

	return item
}

// SKU format: "defindex;quality;attributes" where attributes are optional.
//
// All attributes are separated with semicolon ';'.
//
// Returns *Item, uses NewItem() internally (has default fields, see: type Item)
//
// Returns error if sku lacks defindex or quality.
//
// Returns error if attribute format is invalid.
func FromSKU(sku string) (*Item, error) {
	item := NewItem()
	words := strings.Split(sku, ";")
	if len(words) < 2 {
		return &Item{}, fmt.Errorf("invalid sku, not enough fields: %s", sku)
	}

	defindex, err := strconv.Atoi(words[0])
	if err != nil {
		return &Item{}, fmt.Errorf("invalid sku, couldn't parse defindex: %s", sku)
	}
	item.Defindex = defindex
	quality, err := strconv.Atoi(words[1])
	if err != nil {
		return &Item{}, fmt.Errorf("invalid sku, couldn't parse quality: %s", sku)
	}
	item.Quality = quality

	if len(words) < 3 {
		return item, nil
	}

	attributes := words[2:]
	for _, a := range attributes {
		fmt.Println(a)
		switch a {
		case "uncraftable":
			item.Craftable = false
			continue
		case "untradeable":
			item.Tradable = false
			continue
		case "untradable":
			item.Tradable = false
			continue
		case "australium":
			item.Australium = true
			continue
		case "strange":
			item.Quality2 = QualityStrange
			continue
		case "festive":
			item.Festivized = true
			continue
		}
		if v, ok := strings.CutPrefix(a, "u"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.Particle = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "pk"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.Skin = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "kt-"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.KillstreakTier = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "w"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.Wear = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "td-"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.TargetDefindex = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "od-"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.OutputDefindex = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "oq-"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.OutputQuality = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "n"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.CraftNumber = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "c"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.CrateNumber = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "p"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.Paint = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "ks-"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.Sheen = n
			continue
		}
		if v, ok := strings.CutPrefix(a, "ke-"); ok {
			n, err := strconv.Atoi(v)
			if err != nil {
				return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
			}
			item.Killstreaker = n
			continue
		}
		if a == "" {
			continue
		}
		return &Item{}, fmt.Errorf("invalid attribute '%v'", a)
	}
	return item, nil
}

// Does not validate values.
//
// At least Item.Defindex and Item.Quality should be set.
//
// Ignores all values < 0.
//
// Returns string (SKU).
func ToSKU(i *Item) string {
	sku := fmt.Sprintf("%v;%v", i.Defindex, i.Quality)

	if i.Craftable == false {
		sku += ";uncraftable"
	}
	if i.Tradable == false {
		sku += ";untradable"
	}
	if i.Australium == true {
		sku += ";australium"
	}
	if i.Quality2 == QualityStrange {
		sku += ";strange"
	}
	if i.Festivized == true {
		sku += ";festive"
	}

	if i.Particle > -1 {
		sku += fmt.Sprintf(";u%v", i.Particle)
	}
	if i.Skin > -1 {
		sku += fmt.Sprintf(";pk%v", i.Skin)
	}
	if i.KillstreakTier > 0 {
		sku += fmt.Sprintf(";kt-%v", i.KillstreakTier)
	}
	if i.Wear > -1 {
		sku += fmt.Sprintf(";w%v", i.Wear)
	}
	if i.TargetDefindex > -1 {
		sku += fmt.Sprintf(";td-%v", i.TargetDefindex)
	}
	if i.OutputDefindex > -1 {
		sku += fmt.Sprintf(";od-%v", i.OutputDefindex)
	}
	if i.OutputQuality > -1 {
		sku += fmt.Sprintf(";oq-%v", i.OutputQuality)
	}
	if i.CraftNumber > -1 {
		sku += fmt.Sprintf(";n%v", i.CraftNumber)
	}
	if i.CrateNumber > -1 {
		sku += fmt.Sprintf(";c%v", i.CrateNumber)
	}
	if i.Paint > -1 {
		sku += fmt.Sprintf(";p%v", i.Paint)
	}
	if i.Sheen > -1 {
		sku += fmt.Sprintf(";ks-%v", i.Sheen)
	}
	if i.Killstreaker > -1 {
		sku += fmt.Sprintf(";ke-%v", i.Killstreaker)
	}

	return sku
}
