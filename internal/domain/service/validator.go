package service

type IValidator interface {
	Struct(s interface{}) error
	Var(field interface{}, tag string) error
}
