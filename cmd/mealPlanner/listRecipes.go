package mealPlanner

import (
	"github.com/felixrosatmetlla/mealPlanner/pkg/mealPlanner"
	"github.com/spf13/cobra"
)

var listRecipesCmd = &cobra.Command{
	Use:     "listRecipes",
	Aliases: []string{"listr"},
	Short:   "List existing recipes",
	Run: func(cmd *cobra.Command, args []string) {
		mealPlanner.ListRecipes(tags)
	},
}

func init() {
	rootCmd.AddCommand(listRecipesCmd)
}
