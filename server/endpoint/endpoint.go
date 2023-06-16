package endpoint

type Endpoint struct {
	URL    string
	Method Method
}

func New(url string, method Method) Endpoint {
	return Endpoint{
		URL:    url,
		Method: method,
	}
}
