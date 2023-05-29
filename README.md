# regionlang - Country/Region to Language Mapper

`regionlang` is a Go library that aids in determining the language associated with a given country or region. It comes in handy when you need to localize messages for users whose language preference is unknown.

## Install

```bash
go get github.com/modernice/go-regionlang
```

## Usage

The library's primary feature is the `Find` function which takes a region code as an argument and returns the corresponding language.

Here's a basic example:

```go
package main

import (
	"fmt"

	"github.com/modernice/go-regionlang"
	"golang.org/x/text/language"
)

func main() {
	region := "be" // Belgium
	base, conf := regionlang.Find(region)

	fmt.Println(base.String()) // Output: "fr" (for French)
	fmt.Println(conf == language.Exact) // Output: true
}
```

In the above example, `regionlang.Find("be")` returns French ("fr") as the most likely language for Belgium.

### Custom Language Tags

By default, `go-regionlang` matches against all built-in language tags. However, your application might not support every single built-in language. To specify which language tags to match against, pass the allowed tags to the `Find` function as follows:

```go
package main

import (
	"fmt"

	"github.com/modernice/go-regionlang"
	"golang.org/x/text/language"
)

func main() {
	allowedTags := []language.Tag{language.English, language.Spanish, language.French}
	base, conf := regionlang.Find("some-region-code", allowedTags...)

	fmt.Println(base.String())
	fmt.Println(conf == language.Exact)
}
```

In this example, `Find` will only consider English, Spanish, and French when determining the language for the given region.

## License

This project is licensed under the [MIT](./LICENSE) License. 
