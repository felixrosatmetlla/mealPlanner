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

	getRecipes(*client)
}

func ListRecipes() {
	notionToken := notionapi.Token(os.Getenv("NOTION_TOKEN"))
	fmt.Fprintf(os.Stdout, "'%s' \n", notionToken)

	client := notionapi.NewClient(notionToken)

	fmt.Println("Client Created")

	recipes := getRecipes(*client)

	for index, recipe := range recipes {
		fmt.Printf("%+v\n", recipe)
		if recipeNameProperty, ok := recipe.Properties["Name"].(*notionapi.TitleProperty); ok {
			recipeName := recipeNameProperty.Title[0].PlainText
			fmt.Println(index+1, "-", recipeName)
		} else {
			fmt.Fprintf(os.Stderr, "There was an error getting the recipe name")
		}

	}
}

func getRecipes(client notionapi.Client) []notionapi.Page {
	recipesDbId := notionapi.DatabaseID(os.Getenv("RECIPES_DB_ID"))

	query := new(notionapi.DatabaseQueryRequest)
	recipesDb, err := client.Database.Query(context.Background(), recipesDbId, query)
	if err != nil {
		fmt.Println("Error while getting db")

		// do something
	}

	fmt.Printf("%+v\n", recipesDb.Results)

	return recipesDb.Results
}
