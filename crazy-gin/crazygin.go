package crazy_gin

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup // share all groups
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//c := newContext(w, r)
	//key := c.Method + "-" + c.Path
	//fmt.Println("ServeHTTP, handlers = ", engine.router.handlers)
	//if handler, ok := engine.router.handlers[key]; ok {
	//	handler(c)
	//} else {
	//	fmt.Printf("404 NOT FOUND PATH = %s\n", key)
	//}

	c := newContext(w, r)
	engine.router.handle(c)
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) AddRoute(method, path string, handlerFunc HandlerFunc) {
	engine.router.AddRoute(method, path, handlerFunc)
}

func (engine *Engine) GET(path string, handlerFunc HandlerFunc) {
	engine.router.AddRoute(http.MethodGet, path, handlerFunc)
}

func (engine *Engine) POST(path string, handlerFunc HandlerFunc) {
	engine.router.AddRoute(http.MethodPost, path, handlerFunc)
}

func (engine *Engine) Run(addrAndPort string) error {
	fmt.Printf("crazy-gin is running at %s \n", addrAndPort)
	return http.ListenAndServe(addrAndPort, engine)
}
