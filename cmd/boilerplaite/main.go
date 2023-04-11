package main

import (
	"context"
	"fmt"
	"time"

	logger "github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"

	"github.com/mvrilo/boilerplaite"
)

type Config struct {
	OpenAIKey   string `envconfig:"openai_key"`
	OpenAIModel string `envconfig:"openai_model"`
	Output      string `envconfig:"output"`
	Prompt      string `envconfig:"prompt"`
	Timeout     int    `envconfig:"timeout"`
	Verbose     bool   `envconfig:"verbose"`
}

func main() {
	_ = godotenv.Load()

	var conf Config
	err := envconfig.Process("boilerplaite", &conf)
	if err != nil {
		logger.Fatal(err)
	}

	log := logger.Default()
	log.SetLevel(logger.InfoLevel)

	rootCmd := &cobra.Command{
		Use:     "boilerplaite",
		Example: `boilerplaite -o ./examples/go-cli-example -p "go program, package main, using cobra"`,
		Short:   "boilerplaite is a package and cli tool to generate boilerplate code from a prompt using openai's chatgpt api.",
		Run: func(c *cobra.Command, args []string) {
			if conf.Prompt == "" {
				_ = c.Usage()
				return
			}

			if conf.Verbose {
				log.SetLevel(logger.DebugLevel)
			}

			bp := boilerplaite.New(conf.OpenAIKey, conf.OpenAIModel)

			log.Debug("boilerplaite", "timeout", fmt.Sprintf("%d seconds", conf.Timeout))
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(conf.Timeout)*time.Second)
			defer cancel()

			log.Debug("boilerplaite", "model", conf.OpenAIModel)
			log.Debug("boilerplaite", "request", "chat completion")

			data, tokens, err := bp.Complete(ctx, conf.Prompt)
			if err != nil {
				log.Fatal(err)
			}
			log.Debug("boilerplaite", "tokens", tokens)

			if conf.Output != "" {
				log.Debug("boilerplaite", "output", conf.Output)
				if err = bp.Save(ctx, conf.Output, data); err != nil {
					log.Fatal(err)
				}
				return
			}

			log.Debug("boilerplaite", "output", "STDOUT")
			fmt.Println(string(data))
		},
	}

	rootCmd.Flags().StringVarP(&conf.OpenAIModel, "model", "m", "gpt-3.5-turbo", "OpenAI model")
	rootCmd.Flags().StringVarP(&conf.Output, "output", "o", "", "Output directory")
	rootCmd.Flags().StringVarP(&conf.Prompt, "prompt", "p", "", "Prompt")
	rootCmd.Flags().IntVarP(&conf.Timeout, "timeout", "t", 60, "Timeout in seconds")
	rootCmd.Flags().BoolVarP(&conf.Verbose, "verbose", "v", false, "Verbose logs")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
