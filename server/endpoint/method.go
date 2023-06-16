package endpoint

type Method string

const (
	GET    Method = "GET"
	POST          = "POST"
	PUT           = "PUT"
	PATCH         = "PATCH"
	DELETE        = "DELETE"
)

func (m Method) ToString() string {
	return string(m)
}
