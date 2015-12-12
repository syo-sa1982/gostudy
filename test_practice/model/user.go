package test_practice
import "strings"


type User struct {
	FirstName string
	LastName string
	Divisions []string
}

func NewUser(name string) *User {
	names := strings.Split(name, " ")
	if len(names) >= 2 {
		return &User{FirstName: names[0], LastName: names[1]}
	} else {
		return &User{FirstName: names[0]}
	}

}

func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}

func (user *User) AddDivision(division string) *User {
	user.Divisions = append(user.Divisions, division)
	return user
}