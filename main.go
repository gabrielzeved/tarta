//go:generate tailwindcss build -o static/css/tailwind.css --minify && templ generate

package main

import (
	"fmt"
	"net/http"

	"tarta.com/modules/catalog"
	"tarta.com/modules/common"
	"tarta.com/modules/pages"
)

func main() {

	input := catalog.ProductSearchInput{
		From:                 0,
		To:                   9,
		FullText:             "",
		OrderBy:              "OrderByScoreDESC",
		PriceRange:           "",
		SalesChannel:         "",
		HideUnavailableItems: false,
		SimulationBehavior:   catalog.SimulationBehaviorDefault,
		SelectedFacets:       []catalog.SelectedFacetInput{},
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", common.LayoutHandler(pages.HomePage(input)))
	http.Handle("/about", common.LayoutHandler(pages.AboutPage()))
	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
