package endpoint

type DataType string

const (
	ApplicationJSON DataType = "application/json"
)

const (
	TextPlain DataType = "text/plain"
)

func (dt DataType) ToString() string {
	return string(dt)
}
