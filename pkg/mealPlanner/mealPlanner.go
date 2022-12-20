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

type Meal struct {
	Id    string
	Title string
}

type DayMeals struct {
	Lunch  Meal
	Dinner Meal
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
		dayRecipes.Lunch.Title = recipeName
		dayRecipes.Lunch.Id = string(dayLunchRecipe.ID)

		dinnerRecipes := getRecipes(*client, "Sopar")
		dayDinnerRecipe := getRandomRecipe(dinnerRecipes)

		recipeName = getRecipeName(dayDinnerRecipe)
		dayRecipes.Dinner.Title = recipeName
		dayRecipes.Dinner.Id = string(dayDinnerRecipe.ID)

		weekRecipes = append(weekRecipes, dayRecipes)
	}

	fmt.Printf("Updating calendar page... \n")
	calendarDbId := notionapi.DatabaseID(os.Getenv("MEALS_CALENDAR"))

	query := new(notionapi.DatabaseQueryRequest)

	calendarDb, err := client.Database.Query(context.Background(), calendarDbId, query)
	if err != nil {
		log.Error().Msg("Error while getting db")

		// do something
	}

	log.Debug().Msgf("%+v\n", calendarDb.Results)

	for index, dayMeals := range weekRecipes {
		fmt.Printf("Day %d Lunch: %s \n", index, dayMeals.Lunch.Title)
		fmt.Printf("Day %d Dinner: %s \n", index, dayMeals.Dinner.Title)

		lunchRequest := new(notionapi.PageUpdateRequest)

		var printedProperties = fmt.Sprintf(`{"Recepta": {"type":"relation", "relation": [{"id": "%s"}],"has_more": false}}`, dayMeals.Lunch.Id)
		lunchRequest.Properties.UnmarshalJSON([]byte(printedProperties))

		var lunchIndex = index * 2
		var _, err = client.Page.Update(context.Background(), notionapi.PageID(calendarDb.Results[lunchIndex].ID), lunchRequest)
		if err != nil {
			log.Error().Msg("Error while getting db")

			// do something
		}

		dinnerRequest := new(notionapi.PageUpdateRequest)
		var dinnerProperties = fmt.Sprintf(`{"Recepta": {"type":"relation", "relation": [{"id": "%s"}],"has_more": false}}`, dayMeals.Dinner.Id)
		dinnerRequest.Properties.UnmarshalJSON([]byte(dinnerProperties))

		var dinnerIndex = index*2 + 1
		var _, dinnerErr = client.Page.Update(context.Background(), notionapi.PageID(calendarDb.Results[dinnerIndex].ID), dinnerRequest)
		if dinnerErr != nil {
			log.Error().Msg("Error while getting db")

			// do something
		}
	}

	// var test, error = client.Database.Get(context.Background(), calendarDbId)
	// log.Debug().Msgf("%+v\n", test)
	// log.Debug().Msgf("%+v\n", error)
	// var testPage, errorPage = client.Block.Get(context.Background(), test.Parent.BlockID)
	// log.Debug().Msgf("%+v\n", testPage)
	// log.Debug().Msgf("%+v\n", errorPage)
	// request := new(notionapi.PageUpdateRequest)

	// client.Page.Update(context.Background(), calendarDbId, request)
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
