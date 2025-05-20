package api

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// BuildQueryParams は任意の構造体から URL クエリパラメータを生成する
func BuildQueryParams[T any](opts ...T) (url.Values, error) {
	if len(opts) == 0 {
		return url.Values{}, nil
	}
	return query.Values(opts[0])
}
