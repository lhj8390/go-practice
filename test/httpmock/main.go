package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type RequestBody struct {
	Name string
}

type ResponseBody struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

const (
	MainUrl         = "https://example.com/test1"
	RequestParamUrl = "https://example.com/test2"
	PostUrl         = "https://example.com/test3"
)

func main() {
	resp1, err := RequestHTTP("GET")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(resp1)

	resp2, err := RequestWithParam("GET", map[string]string{"test": "param"})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(resp2)

	resp3, err := RequestWithBody("POST", &RequestBody{Name: "test"})
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(resp3)
}

func RequestHTTP(method string) (string, error) {
	client := &http.Client{
		Timeout: time.Duration(5),
	}

	req, err := http.NewRequest(method, MainUrl, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	resBody, _ := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	return string(resBody), err
}

func RequestWithParam(method string, param map[string]string) (*ResponseBody, error) {
	client := &http.Client{
		Timeout: time.Duration(5),
	}

	req, err := http.NewRequest(method, RequestParamUrl, nil)
	if err != nil {
		return nil, err
	}

	if param != nil {
		query := url.Values{}
		for k, v := range param {
			query.Set(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, _ := io.ReadAll(resp.Body)
	fmt.Println(resBody)
	var respBody *ResponseBody
	if err := json.Unmarshal(resBody, &respBody); err != nil {
		return nil, err
	}
	return respBody, err
}

func RequestWithBody(method string, body *RequestBody) (*ResponseBody, error) {
	client := &http.Client{
		Timeout: time.Duration(5),
	}

	reqBody, err := json.Marshal(body)

	req, err := http.NewRequest(method, PostUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, _ := io.ReadAll(resp.Body)
	fmt.Println(resBody)
	var respBody *ResponseBody
	if err := json.Unmarshal(resBody, &respBody); err != nil {
		return nil, err
	}
	return respBody, err
}
