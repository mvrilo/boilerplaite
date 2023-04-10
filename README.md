# boilerplaite

`boilerplaite` is a command line tool and library to generate boilerplate code from a prompt using openai's chatgpt api.

## Installation

```
go install github.com/mvrilo/boilerplaite/cmd/boilerplaite@latest
```

## Usage

```
boilerplaite -h
boilerplaite is a package and cli tool to generate boilerplate code from a prompt using openai's chatgpt api.

Usage:
  boilerplaite [flags]

Examples:
boilerplaite -o ./examples/go-cli-example -p "go program, package main, using cobra, check imports and errors"

Flags:
  -h, --help            help for boilerplaite
  -m, --model string    OpenAI model (default "gpt-3.5-turbo")
  -o, --output string   Output directory
  -p, --prompt string   Prompt
```

See examples below.

## Configuration

Most options can be used as flags as shown below in "usage".

Except for OpenAI api key which must be set as environment variable: `BOILERPLAITE_OPENAI_KEY`

## Examples

- A Go program using cobra.

```
boilerplaite -o ./examples/go-cli-example -p "go program, package main, using cobra"
```

- A Rust server using axum.

```
boilerplaite -o ./examples/rust-axum-hello -p "hello world program using rust with axum server and cli with clap"
```

Check the output here: [./examples](./examples).
