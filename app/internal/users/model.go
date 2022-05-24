package users

type User struct {
	Id           string `json:"id" bson:"_id,omitempty"`
	FirstName    string `json:"firstName" bson:"firstName"`
	LastName     string `json:"lastName" bson:"lastName"`
	MiddleName   string `json:"middleName" bson:"middleName,omitempty"`
	Age          int64  `json:"age" bson:"age"`
	IsMale       bool   `json:"IsMale" bson:"isMale"`
	Status       string `json:"status" bson:"status"`
	PasswordHash string `json:"-" bson:"password"`
}

type CreateUserDTO struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	Age        int64  `json:"age"`
	IsMale     bool   `json:"isMale"`
	Status     string `json:"status"`
	Password   string `json:"password"`
}
