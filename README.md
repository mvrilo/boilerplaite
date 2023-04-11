# boilerplaite

[![GoDoc](https://godoc.org/github.com/mvrilo/boilerplaite?status.svg)](https://godoc.org/github.com/mvrilo/boilerplaite)
[![Go Report Card](https://goreportcard.com/badge/github.com/mvrilo/boilerplaite)](https://goreportcard.com/report/github.com/mvrilo/boilerplaite)

`boilerplaite` is a command-line tool and library to generate boilerplate code from a prompt using openai's chatgpt api.

## Installation

`go install github.com/mvrilo/boilerplaite/cmd/boilerplaite@latest`

## Usage

```
boilerplaite is a package and cli tool to generate boilerplate code from a prompt using openai's chatgpt api.

Usage:
  boilerplaite [flags]

Examples:
boilerplaite -o ./examples/go-cli-example -p "go program, package main, using cobra"

Flags:
  -h, --help            help for boilerplaite
  -m, --model string    OpenAI model (default "gpt-3.5-turbo")
  -o, --output string   Output directory
  -p, --prompt string   Prompt
  -t, --timeout int     Timeout in seconds (default 60)
  -v, --verbose         Verbose logs
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

- SQL tables and queries

```
./boilerplaite -o ./examples/sql-users-table -p "sql file containing create table statement for users, insert and update queries, the table should include id, name and email"
```

- A Python app using flask listening on port 3000.

```
./boilerplaite -o ./examples/python-flask-hello-world -p "python app using flask, hello world endpoint, listening on port 3000"
```

- A Ruby app using sinatra, fetchin data from sqlite db, listening on port 5000, with extra files.

```
./boilerplaite -o ./examples/ruby-sinatra -p "a ruby app using sinatra, fetching some data from a sqlite database, listening on port 5000, with dockerfile, gemspec and bundle files"
```

Check the output here: [./examples](examples).

## Notes

- Responses may contain innacurate or incorrect data.
- GPT-3.5-turbo is used as default model.
- Beware the high costs when changing the model to GPT-4.
- No overwrites when writing to files.
