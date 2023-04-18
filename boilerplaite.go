package boilerplaite

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/chat"
	yaml "gopkg.in/yaml.v3"
)

var ErrOutputExists = errors.New("Output directory already exists")

const promptTemplate = `You are an AI system called "boilerplaite", your objective is to get me faster in writing code in any programming language, this means generating boilerplate codes. Your answers must be in the format of yaml/yml and it MUST be valid. Don't add any notes. If a file is empty you may skip it. You will provide the code and the filepath where that code should live, so for a task like this "sample package.json file for a nodejs project" you should return something like this:
- filepath: ./package.json
  code: |
  {
	  "name": "sample",
	  "version": "1.0.0",
	  "main": "index.js",
	  "license": "MIT"
  }
Another example would be: "2 files as txt with the characters a and b in each file"
- filepath: ./a.txt
	code: |
	a
- filepath: ./b.txt
	code: |
	b
This is your task: %s`

type entry struct {
	Filepath string `yaml:"filepath"`
	Code     string `yaml:"code"`
}

type Boilerplaite struct {
	chatgpt *chat.Client
}

func New(openaiKey, openaiModel string) *Boilerplaite {
	sess := openai.NewSession(openaiKey)
	sess.HTTPClient.Timeout = 5 * time.Minute
	chatgpt := chat.NewClient(sess, openaiModel)
	return &Boilerplaite{chatgpt}
}

func (b *Boilerplaite) Complete(ctx context.Context, data string) (string, int, error) {
	msgs := []*chat.Message{{Role: "user", Content: fmt.Sprintf(promptTemplate, data)}}
	params := &chat.CreateCompletionParams{Messages: msgs}
	resp, err := b.chatgpt.CreateCompletion(ctx, params)
	if err != nil {
		return "", 0, err
	}

	replymsg := resp.Choices[0].Message
	content := replymsg.Content

	return content, resp.Usage.TotalTokens, nil
}

func (b *Boilerplaite) prepare(ctx context.Context, outdir string) error {
	if _, err := os.Stat(outdir); err == os.ErrExist {
		return ErrOutputExists
	}
	return os.MkdirAll(outdir, 0750)
}

func (b *Boilerplaite) WritePrompt(ctx context.Context, outdir string, data string) error {
	if err := b.prepare(ctx, outdir); err != nil {
		return err
	}
	filepath := filepath.Join(outdir, ".boilerplaite.prompt")
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	if _, err = f.Write([]byte(data)); err != nil {
		return err
	}
	return f.Close()
}

func (b *Boilerplaite) WriteFiles(ctx context.Context, outdir string, data string) error {
	if err := b.prepare(ctx, outdir); err != nil {
		return err
	}

	var out []entry
	if err := yaml.Unmarshal([]byte(data), &out); err != nil {
		return err
	}

	for _, entry := range out {
		file := filepath.Join(outdir, entry.Filepath)

		dir := filepath.Dir(file)
		if err := os.MkdirAll(dir, 0750); err != nil {
			return err
		}

		f, err := os.Create(file)
		if err != nil {
			if err := os.RemoveAll(outdir); err != nil {
				return err
			}
			return err
		}

		if _, err = f.WriteString(entry.Code); err != nil {
			if err := os.Remove(file); err != nil {
				return err
			}
			return err
		}
	}

	return nil
}
