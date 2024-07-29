package mux

import (
	"errors"
	"strings"
)

type ParamsGetter struct{}
type ParamsMap = map[string]string

func UrlMatchesPattern(pattern string, url string) (bool, ParamsMap, error) {
	if pattern == "" || url == "" {
		return false, nil, errors.New("both pattern and url required")
	}

	urlValues := strings.Split(url, "/")[1:]
	params := strings.Split(pattern, "/")[1:]

	if len(urlValues) != len(params) {
		return false, nil, nil
	}

	paramsMap := make(ParamsMap)
	for idx, param := range params {
		if strings.HasPrefix(param, "{") && strings.HasSuffix(param, "}") {
			key := param[1 : len(param)-1]
			paramsMap[key] = urlValues[idx]
			continue
		}

		if param != urlValues[idx] {
			return false, nil, nil
		}
	}

	return true, paramsMap, nil
}
