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
	Name    string
	Enabled bool
	Port    int
	Handler func(port int) server.IServer
}

var (
	servers = []*Config{
		{Enabled: true, Name: "beego", Port: 8081, Handler: beego.New},
		{Enabled: true, Name: "buffalo", Port: 8082, Handler: buffalo.New},
		{Enabled: true, Name: "echo", Port: 8083, Handler: echo.New},
		{Enabled: true, Name: "fasthttp & fasthttp-routing", Port: 8084, Handler: fasthttp.New},
		{Enabled: true, Name: "gin", Port: 8085, Handler: gin.New},
		{Enabled: true, Name: "gocraft", Port: 8086, Handler: gocraft.New},
		{Enabled: true, Name: "goji", Port: 8087, Handler: goji.New},
		{Enabled: true, Name: "http & mux", Port: 8088, Handler: mux.New},
		{Enabled: true, Name: "http-router", Port: 8089, Handler: httprouter.New},
		{Enabled: true, Name: "iris", Port: 8090, Handler: iris.New},
		{Enabled: true, Name: "martini & martini-render", Port: 8091, Handler: martini.New},
		{Enabled: false, Name: "revel", Port: 8092, Handler: revel.New}, // unavailable
		//{Enabled: false, Name: "web", Port: 8093, Handler: web.New}, // unavailable
	}
)

func (c *Config) Available() bool {
	if c.Enabled == false || c.Handler == nil {
		return false
	}
	return true
}

func main() {
	numRequests := 2000
	numGoRoutines := 25

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

	for _, conf := range servers {
		if !conf.Available() {
			continue
		}

		// start web server
		log.Printf(":: %s ::", conf.Name)
		log.Print("starting server")
		server := conf.Handler(conf.Port)
		go server.Start()
		<-time.After(time.Second * 1)

		// run test
		log.Print("running test")
		if _, err = file.WriteString(fmt.Sprintf(":: %s\n", conf.Name)); err != nil {
			panic(err)
		}

		elapsedTime := call(conf.Name, conf.Port, numGoRoutines, numRequests)
		if _, err = file.WriteString(fmt.Sprintf("Elapsed time: %f seconds\n\n", elapsedTime.Seconds())); err != nil {
			panic(err)
		}

		_ = file.Sync()

		log.Print("stopping server")
		if err = server.Stop(); err != nil {
			log.Print("error stopping server")
		}
	}

	log.Print("finished!")
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
