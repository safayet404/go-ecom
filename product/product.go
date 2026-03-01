package product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
}
