package ginger

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

const tag_uri = "uri"
const tag_json = "json"
const tag_form = "form"

// GetRequest get request object from the gin context.
func GetRequest[T any](ctx *gin.Context) *T {
	request := new(T)
	tags := parseTags(request)

	var err error
	for _, tag := range tags {
		switch tag {
		case tag_json:
			err = ctx.ShouldBindJSON(request)
		case tag_form:
			err = ctx.ShouldBindQuery(request)
		case tag_uri:
			err = ctx.ShouldBindUri(request)
		}
		if err != nil {
			panic(err)
		}
	}

	return request
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
