package main

import (
    "fmt"
    "github.com/spf13/cobra"
)

func main() {
    rootCmd := &cobra.Command{
        Use:   "myprogram",
        Short: "My program is a CLI tool",
        Long:  `My program is a CLI tool that helps with boilerplate code generation`,
        Run: func(cmd *cobra.Command, args []string) {
            // your code here
        },
    }

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
