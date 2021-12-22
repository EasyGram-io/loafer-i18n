package i18n

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Printer is printer
type Printer struct {
	lang string
	pt   *message.Printer
}

// PluralRule is Plural rule
type PluralRule struct {
	Pos   int
	Expr  string
	Value int
	Text  string
}

// Message is translation message
type Message map[string]string

// NewPrinter is new printer
func NewPrinter(lang interface{}) *Printer {
	var langTag language.Tag
	switch _lang := lang.(type) {
	case language.Tag:
		langTag = _lang
	case string:
		langTag = language.Make(_lang)
	}
	return &Printer{
		lang: langTag.String(),
		pt:   message.NewPrinter(langTag),
	}
}

// Printf is like fmt.Printf, but using language-specific formatting.
func (p *Printer) Printf(format string, args ...interface{}) {
	format, args = preArgs(format, args...)
	p.pt.Printf(format, args...)
}

// Sprintf is like fmt.Sprintf, but using language-specific formatting.
func (p *Printer) Sprintf(format string, args ...interface{}) string {
	format, args = preArgs(format, args...)
	return p.pt.Sprintf(format, args...)
}

// Fprintf is like fmt.Fprintf, but using language-specific formatting.
func (p *Printer) Fprintf(w io.Writer, key string, a ...interface{}) (n int, err error) {
	format, args := preArgs(key, a...)
	_key := message.Reference(format)
	return p.pt.Fprintf(w, _key, args...)
}

// String is lang
func (p *Printer) String() string {
	return strings.ToLower(p.lang)
}

// Preprocessing parameters in plural form
func preArgs(format string, args ...interface{}) (string, []interface{}) {
	length := len(args)
	if length > 0 {
		lastArg := args[length-1]
		switch lastArg.(type) {
		case []PluralRule:
			rules := lastArg.([]PluralRule)
			// parse rule
			for _, rule := range rules {
				curPosVal := args[rule.Pos-1].(int)
				// Support comparison expression
				if (rule.Expr == "=" && curPosVal == rule.Value) || (rule.Expr == ">" && curPosVal > rule.Value) {
					format = rule.Text
					break
				}
			}
			args = args[0:strings.Count(format, "%")]
		}
	}
	return format, args
}

// Plural is Plural function
func Plural(cases ...interface{}) []PluralRule {
	rules := []PluralRule{}
	// %[1]d=1, %[1]d>1
	re := regexp.MustCompile(`\[(\d+)\][^=>]\s*(\=|\>)\s*(\d+)$`)
	for i := 0; i < len(cases); {
		expr := cases[i].(string)
		if i++; i >= len(cases) {
			return rules
		}
		text := cases[i].(string)
		// cannot match continue
		if !re.MatchString(expr) {
			continue
		}
		matches := re.FindStringSubmatch(expr)
		pos, _ := strconv.Atoi(matches[1])
		value, _ := strconv.Atoi(matches[3])
		rules = append(rules, PluralRule{
			Pos:   pos,
			Expr:  matches[2],
			Value: value,
			Text:  text,
		})
		i++
	}
	return rules
}
