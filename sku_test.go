package tf2sku_test

import (
	"fmt"
	"testing"

	tf2sku "github.com/chasenut/tf2-sku"
)

func TestFromSKU(t *testing.T) {
	num1 := 0
	num2 := 21
	num3 := 321
	num4 := 42
	item1 := tf2sku.NewItem()
	item1.Defindex = num1
	item1.Quality = num2
	item3 := tf2sku.NewItem()
	item3.Defindex = num3
	item3.Quality =	num1
	item3.Australium = true
	item3.Craftable = false
	item3.Tradable = false
	item3.Quality2 = 11 // QualityStrange
	item3.Festivized = true
	item3.Particle = num4
	item3.Skin = num3
	item3.KillstreakTier = num1
	item3.Wear = num2
	item3.TargetDefindex = num2
	item3.OutputDefindex = num2
	item3.OutputQuality = num2
	item3.CraftNumber =	num2
	item3.CrateNumber =	num2
	item3.Paint = num2
	item3.Sheen = num2
	item3.Killstreaker = num2
	testCases := []struct{
		sku			string
		expected	tf2sku.Item
		err			bool
	} {
		{
			sku:		"",
			expected: 	tf2sku.Item{},
			err:		true,
		},
		{
			sku:		"",
			expected: 	tf2sku.Item{},
			err:		true,
		},
		{
			sku:		fmt.Sprintf("%v:%v", num1, num3),
			expected: 	tf2sku.Item{},
			err:		true,
		},
		{
			sku:		fmt.Sprintf("%v;%v", num1, num2),
			expected: 	*item1,
			err:		false,
		},
		{
			sku:		fmt.Sprintf("%v;%v;uncraftable;untradable;australium;strange;festive;u%v;pk%v;kt-%v;w%v;td-%v;od-%v;oq-%v;n%v;c%v;p%v;ks-%v;ke-%v", 
				num3, num1, num4, num3, num1, num2, num2, num2, num2, num2, num2, num2, num2, num2),
			expected: 	*item3,
			err:		false,
		},
	}
	for _, c := range testCases {
		result, err := tf2sku.FromSKU(c.sku)
		errored := err != nil
		if *result != c.expected || errored != c.err {
			t.Errorf("Error FromSKU(), wanterr: %v, errored: %v\nexpected: \n\r'%+v', \ngot: \n\r'%+v'\nSku:%s\n", 
				c.err, 
				errored, 
				c.expected, 
				result, 
				c.sku,
			)
		}
	}
}

func TestToSKU(t *testing.T) {
	num1 := 1
	num2 := 34
	num3 := 423
	num4 := 43
	item1 := tf2sku.NewItem()
	item2 := tf2sku.NewItem()
	item3 := tf2sku.NewItem()
	item2.Defindex = num1
	item2.Quality = num2
	item3.Defindex = num3
	item3.Quality =	num1
	item3.Australium = true
	item3.Craftable = false
	item3.Tradable = false
	item3.Quality2 = 11 // QualityStrange
	item3.Festivized = true
	item3.Particle = num4
	item3.Skin = num3
	item3.KillstreakTier = num1
	item3.Wear = num2
	item3.TargetDefindex = num2
	item3.OutputDefindex = num2
	item3.OutputQuality = num2
	item3.CraftNumber =	num2
	item3.CrateNumber =	num2
	item3.Paint = num2
	item3.Sheen = num2
	item3.Killstreaker = num2
	testCases := []struct{
		i 			tf2sku.Item
		expected	string
		err			bool
	} {
		{
			i: 			*item1,
			expected: 	"0;0",
		},
		{
			i: 			*item2,
			expected: 	fmt.Sprintf("%v;%v", num1, num2),
		},
		{
			i: 			*item3,
			expected: 	fmt.Sprintf("%v;%v;uncraftable;untradable;australium;strange;festive;u%v;pk%v;kt-%v;w%v;td-%v;od-%v;oq-%v;n%v;c%v;p%v;ks-%v;ke-%v", 
				num3, num1, num4, num3, num1, num2, num2, num2, num2, num2, num2, num2, num2, num2),
		},
		
	}
	for _, c := range testCases {
		result := tf2sku.ToSKU(&c.i)
		if result != c.expected {
			t.Errorf("Error ToSKU(), \nexpected: \n\r'%s', \ngot: \n\r'%s'\nItem:%+v\n", 
				c.expected, 
				result, 
				c.i,
			)
		}
	}
}
