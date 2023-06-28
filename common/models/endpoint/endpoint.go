package endpoint

import "fmt"

type Endpoint struct {
	ID       string
	BaseURL  string
	URL      string
	Method   Method
	DataType DataType
}

func New(id string, baseURl string, url string, method Method, dataType DataType) Endpoint {
	return Endpoint{
		ID:       id,
		BaseURL:  baseURl,
		URL:      url,
		Method:   method,
		DataType: dataType,
	}
}

func (e Endpoint) GetRelation() string {
	return fmt.Sprintf(
		`<%s%s;rel="@%s";type="%s"`,
		e.BaseURL,
		e.URL,
		e.ID,
		e.DataType,
	)
}
