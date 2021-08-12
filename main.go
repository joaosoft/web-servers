package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
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
	fiber "web-servers/fiber/server"
	gin "web-servers/gin/server"
	gocraft "web-servers/gocraft/server"
	goji "web-servers/goji/server"
	httprouter "web-servers/httprouter/server"
	iris "web-servers/iris/server"
	martini "web-servers/martini/server"
	mux "web-servers/mux/server"
	revel "web-servers/revel/app/server"
	web "web-servers/web/server"

	"github.com/joaosoft/profiling"
)

type (
	ServerName string

	TestList []*Test

	Test struct {
		Enabled       bool
		Name          string
		NumGoRoutines int
		NumRequests   int
		Servers       []ServerName
	}

	Config struct {
		Name    ServerName
		Enabled bool
		Port    int
		Handler func(port int) server.IServer
	}

	TestResult struct {
		Duration  time.Duration
		Profiling *bytes.Buffer
	}

	Profiling string
)

const (
	ConstServerNameBeego                      ServerName = "beego"
	ConstServerNameBuffalo                    ServerName = "buffalo"
	ConstServerNameEcho                       ServerName = "echo"
	ConstServerNameFastHttpAndFastHttpRouting ServerName = "fasthttp & fasthttp-routing"
	ConstServerNameFiber                      ServerName = "fiber"
	ConstServerNameGin                        ServerName = "gin"
	ConstServerNameGocraft                    ServerName = "gocraft"
	ConstServerNameGoji                       ServerName = "goji"
	ConstServerNameHttpMux                    ServerName = "http & mux"
	ConstServerNameHttpRouter                 ServerName = "http-router"
	ConstServerNameIris                       ServerName = "iris"
	ConstServerNameMartiniMartiniRender       ServerName = "martini & martini-render"
	ConstServerNameRevel                      ServerName = "revel"
	ConstServerNameWeb                        ServerName = "web"

	ConstColorGreen = "\033[32m"
	ConstColorRed   = "\033[31m"
	ConstColorReset = "\033[0m"

	ConstProfilingGoRoutine    Profiling = "goroutine"
	ConstProfilingThreadCreate Profiling = "threadcreate"
	ConstProfilingHeap         Profiling = "heap"
	ConstProfilingAllocs       Profiling = "allocs"
	ConstProfilingCPU          Profiling = "cpu"
	ConstProfilingMemory       Profiling = "memory"
	ConstProfilingGB           Profiling = "gb"
	ConstProfilingBlock        Profiling = "block"
	ConstProfilingMutex        Profiling = "mutex"
)

var (
	servers = map[ServerName]*Config{
		ConstServerNameBeego:                      {Enabled: true, Name: ConstServerNameBeego, Handler: beego.New},
		ConstServerNameBuffalo:                    {Enabled: true, Name: ConstServerNameBuffalo, Handler: buffalo.New},
		ConstServerNameEcho:                       {Enabled: true, Name: ConstServerNameEcho, Handler: echo.New},
		ConstServerNameFastHttpAndFastHttpRouting: {Enabled: true, Name: ConstServerNameFastHttpAndFastHttpRouting, Handler: fasthttp.New},
		ConstServerNameFiber:                      {Enabled: true, Name: ConstServerNameFiber, Handler: fiber.New},
		ConstServerNameGin:                        {Enabled: true, Name: ConstServerNameGin, Handler: gin.New},
		ConstServerNameGocraft:                    {Enabled: true, Name: ConstServerNameGocraft, Handler: gocraft.New},
		ConstServerNameGoji:                       {Enabled: true, Name: ConstServerNameGoji, Handler: goji.New},
		ConstServerNameHttpMux:                    {Enabled: true, Name: ConstServerNameHttpMux, Handler: mux.New},
		ConstServerNameHttpRouter:                 {Enabled: true, Name: ConstServerNameHttpRouter, Handler: httprouter.New},
		ConstServerNameIris:                       {Enabled: true, Name: ConstServerNameIris, Handler: iris.New},
		ConstServerNameMartiniMartiniRender:       {Enabled: true, Name: ConstServerNameMartiniMartiniRender, Handler: martini.New},
		ConstServerNameRevel:                      {Enabled: false, Name: ConstServerNameRevel, Handler: revel.New}, // unavailable
		ConstServerNameWeb:                        {Enabled: false, Name: ConstServerNameWeb, Handler: web.New},     // unavailable
	}

	allServers = []ServerName{
		ConstServerNameBeego,
		ConstServerNameBuffalo,
		ConstServerNameEcho,
		ConstServerNameFastHttpAndFastHttpRouting,
		ConstServerNameFiber,
		ConstServerNameGin,
		ConstServerNameGocraft,
		ConstServerNameGoji,
		ConstServerNameHttpMux,
		ConstServerNameHttpRouter,
		ConstServerNameIris,
		ConstServerNameMartiniMartiniRender,
		ConstServerNameRevel,
		ConstServerNameWeb,
	}

	tests = TestList{
		{Enabled: true, Name: "test 1", NumGoRoutines: 1, NumRequests: 100, Servers: allServers},
		{Enabled: true, Name: "test 2", NumGoRoutines: 1, NumRequests: 200, Servers: allServers},
		{Enabled: true, Name: "test 3", NumGoRoutines: 10, NumRequests: 100, Servers: allServers},
		{Enabled: true, Name: "test 4", NumGoRoutines: 10, NumRequests: 200, Servers: allServers},
		{Enabled: true, Name: "test 5", NumGoRoutines: 20, NumRequests: 100, Servers: allServers},
		{Enabled: true, Name: "test 6", NumGoRoutines: 20, NumRequests: 200, Servers: allServers},
	}

	testProfiling = []Profiling{
		ConstProfilingAllocs,
		ConstProfilingHeap,
	}
)

