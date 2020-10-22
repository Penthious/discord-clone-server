package main

import (
	"discord-clone-server/cmd"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "discord-clone-server",
	}

	rootCmd.SetOut(os.Stdout)
	rootCmd.AddCommand(
		cmd.ServerCmd,
		cmd.MigrateCmd,
		cmd.SeedCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
