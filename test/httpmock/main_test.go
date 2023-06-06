package main

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type Args struct {
	method string
	urls   string
	param  map[string]string
	body   interface{}
}

func TestHttpMockAPI(t *testing.T) {
	// block all HTTP requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// GET
	httpmock.RegisterResponder(
		"GET",
		"https://example.com/test1",
		httpmock.NewStringResponder(200, "hello!"),
	)

	// GET with param
	httpmock.RegisterResponder(
		"GET",
		"https://example.com/test2?test=param",
		func(request *http.Request) (*http.Response, error) {
			responseBody := map[string]interface{}{"id": 1, "name": "A"}
			return httpmock.NewJsonResponse(200, responseBody)
		},
	)

	// POST with body
	httpmock.RegisterResponder(
		"POST",
		"https://example.com/test3",
		func(request *http.Request) (*http.Response, error) {
			body := make(map[string]interface{})
			if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
				return httpmock.NewStringResponse(400, ""), nil
			}
			return httpmock.NewJsonResponse(200, map[string]interface{}{"Name": "test"})
		},
	)

	args1 := Args{
		"GET",
		"https://example.com/test1",
		nil,
		nil,
	}

	t.Run("method GET 기본 테스트", func(t *testing.T) {
		got, err := RequestHTTP(args1.method)
		assert.Equalf(t, "hello!", got, "RequestHTTP(%v)", args1.method)
		assert.Equal(t, nil, err)
	})

	args2 := Args{
		"GET",
		"https://example.com/test2",
		map[string]string{"test": "param"},
		nil,
	}
	t.Run("method GET param 테스트", func(t *testing.T) {
		got, err := RequestWithParam(args2.method, args2.param)
		assert.Equalf(t, &ResponseBody{Id: 1, Name: "A"}, got, "RequestWithParam(%v, %v)", args2.method, args2.param)
		assert.Equal(t, nil, err)
	})

	args3 := Args{
		"POST",
		"https://example.com/test3",
		nil,
		&RequestBody{
			Name: "test",
		},
	}
	t.Run("method POST body 테스트", func(t *testing.T) {
		got, err := RequestWithBody(args3.method, args3.body.(*RequestBody))
		assert.Equalf(t, &ResponseBody{Name: "test"}, got, "RequestWithBody(%v, %v)", args3.method, args3.body)
		assert.Equal(t, nil, err)
	})

}
