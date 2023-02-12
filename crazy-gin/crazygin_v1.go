package crazy_gin

//import (
//	"fmt"
//	"net/http"
//)
//
//type HandlerFunc func(w http.ResponseWriter, r *http.Request)
//
//type Engine struct {
//	router map[string]HandlerFunc
//}
//
//func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	key := r.Method + "-" + r.URL.Path
//	if handler, ok := engine.router[key]; ok {
//		handler(w, r)
//	} else {
//		fmt.Printf("404 NOT FOUND PATH = %s\n", key)
//	}
//}
//
//func New() *Engine {
//	return &Engine{router: make(map[string]HandlerFunc)}
//}
//
//func (engine *Engine) AddRoute(method, path string, handlerFunc HandlerFunc) *Engine {
//	key := method + "-" + path
//	engine.router[key] = handlerFunc
//	return engine
//}
//
//func (engine *Engine) Get(path string, handlerFunc HandlerFunc) {
//	engine.AddRoute(http.MethodGet, path, handlerFunc)
//}
//
//func (engine *Engine) Post(path string, handlerFunc HandlerFunc) {
//	engine.AddRoute(http.MethodPost, path, handlerFunc)
//}
//
//func (engine *Engine) Run(addrAndPort string) error {
//	fmt.Printf("crazy-gin is running at %s \n", addrAndPort)
//	return http.ListenAndServe(addrAndPort, engine)
//}
