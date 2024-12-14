package core

type IContext interface {
	Status(int)
	String(string)
	Json(any)
}
