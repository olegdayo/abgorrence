package endpoint

import (
	"fmt"
	"github.com/offluck/abgorrence/common/models/endpoint/datatype"
)

type Endpoint struct {
	ID       string
	BaseURL  string
	Path     string
	Method   Method
	DataType datatype.DataType
}

func New(id string, baseURl string, path string, method Method, dataType datatype.DataType) Endpoint {
	return Endpoint{
		ID:       id,
		BaseURL:  baseURl,
		Path:     path,
		Method:   method,
		DataType: dataType,
	}
}

func (e Endpoint) GetRelation() string {
	return fmt.Sprintf(
		`<%s%s;rel="@%s";type="%s"`,
		e.BaseURL,
		e.Path,
		e.ID,
		e.DataType,
	)
}
