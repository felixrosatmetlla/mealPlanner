package mealPlanner

import (
	"github.com/felixrosatmetlla/mealPlanner/pkg/mealPlanner"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Plans meals",
	Run: func(cmd *cobra.Command, args []string) {
		mealPlanner.Plan()
	},
}

func init() {
	rootCmd.AddCommand(planCmd)
}
