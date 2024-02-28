package interfaces

type Module interface {
	Setup() error
	Close()
}
