# go-localize

[![GoDoc](https://godoc.org/github.com/m1/go-localize?status.svg)](https://godoc.org/github.com/m1/go-localize)
[![Build Status](https://travis-ci.org/m1/go-localize.svg?branch=master)](https://travis-ci.org/m1/go-localize)
[![Go Report Card](https://goreportcard.com/badge/github.com/m1/go-localize)](https://goreportcard.com/report/github.com/m1/go-localize)
[![Release](https://img.shields.io/github/release/m1/go-localize.svg)](https://github.com/m1/go-localize/releases/latest)
[![codecov](https://codecov.io/gh/m1/go-localize/branch/master/graph/badge.svg)](https://codecov.io/gh/m1/go-localize)

__Simple and easy to use i18n (Internationalization and localization) engine written in Go, used for translating locale strings. 
Use with [go generate](#go-generate) or on the [CLI](#cli). Currently supports JSON and YAML translation files__

## Usage

### Go generate

The suggested way to use go-localize is to use `go generate`. For example, take the following directory structure:

```
goapp
└── localizations_src
    ├── en
    │   └── messages.yaml
    └── es
        ├── customer
        │   └── messages.json
        └── messages.json
```

Example of JSON translation file:

```json
{
  "hello": "Hola",
  "how_are_you": "¿Cómo estás?",
  "whats_your_name": "¿Cuál es tu nombre?",
  "hello_my_name_is": "Hola, mi nombre es {{.name}}"
}
```

Example of YAML translation file:
```yaml
hello: hello
how_are_you: How are you?
whats_your_name: "What's your name?"
hello_my_name_is: Hello my name is {{.name}}
hello_firstname_lastname: Hello {{.firstname}} {{.lastname}}
```

To then generate the localization package, add the following to your `main.go` or another one of your `.go` files:

```go
//go:generate go-localize -input localizations_src -output localizations
// 
```

Now you'll be able to use the localization like so:
```go
l := localizations.New("en", "es")

println(l.Get("messages.how_are_you")) // How are you?

println(l.GetWithLocale("es", "messages.hello_my_name_is", &localizations.Replacements{"name":"steve"})) // "Hola, mi nombre es steve"
```

With `en` being the locale and `es` being the fallback. The localization keys are worked out using folder structure, eg:

`en/customer/messages.json` with the contents being:
```json
{
  "hello": "hello customer!"
}
```
You'll be able to access this using the key: `customer.messages.hello`.

#### Suggestions

It is suggested to instead of using hardcoded locale keys i.e. `en` to use the language keys included in key, i.e: `language.BritishEnglish.String()` 
which is "en-GB"

### CLI

Instead of using go generate you can just generate the localizations manually using `go-localize`:
```
Usage of go-localize:
  -input string
        input localizations folder
  -output string
        where to output the generated package
```
