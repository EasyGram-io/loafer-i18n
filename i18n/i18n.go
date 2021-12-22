package i18n

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var p *Printer

// init
func init() {
	// default language english
	p = NewPrinter(language.English)
}

// SetLang set language
func SetLang(lang interface{}) {
	p = NewPrinter(lang)
}

// GetPrinter
func GetPrinter() *Printer {
	return p
}

// Printf is like fmt.Printf, but using language-specific formatting.
func Printf(format string, args ...interface{}) {
	p.Printf(format, args...)
}

// Sprintf is like fmt.Sprintf, but using language-specific formatting.
func Sprintf(format string, args ...interface{}) string {
	return p.Sprintf(format, args...)
}

// Fprintf is like fmt.Fprintf, but using language-specific formatting.
func Fprintf(w io.Writer, key string, args ...interface{}) (n int, err error) {
	return p.Fprintf(w, key, args...)
}

func unmarshal(path string) (*Message, error) {
	result := &Message{}

	_, err := os.Stat(path)
	if err != nil {
		if !os.IsExist(err) {
			return result, nil
		}
	}

	fileExt := strings.ToLower(filepath.Ext(path))
	if fileExt != ".toml" && fileExt != ".json" && fileExt != ".yaml" {
		return result, fmt.Errorf(Sprintf("File type not supported"))
	}

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return result, err
	}

	if strings.HasSuffix(fileExt, ".json") {
		err := json.Unmarshal(buf, result)
		if err != nil {
			return result, err
		}
	}

	if strings.HasSuffix(fileExt, ".yaml") {
		err := yaml.Unmarshal(buf, result)
		if err != nil {
			return result, err
		}
	}

	if strings.HasSuffix(fileExt, ".toml") {
		_, err := toml.Decode(string(buf), result)
		if err != nil {
			return result, err
		}
	}
	return result, nil

}

func marshal(v interface{}, format string) ([]byte, error) {
	switch format {
	case "json":
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "  ")
		err := encoder.Encode(v)
		return buffer.Bytes(), err
	case "toml":
		var buf bytes.Buffer
		enc := toml.NewEncoder(&buf)
		enc.Indent = ""
		err := enc.Encode(v)
		return buf.Bytes(), err
	case "yaml":
		return yaml.Marshal(v)
	}
	return nil, fmt.Errorf("unsupported format: %s", format)
}
