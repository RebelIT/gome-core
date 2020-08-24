package httpRequest

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

func Post(url string, body []byte, headers map[string]string) (response http.Response, err error) {
	ctx, cncl := context.WithTimeout(context.Background(), time.Second*5)
	defer cncl()

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return http.Response{}, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return http.Response{}, err
	}

	return *resp, nil
}

func Put(url string, body []byte, headers map[string]string) (response http.Response, err error) {
	ctx, cncl := context.WithTimeout(context.Background(), time.Second*5)
	defer cncl()

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return http.Response{}, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return http.Response{}, err
	}

	return *resp, nil
}

func Delete(url string, body []byte, headers map[string]string) (response http.Response, err error) {
	ctx, cncl := context.WithTimeout(context.Background(), time.Second*5)
	defer cncl()

	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	if err != nil {
		return http.Response{}, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return http.Response{}, err
	}

	return *resp, nil
}

func Get(url string, headers map[string]string) (response http.Response, error error) {
	ctx, cncl := context.WithTimeout(context.Background(), time.Second*5)
	defer cncl()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return http.Response{}, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return http.Response{}, err
	}

	return *resp, nil
}
