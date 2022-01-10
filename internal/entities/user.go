package entities

type User struct {
	ID     int
	Type   int
	Name   string
	Email  string
	Role   int
	Status bool
}

func GetUserByID(ID int) *User {
	return nil
}

func CreateNewUser(name string, email string, password string, role int) int {
	return -1
}

func DeleteUserByID(ID int) {

}

func UpdateUser(user *User) {

}

func Login(email string, password string) *User {
	return nil
}
