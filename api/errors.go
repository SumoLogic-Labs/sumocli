package api

type ResponseError struct {
	Id     string   `json:"id"`
	Errors []Errors `json:"errors"`
}

type Errors struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
