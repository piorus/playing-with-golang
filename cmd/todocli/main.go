package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/piotr-rusin/playing-with-golang/internal/app/todocli"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "todocli",
		Short: "TODO CLI manages TODOs.",
	}

	todocli.Configure(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
