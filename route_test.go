package nk;

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewRoute(t *testing.T) {
    assert := assert.New(t)
    assert.NotNil(NewRoute("GET", "/", nil), "NewRoute should return *Route")
}

func TestRouteStatic(t *testing.T) {
    assert := assert.New(t)
    r := NewRoute("GET", "/", nil)
    e, _ := r.Match("GET", "/")
    assert.Equal(e, true)
}

func TestRouteParam(t *testing.T) {
    assert := assert.New(t)
    r := NewRoute("GET", "/users/:id", nil)
    e, m := r.Match("GET", "/users/1")
    assert.Equal(e, true)
    assert.Equal(m["id"], "1")
}

func TestRouteTwoParam(t *testing.T) {
    assert := assert.New(t)
    r := NewRoute("GET", "/users/:uid/files/:fid", nil)
    e, m := r.Match("GET", "/users/1/files/2")
    assert.Equal(e, true)
    assert.Equal(m["uid"], "1")
    assert.Equal(m["fid"], "2")
}