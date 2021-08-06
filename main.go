package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	beego "web-servers/beego/server"
	buffalo "web-servers/buffalo/server"
	echo "web-servers/echo/server"
	fasthttp "web-servers/fasthttp/server"
	gin "web-servers/gin/server"
	gocraft "web-servers/gocraft/server"
	goji "web-servers/goji/server"
	httprouter "web-servers/httprouter/server"
	iris "web-servers/iris/server"
	martini "web-servers/martini/server"
	mux "web-servers/mux/server"
)

type server struct {
	Port    int
	Handler func(port int) error
}

var (
	servers = map[string]*server{
		"http & mux":                  &server{Port: 8081, Handler: mux.Run},
		"gin":                         &server{Port: 8082, Handler: gin.Run},
		"beego":                       &server{Port: 8083, Handler: beego.Run},
		"echo":                        &server{Port: 8084, Handler: echo.Run},
		"martini & martini-render":    &server{Port: 8085, Handler: martini.Run},
		"fasthttp & fasthttp-routing": &server{Port: 8086, Handler: fasthttp.Run},
		"iris":                        &server{Port: 8087, Handler: iris.Run},
		"revel":                       &server{Port: 8088, Handler: nil},
		"buffalo":                     &server{Port: 8089, Handler: buffalo.Run},
		"goji":                        &server{Port: 8090, Handler: goji.Run},
		"gocraft":                     &server{Port: 8091, Handler: gocraft.Run},
		"httprouter":                  &server{Port: 8092, Handler: httprouter.Run},
	}
)

func main() {
	numTimes := 200
	numGoRoutines := 1

	// start servers
	var err error
	for name, conf := range servers {
		if conf.Handler == nil {
			continue
		}

		log.Printf("starting %s server", name)
		go conf.Handler(conf.Port)
	}

	// create output file
	log.Printf("create output file")
	file, err := createFile(".", "generated", "text")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// run test
	for name, conf := range servers {
		log.Printf("running tests on %s", name)
		if conf.Handler == nil {
			continue
		}

		if _, err = file.WriteString(fmt.Sprintf(":: %s\n", name)); err != nil {
			panic(err)
		}

		elapsedTime := call(name, conf.Port, numTimes, numGoRoutines)
		if _, err = file.WriteString(fmt.Sprintf("Elapsed time: %f\n\n", elapsedTime.Seconds())); err != nil {
			panic(err)
		}
	}
}

func call(name string, port, numTimes, numGoRoutines int) time.Duration {
	start := time.Now()

	for i := 0; i <= numTimes; i++ {
		wg := &sync.WaitGroup{}
		for i := 0; i <= numGoRoutines; i++ {
			f := func(id int, wg *sync.WaitGroup) {
				wg.Add(1)
				defer wg.Done()

				url := fmt.Sprintf("http://localhost:%d/v1/persons/%d/addresses/%d", port, id, id)
				_, err := http.Get(url)
				if err != nil {
					log.Printf("\n(%s) error: %s", name, err.Error())
					return
				}
			}

			f(i, wg)
		}
		wg.Wait()

	}

	return time.Since(start)
}

func createFile(folder, name, extension string) (file *os.File, err error) {
	fileName := fmt.Sprintf("%s/%s.%s", folder, name, extension)

	file, err = os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, err
}
