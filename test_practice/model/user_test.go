package test_practice
import "testing"

func TestConstructorWithFullName(t *testing.T) {
	user := NewUser("hoge aaa")

	if user.FirstName != "hoge" {
		t.Error("user's first name should be hoge")
	}

	if user.LastName != "aaa" {
		t.Error("user's last name should be aaa")
	}
}


func TestConstructorWithFirstName(t *testing.T) {
	user := NewUser("hoge")

	if user.FirstName != "hoge" {
		t.Error("user's first name should be hoge")
	}

	if user.LastName != "" {
		t.Error("user's last name should be empty")
	}
}


func TestConstructorWithEmptyString(t *testing.T) {
	user := NewUser("")

	if user.FirstName != "" {
		t.Error("user's first name should be empty")
	}
	if user.LastName != "" {
		t.Error("user's last name should be empty")
	}
}

func TestDevision(t *testing.T) {

	user := NewUser("hoge aaa")

	if len(user.Divisions) != 0 {
		t.Error("default divisions is empty slice")
	}

}

func TestFullName(t *testing.T) {

	fullname := "hoge aaa"
	user := NewUser(fullname)

	if user.FullName() != fullname {
		t.Errorf("fullname should be %s, but %s", fullname, user.FullName())
	}
}

func TestAddDevision(t *testing.T) {

	user := NewUser("hoge aaa")
	division := "test"

	user.AddDivision(division)
	if user.Divisions[0] != division {
		t.Log(user.Divisions)
		t.Errorf("%s division was not added", division)
	}

}