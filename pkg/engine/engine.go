package engine

type Engine interface {
	Start()
	Ssh()
	Stop()
	Rm()
	Validate()
}
