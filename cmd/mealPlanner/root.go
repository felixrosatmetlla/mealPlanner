package mealPlanner

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var tags []string
var rootCmd = &cobra.Command{
	Use:     "mealPlanner",
	Version: version,
	Short:   "mealPlanner - CLI application plan meals",
	Long:    "mealPlanner - CLI application to interact with notion and plan meals",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	rootCmd.PersistentFlags().StringArrayVarP(&tags, "tags", "t", *new([]string), "tags to filter")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error executing the CLI '%s'", err)
		os.Exit(1)
	}
}
