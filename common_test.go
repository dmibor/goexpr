package goexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func msgpacked(t *testing.T, e Expr) Expr {
	b, err := msgpack.Marshal(e)
	if !assert.NoError(t, err) {
		return e
	}
	var e2 interface{}
	err = msgpack.Unmarshal(b, &e2)
	if !assert.NoError(t, err) {
		return e
	}
	return e2.(Expr)
}
