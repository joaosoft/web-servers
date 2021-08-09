package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
	"web-servers/implementation/models"
	iris "web-servers/iris/server"
	martini "web-servers/martini/server"
	mux "web-servers/mux/server"
	//revel "web-servers/revel/app/server"
	//web "web-servers/web/server"
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
		//"revel":                       &server{Port: 8088, Handler: revel.Run},
		"buffalo":    &server{Port: 8089, Handler: buffalo.Run},
		"goji":       &server{Port: 8090, Handler: goji.Run},
		"gocraft":    &server{Port: 8091, Handler: gocraft.Run},
		"httprouter": &server{Port: 8092, Handler: httprouter.Run},
		//"web":        &server{Port: 8093, Handler: web.Run},
	}
)

func main() {
	numRequests := 200
	numGoRoutines := 5

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
	<-time.After(time.Second * 5)
	for name, conf := range servers {
		log.Printf("running tests on %s", name)
		if conf.Handler == nil {
			continue
		}

		if _, err = file.WriteString(fmt.Sprintf(":: %s\n", name)); err != nil {
			panic(err)
		}

		elapsedTime := call(name, conf.Port, numGoRoutines, numRequests)
		if _, err = file.WriteString(fmt.Sprintf("Elapsed time: %f\n\n", elapsedTime.Seconds())); err != nil {
			panic(err)
		}

		file.Sync()
		<-time.After(time.Second * 1)
	}
}

func call(name string, port, numGoRoutines, numRequests int) time.Duration {
	start := time.Now()
	wg := &sync.WaitGroup{}

	for i := 0; i <= numGoRoutines; i++ {
		f := func(name string, id int, wg *sync.WaitGroup, numRequests int) {
			wg.Add(1)
			defer wg.Done()

			for i := 0; i <= numRequests; i++ {
				url := fmt.Sprintf("http://localhost:%d/v1/persons/%d/addresses/%d", port, id, i+1)
				response, err := http.Get(url)
				if err != nil {
					log.Printf("\nERROR 1 (name: %s : request: %d | %d) error: %s", name, id, i+1, err.Error())
					return
				}

				if response != nil {
					defer response.Body.Close()
					bodyResponse, err := ioutil.ReadAll(response.Body)
					if err != nil {
						log.Printf("\nERROR 2 (name: %s : request: %d | %d) error: %s", name, id, i+1, err.Error())
						return
					}

					var address models.Address
					if err = json.Unmarshal(bodyResponse, &address); err != nil {
						log.Printf("\nERROR 3 (name: %s : request: %d | %d) error: %s", name, id, i+1, err.Error())
						log.Printf(string(bodyResponse))
						return
					}

					if address.Id != strconv.Itoa(i+1) {
						log.Printf("\nERROR 4 (name: %s : request: %d | %d) error: %s", name, id, i+1, err.Error())
						log.Printf(string(bodyResponse))
						return
					}
				}
			}
		}

		go f(name, i+1, wg, numRequests)
	}

	wg.Wait()

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
