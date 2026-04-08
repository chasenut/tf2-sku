# Parser for TF2 Stock Keeping Unit (sku)

[![Go Reference](https://pkg.go.dev/badge/github.com/chasenut/tf2-sku.svg)](https://pkg.go.dev/github.com/chasenut/tf2-sku)

Go (Golang) package for parsing TF2 `sku`, commonly found in trading websites, APIs and bots.

If you want to support me and my projects, here is my [Steam Trade Link](https://steamcommunity.com/tradeoffer/new/?partner=1487050196&token=2kpoTi1k).
You can also support me by joining my [Discord Server](https://discord.gg/YjXvX7WfzS).

## Getting started

### Installing

This assumes you already have a working Go environment. 
To pull the _the latest_ version from the main branch, use:
```
go get github.com/chasenut/tf2-sku
```

### Usage

Import the package into your project.
```
import "github.com/chasenut/tf2-sku"
```

To parse `sku` string to `*Item`, use the following:
```
item, err := tf2sku.FromSKU("1071;11")
```

To parse `*Item` to `sku` string, use the following:
```
sku := tf2sku.ToSKU(myItemPointer)
```

### Examples

```
import "github.com/chasenut/tf2-sku"

// SKU of a Mann Co. Supply Crate Key - 5021 is the defindex, 6 is the quality
var sku = "5021;6"

// Converts the sku string into an *Item
var item = tf2sku.FromSKU(sku)
/* ->
&tf2sku.Item{
	item.Defindex		= 5021
	item.Quality		= 6
	item.Craftable		= true
	item.Tradable 		= true
	item.Australium 	= false
	item.Festivized 	= false
    item.Quality2		= -1    // nil
	item.Particle		= -1    // nil
	item.Skin			= -1    // nil
	item.KillstreakTier	= 0     // nil
	item.Wear			= -1    // nil
	item.TargetDefindex	= -1    // nil
	item.OutputDefindex	= -1    // nil
	item.OutputQuality	= -1    // nil
	item.CraftNumber	= -1    // nil
	item.CrateNumber	= -1    // nil
	item.Paint			= -1    // nil
	item.Sheen			= -1    // nil
	item.Killstreaker	= -1    // nil
}
*/
```

```
// Mann Co. Supply Crate Key
var item = tf2sku.NewItem()
/* ->
&tf2sku.Item{
	item.Defindex		= 5021
	item.Quality		= 6
	item.Craftable		= true
	item.Tradable 		= true
	item.Australium 	= false
	item.Festivized 	= false
    item.Quality2		= -1    // nil
	item.Particle		= -1    // nil
	item.Skin			= -1    // nil
	item.KillstreakTier	= 0     // nil
	item.Wear			= -1    // nil
	item.TargetDefindex	= -1    // nil
	item.OutputDefindex	= -1    // nil
	item.OutputQuality	= -1    // nil
	item.CraftNumber	= -1    // nil
	item.CrateNumber	= -1    // nil
	item.Paint			= -1    // nil
	item.Sheen			= -1    // nil
	item.Killstreaker	= -1    // nil
}

// Converts the *Item into an sku string
var sku = tf2sku.ToSKU(&Item)
// -> "5021;6"
```


## Contributing
Contributions are very welcomed, however please follow the below guidelines.

- Feel free to open an issue describing the bug or enhancement so it can be
discussed.  
- Try to maintain the naming integrity present in this project.  
- Create a Pull Request with your changes on the main branch.
