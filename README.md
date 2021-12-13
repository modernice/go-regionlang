# go-regionlang - Detect languages from countries / regions

```sh
go get github.com/modernice/go-regionlang
```

## Use Case

Sometimes you need to find the appropriate language for a country code in order
to localize messages for a user with an unknown language. `go-regionlang` does
exactly that, relying purely on the `golang.org/x/text/language` package:

```go
package example

func getLanguageForRegion() {
	region := "be" // Belgium
	base, conf := regionlang.Find(region)

	base.String() == "fr" // French
	conf == language.Exact
}
```

## Custom Language Tags

By default, `go-regionlang` matches against all built-in language tags. Your
application most probably does not support every single built-in language. You
can specify which language tags to match against:

```go
package example

func getLanguageForRegion(allowedTags []language.Tag) {
	base, conf := regionlang.Find("some-region-code", allowedTags...)
}
```

## License

[MIT](./LICENSE)