func main() {
	var err error
	var result map[ServerName]time.Duration

	if result, err = tests.run(); err != nil {
		panic(err)
	}

	if err = tests.createResultFile(result); err != nil {
		panic(err)
	}

	log.Printf("%sfinished%s", ConstColorRed, ConstColorReset)
}

func (tl TestList) run() (_ map[ServerName]time.Duration, err error) {
	result := make(map[ServerName]time.Duration)
	var testResult map[ServerName]*TestResult

	for _, test := range tl {
		if !test.Enabled {
			continue
		}

		// run test
		if testResult, err = test.run(); err != nil {
			return nil, err
		}

		if err = test.createResultFile(testResult); err != nil {
			return nil, err
		}

		// update result
		for name, tr := range testResult {
			if value, ok := result[name]; ok {
				tr.Duration += value
			}
			result[name] = tr.Duration
		}
	}

	return result, nil
}

func (t *Test) run() (_ map[ServerName]*TestResult, err error) {
	log.Printf("%s:: test: %s%s", ConstColorRed, t.Name, ConstColorReset)

	result := make(map[ServerName]*TestResult)
	for _, s := range t.Servers {
		conf, ok := servers[s]
		if !ok {
			return nil, errors.New(fmt.Sprintf("server '%s' not found", s))
		}

		if !conf.Available() {
			continue
		}

		port, err := conf.GetPort()
		if err != nil {
			return nil, err
		}

		// start web server
		log.Printf(":: %s ::", conf.Name)
		log.Print("starting")
		newServer := conf.Handler(port)
		go newServer.Start()
		<-time.After(time.Second * 1)

		// run test
		log.Print("testing")

		result[conf.Name] = t.call(conf.Name, port)

		log.Print("stopping")
		if err = newServer.Stop(); err != nil {
			log.Print("error stopping server")
		}
		<-time.After(time.Second * 1)
	}

	log.Printf("%sdone%s", ConstColorGreen, ConstColorReset)
	return result, nil
}

func (t *Test) call(name ServerName, port int) (tr *TestResult) {
	tr = &TestResult{
		Profiling: &bytes.Buffer{},
	}
	wg := &sync.WaitGroup{}
	wg.Add(t.NumGoRoutines)

	start := time.Now()

	for i := 1; i <= t.NumGoRoutines; i++ {
		go handler(name, port, i, wg, t.NumRequests)
	}

	if err := tr.runProfiling(); err != nil {
		log.Printf("\nprofiling (name: %s) error: %s", name, err.Error())
		return
	}

	wg.Wait()

	tr.Duration = time.Since(start)

	return tr
}

