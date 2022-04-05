package product

type (
	Product struct {
		ID      int    	`json:"id"`
		BrandID int    	`json:"brand_id"`
		Sku     string 	`json:"sku"`
		Name	string 	`json:"name"`
		Qty	int	`json:"qty"`
		Price	int	`json:"price"`
	}

	ListProducts []Product

)