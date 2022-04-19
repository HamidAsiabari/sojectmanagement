package webserver

import (
	"log"
	"sync"

	"github.com/HamidAsiabari/sojectmanagement/router"

	"github.com/labstack/echo/v4"
)

var initWebServer sync.Once
var ws *WebServer
var port string

type WebServer struct {
	router *echo.Echo
}

func Instance() *WebServer {
	initWebServer.Do(func() {
		ws = &WebServer{}
		ws.router = router.New()
		port = ":1323"
	})

	return ws
}

func (w *WebServer) Start() {
	// start the server

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Start HTTP server on port " + port)
		log.Fatal(w.router.Start(port))
	}()

	wg.Wait()
}
