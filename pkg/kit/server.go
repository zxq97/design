package kit

type Server interface {
	Run() error
	Stop() error
}
