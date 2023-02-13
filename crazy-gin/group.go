package crazy_gin

import "net/http"

type RouterGroup struct {
	engine      *Engine // all groups share a engine instance
	prefix      string
	parent      *RouterGroup  // support nesting
	middlewares []HandlerFunc //support middleware

}

func (g *RouterGroup) Group(prefix string) *RouterGroup {
	engine := g.engine
	newGroup := &RouterGroup{
		engine: engine,
		parent: g,
		prefix: g.prefix + prefix,
	}
	engine.groups = append(engine.groups, newGroup)

	return newGroup
}

func (g *RouterGroup) addRoute(method, comp string, handlerFunc HandlerFunc) {
	pattern := g.prefix + comp
	g.engine.router.AddRoute(method, pattern, handlerFunc)
}

func (g *RouterGroup) GET(path string, handlerFunc HandlerFunc) {
	g.addRoute(http.MethodGet, path, handlerFunc)
}

func (g *RouterGroup) POST(path string, handlerFunc HandlerFunc) {
	g.addRoute(http.MethodPost, path, handlerFunc)
}
