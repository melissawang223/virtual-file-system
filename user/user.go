package user

var Users []User

type User struct {
	Name    string
	Folders map[string]string
}
