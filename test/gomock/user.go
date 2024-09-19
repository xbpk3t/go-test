package gomock

//go:generate mockgen -source user.go -destination user_mock.go -package gomock
type UserInterface interface {
	FindOne(id int) (*UserStruct, error)
}

type UserStruct struct {
	Name string
}
