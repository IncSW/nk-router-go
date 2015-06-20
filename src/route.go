package nk;

import (
    "regexp"
    "strings"
)

var r, _ = regexp.Compile(":([_A-Za-z]+)")

type Route struct {
    method string
    path string
    pattern *regexp.Regexp
    parameters []string
    handler func()
}

func (r Route) Match(method string, url string) (bool, map[string]string) {
    if r.method != method {
        return false, nil
    }
    if len(r.parameters) == 0 {
        return r.path == url, nil
    }
    if r.pattern.MatchString(url) == false {
        return false, nil
    }
    m := make(map[string]string)
    sf := r.pattern.FindAllStringSubmatch(url, -1)[0][1:];
    for i, p := range sf {
        m[r.parameters[i]] = p
    }
    return true, m
}

func NewRoute(method string, path string, handler func()) *Route {
    ps := path
    parameters := []string{}
    for _, p := range r.FindAllString(path, -1) {
        ps = strings.Replace(ps, p, "([^/]+)", 1)
        parameters = append(parameters, strings.Replace(p, ":", "", 1))
    }
    pattern, _ := regexp.Compile(strings.Join([]string {"^", ps, "$"}, ""))
    return &Route {method, path, pattern, parameters, handler}
}