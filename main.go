//go:generate tailwindcss build -o static/css/tailwind.css --minify && templ generate

package main

import (
	"fmt"
	"net/http"

	"github.com/hasura/go-graphql-client"
	"tarta.com/modules/catalog"
	"tarta.com/modules/common"
	"tarta.com/modules/pages"
)

var GraphQLClient = graphql.NewClient("https://agenciam3.myvtex.com/_v/private/graphql/v1", nil)

func main() {

	result, err := catalog.ProductSearch(GraphQLClient, catalog.ProductSearchInput{
		From:                 0,
		To:                   9,
		FullText:             "",
		OrderBy:              "OrderByScoreDESC",
		PriceRange:           "",
		SalesChannel:         "",
		HideUnavailableItems: false,
		SimulationBehavior:   catalog.SimulationBehaviorDefault,
		SelectedFacets:       []catalog.SelectedFacetInput{},
	})

	print(err.Error())

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", common.LayoutHandler(pages.HomePage(result)))
	http.Handle("/about", common.LayoutHandler(pages.AboutPage()))
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
