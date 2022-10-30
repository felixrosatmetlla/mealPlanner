package mealPlanner

import (
	"context"
	"fmt"
	"os"

	"github.com/jomei/notionapi"
	"github.com/rs/zerolog/log"
)

func Plan() {
	notionToken := notionapi.Token(os.Getenv("NOTION_TOKEN"))
	log.Debug().Msgf("'%s' \n", notionToken)

	client := notionapi.NewClient(notionToken)

	log.Debug().Msg("Client Created")
	fmt.Println()

	getRecipes(*client)
}

func ListRecipes() {
	notionToken := notionapi.Token(os.Getenv("NOTION_TOKEN"))
	client := notionapi.NewClient(notionToken)

	log.Debug().Msg("Client Created")

	recipes := getRecipes(*client)

	for index, recipe := range recipes {
		log.Debug().Msgf("%+v\n", recipe)
		if recipeNameProperty, ok := recipe.Properties["Name"].(*notionapi.TitleProperty); ok {
			recipeName := recipeNameProperty.Title[0].PlainText
			fmt.Printf("%d. %s \n", index+1, recipeName)
		} else {
			log.Error().Msg("There was an error getting the recipe name")
		}

	}
}

func getRecipes(client notionapi.Client) []notionapi.Page {
	recipesDbId := notionapi.DatabaseID(os.Getenv("RECIPES_DB_ID"))

	query := new(notionapi.DatabaseQueryRequest)
	recipesDb, err := client.Database.Query(context.Background(), recipesDbId, query)
	if err != nil {
		log.Error().Msg("Error while getting db")

		// do something
	}

	log.Debug().Msgf("%+v\n", recipesDb.Results)

	return recipesDb.Results
}
