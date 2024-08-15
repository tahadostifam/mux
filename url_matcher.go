package mux

import (
	"errors"
	"net/http"
)

type paramsGetter struct{}
type ParamsMap = map[string]string

func Params(r *http.Request) ParamsMap {
	return r.Context().Value(paramsGetter{}).(ParamsMap)
}

func urlMatchesPattern(pattern, url string, paramsMap ParamsMap) (bool, ParamsMap, error) {
	if pattern == "" || url == "" {
		return false, nil, errors.New("both pattern and url required")
	}

	patternLen := len(pattern)
	urlLen := len(url)

	var i, j, paramStart int

	// Clear the map without reallocating
	for k := range paramsMap {
		delete(paramsMap, k)
	}

	for i < patternLen && j < urlLen {
		if pattern[i] == '{' {
			paramStart = i + 1
			for i < patternLen && pattern[i] != '}' {
				i++
			}
			if i == patternLen {
				return false, nil, errors.New("invalid pattern: unclosed '{'")
			}
			key := pattern[paramStart:i]
			valueStart := j
			for j < urlLen && url[j] != '/' {
				j++
			}
			paramsMap[key] = url[valueStart:j]
			i++ // Skip closing '}'
		} else if pattern[i] == '/' {
			if url[j] != '/' {
				return false, nil, nil
			}
			i++
			j++
		} else {
			if i == patternLen || j == urlLen || pattern[i] != url[j] {
				return false, nil, nil
			}
			i++
			j++
		}
	}

	if i != patternLen || j != urlLen {
		return false, nil, nil
	}

	return true, paramsMap, nil
}
