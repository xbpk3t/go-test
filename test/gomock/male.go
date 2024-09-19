package gomock

//go:generate mockgen -source=male.go -destination=male_mock.go -package=gomock
type Male interface {
	Get(id int64) error
}

type User struct {
	Person Male
}

func NewUser(male Male) *User {
	return &User{Person: male}
}

func (u *User) GetUserInfo(id int64) error {
	return u.Person.Get(id)
}
