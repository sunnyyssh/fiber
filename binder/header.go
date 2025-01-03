package binder

import (
	"reflect"
	"strings"

	"github.com/gofiber/utils/v2"
	"github.com/valyala/fasthttp"
)

// v is the header binder for header request body.
type HeaderBinding struct {
	EnableSplitting bool
}

// Name returns the binding name.
func (*HeaderBinding) Name() string {
	return "header"
}

// Bind parses the request header and returns the result.
func (b *HeaderBinding) Bind(req *fasthttp.Request, out any) error {
	data := make(map[string][]string)
	req.Header.VisitAll(func(key, val []byte) {
		k := utils.UnsafeString(key)
		v := utils.UnsafeString(val)

		if b.EnableSplitting && strings.Contains(v, ",") && equalFieldType(out, reflect.Slice, k) {
			values := strings.Split(v, ",")
			for i := 0; i < len(values); i++ {
				data[k] = append(data[k], values[i])
			}
		} else {
			data[k] = append(data[k], v)
		}
	})

	return parse(b.Name(), out, data)
}

// Reset resets the HeaderBinding binder.
func (b *HeaderBinding) Reset() {
	b.EnableSplitting = false
}
