package users

type User struct {
	Id         int    `json:"id"`
	FirstName  string `json:"first_Name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Age        int64  `json:"age"`
	IsMale     bool   `json:"is_male"`
	Status     string `json:"status"`
}
type ErrorMsg struct {
	Msg string `json:"msg"`
	Err string `json:"err"`
}
