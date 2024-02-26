package catalog

import (
	"context"
	"fmt"

	"github.com/hasura/go-graphql-client"
)

type SelectedFacetInput struct {
	Key   string
	Value string
}

type SimulationBehavior string

const (
	SimulationBehaviorDefault SimulationBehavior = "default"
	SimulationBehaviorOnly1P  SimulationBehavior = "only1P"
	SimulationBehaviorSkip    SimulationBehavior = "skip"
	SimulationBehaviorAsync   SimulationBehavior = "async"
)

type ProductSearchInput struct {
	From                 int
	To                   int
	FullText             string
	OrderBy              string
	PriceRange           string
	SalesChannel         string
	HideUnavailableItems bool
	SimulationBehavior   SimulationBehavior
	SelectedFacets       []SelectedFacetInput
}

type ContextDirective struct {
	provider string
}

func (cd ContextDirective) Type() graphql.OptionType {
	return graphql.OptionTypeOperationDirective
}

func (cd ContextDirective) String() string {
	return fmt.Sprintf("@context(provider: \"%s\")", cd.provider)
}

type ProductSearchQuery struct {
	ProductSearch struct {
		RecordsFiltered int
		Products        []ProductFragment
	} `graphql:"productSearch(from: $from, fullText: $fullText, hideUnavailableItems: $hideUnavailableItems, orderBy: $orderBy, priceRange: $priceRange, salesChannel: $salesChannel, selectedFacets: $selectedFacets, simulationBehavior: $simulationBehavior, to: $to) @context(provider: \"vtex.search-graphql\")"`
}

func ProductSearch(client *graphql.Client, input ProductSearchInput) (ProductSearchQuery, error) {

	var query = ProductSearchQuery{}

	var variables = map[string]interface{}{
		"from":                 graphql.Int(input.From),
		"to":                   graphql.Int(input.To),
		"fullText":             graphql.String(input.FullText),
		"orderBy":              graphql.String(input.OrderBy),
		"priceRange":           graphql.String(input.PriceRange),
		"salesChannel":         graphql.String(input.SalesChannel),
		"hideUnavailableItems": graphql.Boolean(input.HideUnavailableItems),
		"simulationBehavior":   input.SimulationBehavior,
		"selectedFacets":       input.SelectedFacets,
	}

	err := client.Query(context.Background(), &query, variables, graphql.OperationName("ProductSearch"), ContextDirective{provider: "vtex.search-graphql"})

	return query, err
}
