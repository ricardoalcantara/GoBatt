package core

type IEngine interface {
	Get(string, func(IContext))
	Post(string, func(IContext))
	Put(string, func(IContext))
	Delete(string, func(IContext))
	Patch(string, func(IContext))
	Start()
}
