package pkg

import (
	"errors"
	"net/url"
	"reflect"
	"strings"
)

// argument should be a struct
func BuildQueryParams(x any) (string, error) {
	v := reflect.ValueOf(x)

	if v.Kind() != reflect.Struct {
		return "", errors.New("x must be a struct")
	}

	params := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		tags := strings.Split(v.Type().Field(i).Tag.Get("json"), ",")
		if !v.Field(i).IsZero() || !isOmitempty(tags) {
			params.Add(tags[0], v.Field(i).String())
		}
	}

	return params.Encode(), nil
}

// check if its have tag omitempty
func isOmitempty(tags []string) bool {
	for _, tag := range tags {
		if tag == "omitempty" {
			return true
		}
	}
	return false
}
