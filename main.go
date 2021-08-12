package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
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
		Profiling map[ProfilingName]*bytes.Buffer
	}

	ProfilingName string

	ShowUiCmdFunc func(fileName string) error
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

	ConstProfilingNameGoRoutine    ProfilingName = "goroutine"
	ConstProfilingNameThreadCreate ProfilingName = "threadcreate"
	ConstProfilingNameTrace        ProfilingName = "trace"
	ConstProfilingNameHeap         ProfilingName = "heap"
	ConstProfilingNameAllocs       ProfilingName = "allocs"
	ConstProfilingNameCPU          ProfilingName = "cpu"
	ConstProfilingNameMemory       ProfilingName = "memory"
	ConstProfilingNameGB           ProfilingName = "gb"
	ConstProfilingNameBlock        ProfilingName = "block"
	ConstProfilingNameMutex        ProfilingName = "mutex"

	goToolCmdProf  = "pprof"
	goToolCmdTrace = "trace"
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
		{Enabled: false, Name: "test 1", NumGoRoutines: 1, NumRequests: 100, Servers: allServers},
		{Enabled: false, Name: "test 2", NumGoRoutines: 1, NumRequests: 200, Servers: allServers},
		{Enabled: false, Name: "test 3", NumGoRoutines: 10, NumRequests: 100, Servers: allServers},
		{Enabled: false, Name: "test 4", NumGoRoutines: 10, NumRequests: 200, Servers: allServers},
		{Enabled: false, Name: "test 5", NumGoRoutines: 20, NumRequests: 100, Servers: allServers},
		{Enabled: false, Name: "test 6", NumGoRoutines: 20, NumRequests: 200, Servers: allServers},

		{Enabled: true, Name: "test all", NumGoRoutines: 10, NumRequests: 1, Servers: []ServerName{ConstServerNameBeego}},
	}

	testProfiling = []ProfilingName{
		ConstProfilingNameHeap,
		ConstProfilingNameMemory,
		ConstProfilingNameCPU,
		ConstProfilingNameTrace,
	}

	showProfilingCmd = map[ProfilingName]ShowUiCmdFunc{
		ConstProfilingNameTrace:  func(fileName string) error { return showGoToolUI(goToolCmdTrace, fileName) },
		ConstProfilingNameCPU:    func(fileName string) error { return showGoToolUI(goToolCmdProf, fileName) },
		ConstProfilingNameMemory: func(fileName string) error { return showGoToolUI(goToolCmdProf, fileName) },
	}

	printProfilingLines = map[ProfilingName][]string{
		ConstProfilingNameAllocs: []string{
			"# Alloc =",
			"# TotalAlloc =",
			"# Sys =",
			"# Mallocs =",
			"# Frees =",
			"# HeapAlloc =",
			"# HeapSys =",
			"# HeapIdle =",
			"# HeapInuse =",
			"# HeapReleased =",
			"# HeapObjects =",
			"# Stack =",
			"# MSpan =",
			"# MCache =",
			"# NumGC =",
			"# NumForcedGC =",
		},
		ConstProfilingNameHeap: []string{
			"# Alloc =",
			"# TotalAlloc =",
			"# Sys =",
			"# Mallocs =",
			"# Frees =",
			"# HeapAlloc =",
			"# HeapSys =",
			"# HeapIdle =",
			"# HeapInuse =",
			"# HeapReleased =",
			"# HeapObjects =",
			"# Stack =",
			"# MSpan =",
			"# MCache =",
			"# NumGC =",
			"# NumForcedGC =",
		},
	}
)

func main() {
	var err error
	var result map[ServerName]time.Duration
	now := time.Now()

	if result, err = tests.run(now); err != nil {
		panic(err)
	}

	if err = tests.createResultFile(result, now); err != nil {
		panic(err)
	}

	log.Printf("%sfinished%s", ConstColorRed, ConstColorReset)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	log.Println("to cancel: Ctrl + C")
	select {
	case <-termChan:
		log.Println("received term signal")
	}
}

