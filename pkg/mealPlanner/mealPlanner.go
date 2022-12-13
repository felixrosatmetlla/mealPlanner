package mealPlanner

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jomei/notionapi"
	"github.com/rs/zerolog/log"
)

type DayMeals struct {
	Lunch  string
	Dinner string
}

func Plan(tags []string) {
	notionToken := notionapi.Token(os.Getenv("NOTION_TOKEN"))
	log.Debug().Msgf("'%s' \n", notionToken)

	client := notionapi.NewClient(notionToken)

	log.Debug().Msg("Client Created")
	fmt.Println()

	var recipes []notionapi.Page
	if len(tags) != 0 {
		recipes = getRecipes(*client, tags[0])
	} else {
		recipes = getRecipes(*client, "")
	}

	chosenRecipe := getRandomRecipe(recipes)

	log.Debug().Msgf("%+v\n", chosenRecipe)
	recipeName := getRecipeName(chosenRecipe)
	fmt.Printf("Random receipe chosen: %s \n", recipeName)

	var weekRecipes []DayMeals

	for i := 0; i < 7; i++ {
		lunchRecipes := getRecipes(*client, "Dinar")
		dayLunchRecipe := getRandomRecipe(lunchRecipes)

		var dayRecipes DayMeals

		recipeName := getRecipeName(dayLunchRecipe)
		dayRecipes.Lunch = recipeName

		dinnerRecipes := getRecipes(*client, "Sopar")
		dayDinnerRecipe := getRandomRecipe(dinnerRecipes)

		recipeName = getRecipeName(dayDinnerRecipe)
		dayRecipes.Dinner = recipeName

		weekRecipes = append(weekRecipes, dayRecipes)
	}

	for index, dayMeals := range weekRecipes {
		fmt.Printf("Day %d Lunch: %s \n", index, dayMeals.Lunch)
		fmt.Printf("Day %d Dinner: %s \n", index, dayMeals.Dinner)
	}
}

func ListRecipes(tags []string) {
	notionToken := notionapi.Token(os.Getenv("NOTION_TOKEN"))
	client := notionapi.NewClient(notionToken)

	log.Debug().Msg("Client Created")

	var recipes []notionapi.Page
	if len(tags) != 0 {
		recipes = getRecipes(*client, tags[0])
	} else {
		recipes = getRecipes(*client, "")
	}

	for index, recipe := range recipes {
		log.Debug().Msgf("%+v\n", recipe)
		recipeName := getRecipeName(recipe)
		fmt.Printf("%d. %s \n", index+1, recipeName)

	}
}

func getRecipes(client notionapi.Client, tag string) []notionapi.Page {
	recipesDbId := notionapi.DatabaseID(os.Getenv("RECIPES_DB_ID"))

	query := new(notionapi.DatabaseQueryRequest)
	if tag != "" {
		filter := new(notionapi.PropertyFilter)
		filter.Property = "Tags"
		filter.MultiSelect = new(notionapi.MultiSelectFilterCondition)
		filter.MultiSelect.Contains = tag
		query.Filter = filter
	}
	recipesDb, err := client.Database.Query(context.Background(), recipesDbId, query)
	if err != nil {
		log.Error().Msg("Error while getting db")

		// do something
	}

	log.Debug().Msgf("%+v\n", recipesDb.Results)

	return recipesDb.Results
}

func getRecipeName(recipe notionapi.Page) string {
	var recipeName string

	if recipeNameProperty, ok := recipe.Properties["Name"].(*notionapi.TitleProperty); ok {
		recipeName = recipeNameProperty.Title[0].PlainText

	} else {
		log.Error().Msg("There was an error getting the recipe name")
	}

	return recipeName
}

func getRandomRecipe(recipes []notionapi.Page) notionapi.Page {
	limitNumber := len(recipes)
	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(limitNumber)

	return recipes[index]
}
