package nk;

type Router struct {
    routes map[string]*Route
}

func (r *Router) Add(name string, route *Route) {
	r.routes[name] = route
}

func (r Router) Match(method string, url string) (*Route, map[string]string) {
    for _, route := range r.routes {
        if ok, p := route.Match(method, url); ok {
            return route, p
        }
    }
    return nil, nil
}

func NewRouter() *Router {
    return &Router {make(map[string]*Route)}
}