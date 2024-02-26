package catalog

type (
	Image struct {
		ImageUrl string
	}
	Installment struct {
		NumberOfInstallments int     `graphql:"NumberOfInstallments"`
		Value                float32 `graphql:"Value"`
	}
	Discount struct {
		Name string
	}
	CommertialOffer struct {
		ListPrice          float32       `graphql:"ListPrice"`
		Price              float32       `graphql:"Price"`
		AvailableQuantity  int           `graphql:"AvailableQuantity"`
		Installments       []Installment `graphql:"Installments"`
		DiscountHighlights []Discount
	}
	Variation struct {
		Name   string
		Values []string
	}
	Seller struct {
		SellerId        string
		CommertialOffer CommertialOffer
	}
	ItemFragment struct {
		ItemId     string
		Name       string
		Images     []Image
		Variations []Variation
		Sellers    []Seller
	}
	Category struct {
		Href string
		Name string
	}
	Property struct {
		Name   string
		Values []string
	}
	ClusterHighlight struct {
		Id   string
		Name string
	}
	ProductFragment struct {
		ProductId         string
		ProductName       string
		Brand             string
		ProductReference  string
		Description       string
		Items             []ItemFragment
		CategoryTree      []Category
		Properties        []Property
		ClusterHighlights []ClusterHighlight
	}
)
