# Loafer-i18n

Loafer-i18n is a Go package and a command that helps you translate Go programs into multiple languages.

- Supports pluralized strings with =x or >x expression.
- Supports strings with similar to [fmt.Sprintf](https://golang.org/pkg/fmt/) format syntax.
- Supports message files of any format (e.g. JSON, TOML, YAML).

# Package i18n

The i18n package provides support for looking up messages according to a set of locale preferences.

```go
package main

import (
	"fmt"
	"os"

	_ "github.com/EasyGram-io/loafer-i18n/example/catalog"
	"github.com/EasyGram-io/loafer-i18n/i18n"
	"golang.org/x/text/language"
)

func main() {

	i18n.SetLang(language.SimplifiedChinese)

	i18n.Printf(`hello world!`)
	fmt.Println()

	name := `EasyGram.io`

	i18n.Printf(`hello %s!`, name)
	fmt.Println()

	i18n.Printf(`%s has %d cat.`, name, 1)
	fmt.Println()

	i18n.Printf(`%s has %d cat.`, name, 2, i18n.Plural(
		`%[2]d=1`, `%s has %d cat.`,
		`%[2]d>1`, `%s has %d cats.`,
	))
	fmt.Println()

	i18n.Fprintf(os.Stderr, `%s have %d apple.`, name, 2, i18n.Plural(
		`%[2]d=1`, `%s have an apple.`,
		`%[2]d=2`, `%s have two apples.`,
		`%[2]d>2`, `%s have %d apples.`,
	))
	fmt.Println()
}
```

# Command loaferi18n

The loaferi18n command manages message files used by the i18n package.

```
go get -u github.com/EasyGram-io/loafer-i18n/loaferi18n
loaferi18n -h

  update, u    merge translations and generate catalog
  extract, e   extracts strings to be translated from code
  generate, g  generates code to insert translated messages
```

### Extracting messages

Use `loaferi18n extract . ./locales/en.json` to extract all i18n.Sprintf function literals in Go source files to a message file for translation.

`./locales/en.json`
```json
{
  "%s has %d cat.": "%s has %d cat.",
  "%s has %d cats.": "%s has %d cats.",
  "%s have %d apples.": "%s have %d apples.",
  "%s have an apple.": "%s have an apple.",
  "%s have two apples.": "%s have two apples.",
  "hello %s!": "hello %s!",
  "hello world!": "hello world!"
}
```

### Translating a new language

1. Create an empty message file for the language that you want to add (e.g. `zh-Hans.json`).
2. Run `loaferi18n update ./locales/en.json ./locales/zh-Hans.json` to populate `zh-Hans.json` with the mesages to be translated.

	`./locales/zh-Hans.json`
	```json
	{
	  "%s has %d cat.": "%s有%d只猫。",
	  "%s has %d cats.": "%s有%d只猫。",
	  "%s have %d apples.": "%s有%d个苹果。",
	  "%s have an apple.": "%s有一个苹果。",
	  "%s have two apples.": "%s有两个苹果。",
	  "hello %s!": "你好%s！",
	  "hello world!": "你好世界！"
	}
	```
3. After `zh-Hans.json` has been translated, run `loaferi18n generate --pkg=catalog ./locales ./catalog/catalog.go`.

4. Import `catalog` package in main.go, example: `import _ "github.com/EasyGram-io/loafer-i18n/example/catalog"` 

### Translating new messages

If you have added new messages to your program:

1. Run `loaferi18n extract` to update `./locales/en.json` with the new messages.
2. Run `loaferi18n update ./locales/en.json` to generate updated `./locales/new-language.json` files.
3. Translate all the messages in the `./locales/new-language.json` files.
4. Run `loaferi18n generate --pkg=catalog ./locales ./catalog/catalog.go` to merge the translated messages into the go files.

## For examples:

- Look at an example [application](https://github.com/EasyGram-io/loafer-i18n/tree/master/example).

## Thanks:

- https://xuanwo.io/2019/12/11/golang-i18n/
- https://github.com/qingstor/qsctl/tree/master/pkg/i18n
- https://github.com/golang/text
- https://github.com/nicksnyder/go-i18n

## License

Loafer-i18n is available under the MIT license. See the [LICENSE](LICENSE) file for more info.