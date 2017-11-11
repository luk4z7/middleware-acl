package error

type ErrorsAPI struct {
	Errors []Errors `json:"errors"`
	Url    string   `json:"url"`
	Method string   `json:"method"`
}

type Errors struct {
	ParameterName string `json:"parameter_name"`
	Type          string `json:"type"`
	Message       string `json:"message"`
}

type Err struct {
	Name string
}

func (e *Err) Error() string {
	return e.Name
}