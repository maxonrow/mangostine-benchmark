package main

import (
	"os"

	"github.com/mangostine-benchmark-go/mangostine"
	"github.com/spf13/cobra"
)

var (
	defaultCLIHome = os.ExpandEnv("$HOME/.mangostine")
)

func main() {
	cobra.EnableCommandSorting = true

	rootCmd := &cobra.Command{
		Use:   "mangostine",
		Short: "mangostine client",
	}

	rootCmd.AddCommand(
		mangostine.CreateKeyPairCommand(),
		mangostine.SendTxCommand(),
	)
	// executor := (rootCmd, "mangostine", defaultCLIHome)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

