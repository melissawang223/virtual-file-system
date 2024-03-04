package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Virtual File System",
	Short: "Virtual File System",
	Long:  `Virtual File System written in Go.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
