package user

type(
	User struct{
		ID	int	`json:"id"`
		Name	string	`json:"name"`
		Email	string	`json:"email"`
		Address	string	`json:"address"`
	}
	Users []User
)