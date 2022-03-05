package data

type testUser struct {
	Email    string
	Username string
	Phone    string
	Password string
}

func TestUsers() []testUser {
	return []testUser{
		{Email: "aaaaa@email.com", Username: "a", Password: "1234"},
		{Email: "bbbbb@email.com", Username: "b", Password: "1234"},
		{Phone: "+1 444 444 4446", Username: "c", Password: "1234"},
		{Phone: "+1 444 444 4447", Username: "d", Password: "1234"},
	}
}
