package main

type Response struct {
	Result string `json:"result"`
}

func NewResponse(result string) *Response {
	return &Response{
		Result: result,
	}
}
