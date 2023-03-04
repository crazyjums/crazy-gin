package crazy_gin

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strings"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup // share all groups
	htmlTemplates *template.Template
	funcMap       template.FuncMap
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHtmlGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob(pattern))
}

func (g *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))

	return func(c *Context) {
		file := c.Param("filepath")

		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Resp)
	}
}

func (g *RouterGroup) Static(relativePath, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	g.GET(urlPattern, handler)
}

func (g *RouterGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc

	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	c := newContext(w, r)
	c.handlers = middlewares
	c.engine = engine
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
