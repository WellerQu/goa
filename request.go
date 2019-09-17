package goa

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

// Request 请求对象
type Request struct {
	*http.Request
	RoutePath string
}

// Queries 获取查询字符串中的请求参数
func (r *Request) Queries() (url.Values, error) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return nil, err
	}

	return queryForm, nil
}

// Params 路径参数
type Params map[string]string

// Params 获取路径字符串中的请求参数
func (r *Request) Params() (Params, error) {
	if r.RoutePath == "" {
		return nil, errors.New("缺少RoutePath值")
	}

	preCompileRegExp := regexp.MustCompile(`:([^/]+)`)
	regExpStr := preCompileRegExp.ReplaceAllString(r.RoutePath, "(?P<${1}>[^/]+)")
	routeRegExp := regexp.MustCompile(regExpStr)

	if !routeRegExp.MatchString(r.URL.Path) {
		return nil, fmt.Errorf("请求路径%s 与 路由%s 不匹配", r.URL.Path, r.RoutePath)
	}

	params := make(Params)
	names := routeRegExp.SubexpNames()
	matches := routeRegExp.FindStringSubmatch(r.URL.Path)

	for i, n := range names {
		if i > 0 && n != "" {
			params[n] = matches[i]
		}
	}

	return params, nil
}

// JSON 获取以JSON格式发送的请求参数
func (r *Request) JSON(m interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(m)
}
