package server

type RequestType int

const (
	Empty RequestType = iota
	Filled
)

type Request interface {
	Validate() error
	Type() RequestType
}
