package mealPlanner

import (
	"context"
	"fmt"
	"os"

	"github.com/jomei/notionapi"
)

func Plan() {
	notionToken := notionapi.Token(os.Getenv("NOTION_TOKEN"))
	fmt.Fprintf(os.Stdout, "'%s' \n", notionToken)

	client := notionapi.NewClient(notionToken)

	fmt.Println("Client Created")
	recipesDbId := notionapi.DatabaseID(os.Getenv("RECIPES_DB_ID"))

	query := new(notionapi.DatabaseQueryRequest)
	recipesDb, err := client.Database.Query(context.Background(), recipesDbId, query)
	if err != nil {
		fmt.Println("Error while getting db")

		// do something
	}

	fmt.Printf("%+v\n", recipesDb.Results)
}
