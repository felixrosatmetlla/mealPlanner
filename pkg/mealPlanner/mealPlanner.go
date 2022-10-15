package mealPlanner

import (
	"fmt"
	"os"
)

func Plan() {
	notionToken := os.Getenv("NOTION_TOKEN")
	fmt.Fprintf(os.Stdout, "'%s'", notionToken)
}
