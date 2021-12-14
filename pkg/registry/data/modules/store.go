package modules

type ModuleStore interface {
	Init() error
	ReadAll(limit int, offset int) ([]*Module, error)
	ReadOne(name string) (*Module, error)
}
