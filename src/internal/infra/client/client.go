package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func Get(ctx context.Context, uri string) (response ResponseResult, err error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return response, err
	}

	response, err = send(request)
	if err != nil {
		return response, err
	}

	statusCode := response.GetStatusCode()

	if statusCode > 399 {
		return response, fmt.Errorf("ERR : Fala na requisição : %d - %s", statusCode, response.GetStatus())
	}

	return response, err
}

func send(request *http.Request) (result ResponseResult, err error) {
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	result.statusCode = response.StatusCode
	result.status = response.Status

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}

	result.body = responseBody

	return result, err
}
