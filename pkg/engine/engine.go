package engine

// Engine provides an interface for wrappers around external third party
// technologies which can be used to create a temporary linux environment. For
// example, we could have a `RktEngine`, `LibvirtdEngine`, etc...
type Engine interface {
	Start()
	SSH()
	Stop()
	Rm()
	Validate()
}
