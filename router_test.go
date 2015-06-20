package nk;

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
    assert := assert.New(t)
    assert.NotNil(NewRouter())
}

func TestRouterAdd(t *testing.T) {
    r := NewRouter()
    r.Add("main", NewRoute("GET", "/", nil))
}

func TestRouterStatic(t *testing.T) {
    assert := assert.New(t)
    r := NewRouter()
    r.Add("main", NewRoute("GET", "/", nil))
    ro, _ := r.Match("GET", "/")
    assert.NotNil(ro)
}

func TestRouterParam(t *testing.T) {
    assert := assert.New(t)
    r := NewRouter()
    r.Add("user", NewRoute("GET", "/post/:id", nil))
    ro, p := r.Match("GET", "/post/260227")
    assert.NotNil(ro)
    assert.NotNil(p)
    assert.Equal(p["id"], "260227")
}

func TestRouterTwoParam(t *testing.T) {
    assert := assert.New(t)
    r := NewRouter()
    r.Add("user", NewRoute("GET",  "/users/:uid/files/:fid", nil))
    ro, p := r.Match("GET", "/users/1/files/2")
    assert.NotNil(ro)
    assert.NotNil(p)
    assert.Equal(p["uid"], "1")
    assert.Equal(p["fid"], "2")
}

func GetRouter() *Router {
    r := NewRouter()
    r.Add("main", NewRoute("GET", "/", nil))
    r.Add("all", NewRoute("GET", "/all", nil))
    r.Add("search", NewRoute("GET", "/search", nil))
    r.Add("hub", NewRoute("GET", "/hub/:name", nil))
    r.Add("hub/type", NewRoute("GET", "/hub/:name/:type", nil))
    r.Add("post", NewRoute("GET", "/post/:id", nil))
    return r
}

func BenchmarkMain(b *testing.B) {
    r := GetRouter()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        r.Match("GET", "/")
    }
}

func BenchmarkAll(b *testing.B) {
    r := GetRouter()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        r.Match("GET", "/all")
    }
}

func BenchmarkSearch(b *testing.B) {
    r := GetRouter()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        r.Match("GET", "/search")
    }
}

func BenchmarkHub(b *testing.B) {
    r := GetRouter()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        r.Match("GET", "/hub/programming")
    }
}

func BenchmarkHubType(b *testing.B) {
    r := GetRouter()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        r.Match("GET", "/hub/programming/all")
    }
}

func BenchmarkPost(b *testing.B) {
    r := GetRouter()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        r.Match("GET", "/post/260227")
    }
}