package prometheus

import (
	"context"
	"net/http"
)

func NewApiC(cli Client) *apiClient {
	return &apiClient{Client: cli}
}

func (c apiClient) Do(ctx context.Context, req *http.Request) (*http.Response, []byte, error) {
	return c.do(ctx, req)
}
