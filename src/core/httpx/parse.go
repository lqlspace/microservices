package httpx

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	ContentType = "Content-Type"
	ApplicationJson = "application/json"

	maxBodyLen = 8 << 20
)

func Parse(r *http.Request, v interface{}, params map[string]string) error {
	if params != nil {
		if err := parseParams(r, params); err != nil {
			return err
		}
	}

	if err := parseBody(r, v); err != nil {
		return err
	}

	return nil
}

func parseParams(r *http.Request, params map[string]string) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	for name := range r.Form {
		formValue := r.Form.Get(name)
		if len(formValue) > 0 {
			params[name] = formValue
		}
	}

	return nil
}

func parseBody(r *http.Request, v interface{}) error {
	if !withJsonBody(r) {
		return nil
	}

	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(v)
}

func withJsonBody(r *http.Request) bool {
	return r.ContentLength > 0 && strings.Contains(r.Header.Get(ContentType), ApplicationJson)
}
