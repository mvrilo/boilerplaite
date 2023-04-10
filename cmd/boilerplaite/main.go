package main

import (
	"context"

	"github.com/charmbracelet/log"
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
}

func main() {
	godotenv.Load()

	var conf Config
	err := envconfig.Process("boilerplaite", &conf)
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := &cobra.Command{
		Use:     "boilerplaite",
		Example: `boilerplaite -o ./examples/go-cli-example -p "go program, package main, using cobra"`,
		Short:   "boilerplaite is a package and cli tool to generate boilerplate code from a prompt using openai's chatgpt api.",
		Run: func(c *cobra.Command, args []string) {
			if conf.Output == "" {
				log.Fatal("Output dir flag is required")
			}
			if conf.Prompt == "" {
				log.Fatal("Prompt flag is required")
			}
			bp := boilerplaite.New(conf.OpenAIKey, conf.OpenAIModel)
			ctx := context.Background()
			data, err := bp.Complete(ctx, conf.Prompt)
			if err != nil {
				log.Fatal(err)
			}
			err = bp.Save(ctx, conf.Output, data)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.Flags().StringVarP(&conf.OpenAIModel, "model", "m", "gpt-3.5-turbo", "OpenAI model")
	rootCmd.Flags().StringVarP(&conf.Output, "output", "o", "", "Output directory")
	rootCmd.Flags().StringVarP(&conf.Prompt, "prompt", "p", "", "Prompt")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
