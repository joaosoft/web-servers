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
	"web-servers/domain/models"
	"web-servers/domain/server"
	echo "web-servers/echo/server"
	fasthttp "web-servers/fasthttp/server"
	gin "web-servers/gin/server"
	gocraft "web-servers/gocraft/server"
	goji "web-servers/goji/server"
	httprouter "web-servers/httprouter/server"
	iris "web-servers/iris/server"
	martini "web-servers/martini/server"
	mux "web-servers/mux/server"
	revel "web-servers/revel/app/server"
	//web "web-servers/web/server"
)

type Config struct {
	Enabled bool
	Port    int
	Handler func(port int) server.IServer
}

var (
	servers = map[string]*Config{
		"http & mux":                  {Enabled: true, Port: 8081, Handler: mux.New},
		"gin":                         {Enabled: true, Port: 8082, Handler: gin.New},
		"beego":                       {Enabled: true, Port: 8083, Handler: beego.New},
		"echo":                        {Enabled: true, Port: 8084, Handler: echo.New},
		"martini & martini-render":    {Enabled: true, Port: 8085, Handler: martini.New},
		"fasthttp & fasthttp-routing": {Enabled: true, Port: 8086, Handler: fasthttp.New},
		"iris":                        {Enabled: true, Port: 8087, Handler: iris.New},
		"revel":                       {Enabled: false, Port: 8088, Handler: revel.New}, // unavailable
		"buffalo":                     {Enabled: true, Port: 8089, Handler: buffalo.New},
		"goji":                        {Enabled: true, Port: 8090, Handler: goji.New},
		"gocraft":                     {Enabled: true, Port: 8091, Handler: gocraft.New},
		"httprouter":                  {Enabled: true, Port: 8092, Handler: httprouter.New},
		//"web":                         {Enabled: false, Port: 8093, Handler: web.New}, // unavailable
	}
)

func (c *Config) Available() bool {
	if c.Enabled == false || c.Handler == nil {
		return false
	}
	return true
}

func main() {
	numRequests := 100
	numGoRoutines := 5

	// start servers
	var err error
	for name, conf := range servers {
		if !conf.Available() {
			continue
		}

		log.Printf("starting %s server", name)
		server := conf.Handler(conf.Port)
		go server.Start()
	}
	// create output file
	log.Printf("create output file")
	file, err := createFile("./generated/", time.Now().Format(time.RFC3339), "txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(fmt.Sprintf("Number of Go Routines: %d\nNumber of Requests: %d\n\n", numGoRoutines, numRequests)); err != nil {
		panic(err)
	}
	_ = file.Sync()

	// run test
	<-time.After(time.Second * 5)
	for name, conf := range servers {
		if !conf.Available() {
			continue
		}

		log.Printf("running tests on %s", name)
		if _, err = file.WriteString(fmt.Sprintf(":: %s\n", name)); err != nil {
			panic(err)
		}

		elapsedTime := call(name, conf.Port, numGoRoutines, numRequests)
		if _, err = file.WriteString(fmt.Sprintf("Elapsed time: %f seconds\n\n", elapsedTime.Seconds())); err != nil {
			panic(err)
		}

		_ = file.Sync()
		<-time.After(time.Second * 1)
	}
}

func call(name string, port, numGoRoutines, numRequests int) time.Duration {
	start := time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(numGoRoutines)

	for i := 0; i < numGoRoutines; i++ {
		f := func(name string, id int, wg *sync.WaitGroup, numRequests int) {
			defer wg.Done()

			for i := 0; i < numRequests; i++ {
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
