package ginger

import (
	"reflect"

	"github.com/WebXense/sql"
	"github.com/gin-gonic/gin"
)

// Request get request object from the gin context.
func Request[T any](ctx *gin.Context) *T {
	request := new(T)
	tags := parseTags(request)

	for _, tag := range tags {
		switch tag {
		case tag_json:
			ctx.ShouldBindJSON(request)
		case tag_form:
			ctx.ShouldBindQuery(request)
		case tag_uri:
			ctx.ShouldBindUri(request)
		}
	}

	return request
}

// PaginationRequest get pagination request from the gin context.
func PaginationRequest(ctx *gin.Context) *sql.Pagination {
	return Request[sql.Pagination](ctx)
}

// SortRequest get sort request from the gin context.
func SortRequest(ctx *gin.Context) *sql.Sort {
	return Request[sql.Sort](ctx)
}

// parseTags get the tags related to the request method
func parseTags[T any](request T) []string {
	m := make(map[string]bool)

	for i := 0; i < reflect.TypeOf(request).Elem().NumField(); i++ {
		tag := reflect.TypeOf(request).Elem().Field(i).Tag
		for _, key := range []string{tag_json, tag_form, tag_uri} {
			value := tag.Get(key)
			if len(value) > 0 {
				m[key] = true
				break
			}
		}
	}

	result := make([]string, 0)
	for key := range m {
		result = append(result, key)
	}
	return result
}
