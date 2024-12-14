package core

type IModule interface {
	Register(*GoBatt)
	Init(*GoBatt)
}
