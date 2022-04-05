package helper

type (
	Response struct{
		User		User		`json:"user"`
		Items		[]Item		`json:"items"`
		Total		int		`json:"total"`
		Discount	int		`json:"discount"`
		GrandTotal	int		`json:"grand_total"`
		Reason		string		`json:"reason"`
	}

	User struct{
		Name	string	`json:"name"`
		Email	string	`json:"email"`
		Address	string	`json:"address"`
	}

	Item struct{
		Sku	string	`json:"sku"`
		Name	string	`json:"name"`
		Qty	int	`json:"qty"`
	}
)