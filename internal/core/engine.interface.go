package core

type IEngine interface {
	GET(string, func(IContext))
	Run()
}
