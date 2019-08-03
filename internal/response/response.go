package response

type ResponseError struct {
	Message string `json:"message"`
}

type ResponsePayload struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
	Error *ResponseError `json:"error"`
}

func Create(data interface{}, err error) *ResponsePayload {
	var apiErr *ResponseError

	if err != nil {
		apiErr = &ResponseError{
			Message: err.Error(),
		}
	}

	return &ResponsePayload{
		Data: data,
		Error: apiErr,
	}
}