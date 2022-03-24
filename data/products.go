package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string	`json:"_"`
	UpdatedOn   string	`json:"_"`
	DeletedOn   string  `json:"-"`
}




type Products []*Product

func GetProducts() Products {
	return productList
}

func (p*Products) ToJson(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error{
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *Product)  {
	p.ID = getNextID()
	productList = append(productList,p)
}

func getNextID() int {
	lp := productList[len(productList) - 1]

	return lp.ID + 1
}


var productList = []*Product{

	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milk coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	&Product{
		ID:          2,
		Name:        "Mccoffee",
		Description: "Coffe without milk",
		Price:       1.99,
		SKU:         "efg456",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
