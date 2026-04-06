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

## Contributing
Contributions are very welcomed, however please follow the below guidelines.

- Feel free to open an issue describing the bug or enhancement so it can be
discussed.  
- Try to maintain the naming integrity present in this project.  
- Create a Pull Request with your changes on the main branch.
