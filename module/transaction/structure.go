package transaction

type(

	Transaction struct{
		ID		int	`json:"id"`
		UserID		int	`json:"user_id"`
		Datetime	string	`json:"datetime"`
		Total		int	`json:"total"`
		Discount	int	`json:"discount"`
		GrandTotal	int	`json:"grand_total"`
	}

	Transactions []Transaction

	TransactionDetail struct{
		ID		int	`json:"id"`
		TransactionID	int	`json:"transaction_id"`
		Price		int	`json:"price"`
		Qty		int	`json:"qty"`
		SubTotal	int	`json:"sub_total"`
	}

	TransactionDetails []TransactionDetail

	Item struct{
		Sku	string	`json:"sku"`
		Qty	int	`json:"qty"`
	}

	Pembelian struct{
		UserID	int	`json:"user_id"`
		Items	[]Item	`json:"items"`
	}
)