func (tr TestResult) runProfiling() error {
	tr.Profiling.WriteString("\n::Profiling")

	for _, p := range testProfiling {
		switch p {
		case ConstProfilingGoRoutine:
			tr.Profiling.WriteString("\n:: GoRoutines\n")
			if err := profiling.GoRoutine(tr.Profiling); err != nil {
				return err
			}
		case ConstProfilingThreadCreate:
			tr.Profiling.WriteString("\n:: ThreadCreate\n")
			if err := profiling.ThreadCreate(tr.Profiling); err != nil {
				return err
			}
		case ConstProfilingHeap:
			tr.Profiling.WriteString("\n:: Heap\n")
			if err := profiling.Heap(tr.Profiling); err != nil {
				return err
			}
		case ConstProfilingAllocs:
			tr.Profiling.WriteString("\n:: Allocs\n")
			if err := profiling.Allocs(tr.Profiling); err != nil {
				return err
			}
		case ConstProfilingCPU:
			tr.Profiling.WriteString("\n:: CPU\n")
			if err := profiling.CPU(time.Second*1, tr.Profiling); err != nil {
				return err
			}
		case ConstProfilingMemory:
			tr.Profiling.WriteString("\n:: Memory\n")
			if err := profiling.Memory(tr.Profiling); err != nil {
				return err
			}
		case ConstProfilingGB:
			tr.Profiling.WriteString("\n:: Garbage Collection\n")
			if err := profiling.GC(tr.Profiling); err != nil {
				return err
			}
		case ConstProfilingBlock:
			tr.Profiling.WriteString("\n:: Block\n")
			if err := profiling.Block(10, tr.Profiling); err != nil {
				return err
			}
		case ConstProfilingMutex:
			tr.Profiling.WriteString("\n\n:: Mutex\n")
			if err := profiling.Mutex(10, tr.Profiling); err != nil {
				return err
			}
		}
	}

	return nil
}

func handler(name ServerName, port, id int, wg *sync.WaitGroup, numRequests int) {
	defer wg.Done()

	for index := 1; index <= numRequests; index++ {
		url := fmt.Sprintf("http://localhost:%d/v1/persons/%d/addresses/%d", port, id, index)
		response, err := http.Get(url)
		if err != nil {
			log.Printf("\nERROR 1 (name: %s, request: %d | %d) error: %s", name, id, index, err.Error())
			return
		}

		if response != nil {
			defer response.Body.Close()
			bodyResponse, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("\nERROR 2 (name: %s, request: %d | %d) error: %s", name, id, index, err.Error())
				return
			}

			var address models.Address
			if err = json.Unmarshal(bodyResponse, &address); err != nil {
				log.Printf("\nERROR 3 (name: %s, request: %d | %d) error: %s", name, id, index, err.Error())
				log.Printf(string(bodyResponse))
				return
			}

			if address.Id != strconv.Itoa(index) {
				log.Printf("\nERROR 4 (name: %s, request: %d | %d) invalid address %s", name, id, index, address.Id)
				log.Printf(string(bodyResponse))
				return
			}
		}
	}
}

func (t *Test) createResultFile(result map[ServerName]*TestResult) (err error) {
	var file *os.File

	name := fmt.Sprintf("%s - %s", time.Now().Format(time.RFC3339), t.Name)
	if file, err = createFile("./generated/", name, "txt"); err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(
		fmt.Sprintf("Test: %s\nNumber of Go Routines: %d\nNumber of Requests: %d\n\n",
			t.Name, t.NumGoRoutines, t.NumRequests)); err != nil {
		return err
	}

	for serverName, tr := range result {
		if _, err = file.WriteString(fmt.Sprintf(":: %s\n", serverName)); err != nil {
			return err
		}
		if _, err = file.WriteString(fmt.Sprintf("Elapsed time: %f seconds\n\n", tr.Duration.Seconds())); err != nil {
			return err
		}

		if _, err = file.Write(tr.Profiling.Bytes()); err != nil {
			return err
		}
	}

	return nil
}

func (tl TestList) createResultFile(result map[ServerName]time.Duration) (err error) {
	var file *os.File
	name := fmt.Sprintf("%s - result", time.Now().Format(time.RFC3339))
	if file, err = createFile("./generated/", name, "txt"); err != nil {
		return err
	}
	defer file.Close()

	for serverName, duration := range result {
		if _, err = file.WriteString(fmt.Sprintf(":: %s\n", serverName)); err != nil {
			return err
		}
		if _, err = file.WriteString(fmt.Sprintf("Elapsed time: %f seconds\n\n", duration.Seconds())); err != nil {
			return err
		}
	}

	return nil
}

func createFile(folder, name, extension string) (file *os.File, err error) {
	fileName := fmt.Sprintf("%s/%s.%s", folder, name, extension)

	file, err = os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, err
}

func (c *Config) Available() bool {
	if c.Enabled == false || c.Handler == nil {
		return false
	}
	return true
}

func (c *Config) GetPort() (int, error) {
	if c.Port > 0 {
		return c.Port, nil
	}

	return getFreePort()
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
