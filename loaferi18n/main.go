package main

//go:generate go run main.go extract . ./locales/en.json
//go:generate go run main.go update ./locales/en.json ./locales/zh-Hans.json
//go:generate go run main.go update ./locales/en.json ./locales/zh-Hant.json
//go:generate go run main.go generate --pkg=catalog ./locales ./catalog/catalog.go

import (
	"fmt"
	"log"
	"os"

	"github.com/EasyGram-io/loafer-i18n/i18n"
	_ "github.com/EasyGram-io/loafer-i18n/loaferi18n/catalog"
	"github.com/Xuanwo/go-locale"
	"github.com/urfave/cli/v2"
)

func main() {
	// Detect OS language
	tag, _ := locale.Detect()

	// Set Language
	i18n.SetLang(tag)

	appName := "loaferi18n"

	app := &cli.App{
		HelpName: appName,
		Name:     appName,
		Usage:    i18n.Sprintf(`a tool for managing message translations.`),
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     i18n.Sprintf(`merge translations and generate catalog`),
				UsageText: i18n.Sprintf(`%s update srcfile destfile`, appName),
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "flush",
						Aliases: []string{"f"},
						Value:   false,
						Usage:   fmt.Sprintf(`flush messages`),
					},
				},
				Action: func(c *cli.Context) error {
					srcFile := c.Args().Get(0)
					if len(srcFile) == 0 {
						return fmt.Errorf(i18n.Sprintf(`srcfile cannot be empty`))
					}

					destFile := c.Args().Get(1)
					if len(destFile) == 0 {
						return fmt.Errorf(i18n.Sprintf(`destfile cannot be empty`))
					}

					flush := c.Bool("flush")
					err := i18n.Update(srcFile, destFile, flush)

					return err
				},
			},
			{
				Name:      "extract",
				Aliases:   []string{"e"},
				Usage:     i18n.Sprintf(`extracts strings to be translated from code`),
				UsageText: i18n.Sprintf(`%s extract [path] [outfile]`, appName),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "pkg",
						Value: "i18n",
						Usage: i18n.Sprintf(`package name`),
					},
				},
				Action: func(c *cli.Context) error {
					path := c.Args().Get(0)
					if len(path) == 0 {
						path = "."
					}
					outFile := c.Args().Get(1)
					if len(outFile) == 0 {
						outFile = "./locales/en.json"
					}
					pkgName := c.String("pkg")
					err := i18n.Extract(pkgName, []string{
						path,
					}, outFile)
					return err
				},
			},
			{
				Name:      "generate",
				Aliases:   []string{"g"},
				Usage:     i18n.Sprintf(`generates code to insert translated messages`),
				UsageText: i18n.Sprintf(`%s generate [path] [outfile]`, appName),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "pkg",
						Value: "catalog",
						Usage: i18n.Sprintf(`generated go file package name`),
					},
				},
				Action: func(c *cli.Context) error {
					path := c.Args().Get(0)
					if len(path) == 0 {
						path = "./locales"
					}
					outFile := c.Args().Get(1)
					if len(outFile) == 0 {
						outFile = "./catalog/catalog.go"
					}
					pkgName := c.String("pkg")
					err := i18n.Generate(
						pkgName,
						[]string{
							path,
						}, outFile)
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
