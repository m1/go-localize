// Code generated by go-localize; DO NOT EDIT.
// This file was generated by robots at
// 2020-11-11 22:23:20.406208 -0800 PST m=+0.001173998

package localizations

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var localizations = map[string]string{
	"en.messages.hello": "hello",
	"en.messages.hello_firstname_lastname": "Hello {{.firstname}} {{.lastname}}",
	"en.messages.hello_my_name_is": "Hello my name is {{.name}}",
	"en.messages.how_are_you": "How are you?",
	"en.messages.whats_your_name": "What's your name?",
	"es.customer.messages.hello": "hello customer!",
	"es.messages.hello": "Hola",
	"es.messages.hello_my_name_is": "Hola, mi nombre es {{.name}}",
	"es.messages.how_are_you": "¿Cómo estás?",
	"es.messages.whats_your_name": "¿Cuál es tu nombre?",
}

type Replacements map[string]interface{}

type Localizer struct {
	Locale	 string
	FallbackLocale string
	Localizations  map[string]string
}

func New(locale string, fallbackLocale string) *Localizer {
	t := &Localizer{Locale: locale, FallbackLocale: fallbackLocale}
	t.Localizations = localizations
	return t
}

func (t Localizer) SetLocales(locale, fallback string) Localizer {
	t.Locale = locale
	t.FallbackLocale = fallback
	return t
}

func (t Localizer) SetLocale(locale string) Localizer {
	t.Locale = locale
	return t
}

func (t Localizer) SetFallbackLocale(fallback string) Localizer {
	t.FallbackLocale = fallback
	return t
}

func (t Localizer) GetWithLocale(locale, key string, replacements ...*Replacements) string {
	str, ok := t.Localizations[t.getLocalizationKey(locale, key)]
	if !ok {
		str, ok = t.Localizations[t.getLocalizationKey(t.FallbackLocale, key)]
		if !ok {
			return key
		}
	}

        // If the str doesn't have any substitutions, no need to
        // template.Execute.
	if strings.Index(str, "}}") == -1 {
                return str
        }

	return t.replace(str, replacements...)
}

func (t Localizer) Get(key string, replacements ...*Replacements) string {
	str := t.GetWithLocale(t.Locale, key, replacements...)
	return str
}

func (t Localizer) getLocalizationKey(locale string, key string) string {
	return fmt.Sprintf("%v.%v", locale, key)
}

func (t Localizer) replace(str string, replacements ...*Replacements) string {
	b := &bytes.Buffer{}
	tmpl, err := template.New("").Parse(str)
	if err != nil {
		return str
	}

	replacementsMerge := Replacements{}
	for _, replacement := range replacements {
		for k, v := range *replacement {
			replacementsMerge[k] = v
		}
	}

	err = template.Must(tmpl, err).Execute(b, replacementsMerge)
	if err != nil {
		return str
	}
	buff := b.String()
	return buff
}
