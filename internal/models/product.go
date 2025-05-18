// package models

// import (
// 	"encoding/json"
// )

// type Products struct {
// 	Count    int         `json:"count"`
// 	Total    int         `json:"total"`
// 	// Products []*Products `json:"products"`
// 	Products []*Product `json:"products"`
// }

// type Product struct {
// 	ID               int    `json:"id"`
// 	// ID               string    `json:"id"`
// 	SKU              string    `json:"sku"`
// 	ShortDescription *string   `json:"shortDescription"`
// 	LongDescription  *string   `json:"longDescription"`
// 	Brand            string    `json:"brand"`
// 	Price            *Money    `json:"price"`
// 	Features         []string  `json:"features,omitempty"`
// 	Images 			 ImageList `json:"images"`
// 	// Images           []*Image  `json:"images"`
// 	Options          []*Option `json:"options"`
// 	Variants         []*Variant  `json:"variants"`
// }

// type Variant struct {
// 	ID    int    `json:"id"`
// 	EAN   string `json:"ean"`
// 	SKU   string `json:"sku"`
// 	Prices struct {
// 		Price           float64     `json:"price"`
// 		MembershipPrice interface{} `json:"membershipPrice"` // float or empty string/null
// 		CurrencyCode    string      `json:"currencyCode"`
// 	} `json:"prices"`
// 	Inventory struct {
// 		Count     interface{} `json:"count"` // Sometimes null
// 		IsInStock bool        `json:"isInStock"`
// 	} `json:"inventory"`
// 	Options []struct {
// 		ID    int `json:"id"`
// 		Value []struct {
// 			Label        string `json:"label"`
// 			PreorderDate string `json:"preorderDate"`
// 		} `json:"value"`
// 		Group string `json:"group"`
// 	} `json:"options"`
// }


// type Money struct {
// 	Amount   string `json:"amount"`
// 	Currency string `json:"currency"`
// }

// type Image struct {
// 	Url     string  `json:"url"`
// 	AltText *string `json:"altText"`
// }

// type Option struct {
// 	ID      string `json:"id"`
// 	SKU     string `json:"sku"`
// 	Price   *Money `json:"price"`
// 	InStock bool   `json:"inStock"`
// 	Size    string `json:"size"`
// 	Colour  string `json:"colour"`
// }

// type ImageList []*Image

// func (il *ImageList) UnmarshalJSON(data []byte) error {
// 	var urls []string
// 	if err := json.Unmarshal(data, &urls); err != nil {
// 		return err
// 	}

// 	var result []*Image
// 	for _, u := range urls {
// 		result = append(result, &Image{
// 			Url:     u,
// 			AltText: nil,
// 		})
// 	}
// 	*il = result
// 	return nil
// }

package models

type Products struct {
	Count    int        `json:"count"`
	Total    int        `json:"total"`
	Products []*Product `json:"products"`
}

type Product struct {
	ID               int       `json:"id"`
	SKU              string    `json:"sku"`
	ShortDescription *string   `json:"shortDescription"`
	LongDescription  *string   `json:"longDescription"`
	Brand            string    `json:"brand"`
	Price            *Money    `json:"price"`
	Features         []string  `json:"features,omitempty"`
	Images           []*Image  `json:"images"`
	Options          []*Option `json:"options"`
	Variants         []*Variant  `json:"variants"`
}

type Variant struct {
	ID    int    `json:"id"`
	EAN   string `json:"ean"`
	SKU   string `json:"sku"`
	Prices struct {
		Price           float64     `json:"price"`
		MembershipPrice interface{} `json:"membershipPrice"`
		CurrencyCode    string      `json:"currencyCode"`
	} `json:"prices"`
	Inventory struct {
		Count     interface{} `json:"count"`
		IsInStock bool        `json:"isInStock"`
	} `json:"inventory"`
	Options []struct {
		ID    int `json:"id"`
		Value []struct {
			Label        string `json:"label"`
			PreorderDate string `json:"preorderDate"`
		} `json:"value"`
		Group string `json:"group"`
	} `json:"options"`
}

type Money struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Image struct {
	Url     string  `json:"url"`
	AltText *string `json:"altText"`
}

type Option struct {
	ID      string `json:"id"`
	SKU     string `json:"sku"`
	Price   *Money `json:"price"`
	InStock bool   `json:"inStock"`
	Size    string `json:"size"`
	Colour  string `json:"colour"`
}

