package client

type ResponseResult struct {
	statusCode int
	status     string
	body       []byte
}

func (r *ResponseResult) GetStatusCode() int {
	return r.statusCode
}

func (r *ResponseResult) GetStatus() string {
	return r.status
}

func (r *ResponseResult) GetBody() []byte {
	return r.body
}
