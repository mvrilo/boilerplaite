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

const promptTemplate = `You are an AI system called "boilerplaite", your objective is to get me faster in writing code in any programming language, this means generating boilerplate codes. Your answers must be in the format of valid Yaml files only. You will provide the code and the filepath where that code should live, so for a task like this "sample package.json file for a nodejs project" you should return something like this:
- filepath: ./package.json
  code: |
  {
	  "name": "sample",
	  "version": "1.0.0",
	  "description": "",
	  "main": "index.js",
	  "scripts": {
	  	"test": "echo \"Error: no test specified\" && exit 1"
	  },
	  "author": "",
	  "license": "MIT"
  }
Another example would be: "2 files as txt with the characters a and b in each file"
- filepath: ./a.txt
	code: |
	a
- filepath: ./b.txt
	code: |
	a
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

func (b *Boilerplaite) Complete(ctx context.Context, data string) (string, error) {
	msgs := []*chat.Message{{Role: "user", Content: fmt.Sprintf(promptTemplate, data)}}
	params := &chat.CreateCompletionParams{Messages: msgs}
	resp, err := b.chatgpt.CreateCompletion(ctx, params)
	if err != nil {
		return "", err
	}

	replymsg := resp.Choices[0].Message
	content := replymsg.Content

	return content, nil
}

func (b *Boilerplaite) Save(ctx context.Context, outdir string, data string) error {
	var out []entry
	if err := yaml.Unmarshal([]byte(data), &out); err != nil {
		return err
	}

	if _, err := os.Stat(outdir); err == os.ErrExist {
		return errors.New("Output directory already exists")
	}

	for _, entry := range out {
		file := filepath.Join(outdir, entry.Filepath)

		dir := filepath.Dir(file)
		if err := os.MkdirAll(dir, 0750); err != nil {
			return err
		}

		println(file)
		f, err := os.Create(file)
		if err != nil {
			os.RemoveAll(outdir)
			return err
		}

		_, err = f.WriteString(entry.Code)
		if err != nil {
			return err
		}
	}

	return nil
}
