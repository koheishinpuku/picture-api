package Models

type Userdata struct {
	Id       int    `json:id`
	Username string `json:table_name`
	Email    string `json:email`
	Password string `json:password`
}