func (tl TestList) run(now time.Time) (_ map[ServerName]time.Duration, err error) {
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

		if err = test.createResultFile(testResult, now); err != nil {
			return nil, err
		}

		if err = test.createProfileFiles(testResult, now); err != nil {
			return nil, err
		}

		// sum test duration for each server
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

		// clean garbage collector
		runtime.GC()
		<-time.After(time.Second * 5)

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
		Profiling: make(map[ProfilingName]*bytes.Buffer),
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
	for _, p := range testProfiling {
		buffer := &bytes.Buffer{}
		tr.Profiling[p] = buffer

		switch p {
		case ConstProfilingNameGoRoutine:
			if err := profiling.GoRoutine(buffer); err != nil {
				return err
			}
		case ConstProfilingNameThreadCreate:
			if err := profiling.ThreadCreate(buffer); err != nil {
				return err
			}
		case ConstProfilingNameHeap:
			if err := profiling.Heap(buffer); err != nil {
				return err
			}
		case ConstProfilingNameAllocs:
			if err := profiling.Allocs(buffer); err != nil {
				return err
			}
		case ConstProfilingNameCPU:
			if err := profiling.CPU(time.Second*5, buffer); err != nil {
				return err
			}
		case ConstProfilingNameMemory:
			if err := profiling.Memory(buffer); err != nil {
				return err
			}
		case ConstProfilingNameTrace:
			if err := profiling.Trace(time.Second*5, buffer); err != nil {
				return err
			}
		case ConstProfilingNameGB:
			if err := profiling.GC(buffer); err != nil {
				return err
			}
		case ConstProfilingNameBlock:
			if err := profiling.Block(100, buffer); err != nil {
				return err
			}
		case ConstProfilingNameMutex:
			if err := profiling.Mutex(100, buffer); err != nil {
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

func (t *Test) createResultFile(result map[ServerName]*TestResult, now time.Time) (err error) {
	var file *os.File

	name := fmt.Sprintf("./generated/%s - %s.txt", now.Format(time.RFC3339), t.Name)
	if file, err = createFile(name); err != nil {
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
	}

	return nil
}

func (t *Test) createProfileFiles(result map[ServerName]*TestResult, now time.Time) (err error) {

	for _, pName := range testProfiling {
		for serverName, tr := range result {
			if prof, ok := tr.Profiling[pName]; ok {
				buffer := &bytes.Buffer{}
				name := fmt.Sprintf("./generated/%s - %s - %s - profiling: %s.txt", now.Format(time.RFC3339), t.Name, serverName, pName)

				file, err := createFile(name)
				if err != nil {
					return err
				}
				defer file.Close()

				if prefixLines, ok := printProfilingLines[pName]; ok {
					reader := bufio.NewReader(prof)

				next:
					line, _, err := reader.ReadLine()
					for err == nil {

						for _, prefix := range prefixLines {
							if bytes.HasPrefix(line, []byte(prefix)) {
								buffer.Write(line)
								buffer.WriteString("\n")
								break
							}
						}
						goto next
					}
				} else {
					buffer = prof
				}

				if _, err = file.Write(buffer.Bytes()); err != nil {
					return err
				}

				_ = file.Sync()

				<-time.After(time.Second * 1)

				if f, ok := showProfilingCmd[pName]; ok {
					if err = f(name); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (tl TestList) createResultFile(result map[ServerName]time.Duration, now time.Time) (err error) {
	var file *os.File
	name := fmt.Sprintf("./generated/%s - result.txt", now.Format(time.RFC3339))
	if file, err = createFile(name); err != nil {
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

func createFile(name string) (file *os.File, err error) {
	file, err = os.Create(name)
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

func showGoToolUI(command, fileName string) error {
	cmd := exec.Command("go", "tool", command, "-http=:", fileName)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("cannot start pprof UI: %v", err)
	}

	return nil
}
