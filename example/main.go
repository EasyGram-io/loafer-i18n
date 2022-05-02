package main

//go:generate loaferi18n extract . ./locales/en.json
//go:generate loaferi18n update ./locales/en.json ./locales/zh-Hans.json
//go:generate loaferi18n generate --pkg=catalog ./locales ./catalog/catalog.go
//go:generate go build -o example

import (
	"fmt"
	"os"

	"golang.org/x/text/language"

	_ "github.com/EasyGram-io/loafer-i18n/example/catalog"
	"github.com/EasyGram-io/loafer-i18n/i18n"
)

func main() {

	p := i18n.NewPrinter(language.SimplifiedChinese)
	p.Printf(`hello world.`)
	fmt.Println()

	i18n.SetLang(language.SimplifiedChinese)

	i18n.Printf(`hello world!`)
	fmt.Println()

	i18n.Printf(`hello world!`, i18n.Domain{`example`})
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
