package brand

type(
	Brand struct{
		ID	int	`json:"id"`
		Name	string	`json:"name"`
	}

	Brands []Brand
)