package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	// we assume that the server is up and running
	generateFile := false

	elapsedTime := call(8081, 100)
	content := fmt.Sprintf("Elapsed time: %f", elapsedTime.Seconds())

	if generateFile {
		if err := createFile(".", "generated", "text", []byte(content)); err != nil {
			panic(err)
		}
	} else {
		fmt.Println(content)
	}
}

func call(port, numGoRoutines int) time.Duration {
	wg := &sync.WaitGroup{}

	start := time.Now()
	for i := 0; i <= numGoRoutines; i++ {
		go func(id int, wg *sync.WaitGroup) {
			wg.Add(1)
			defer wg.Done()

			url := fmt.Sprintf("http://localhost:%d/v1/persons/%d/addresses/%d", port, id, id)
			_, err := http.Get(url)
			if err != nil {
				panic(err)
			}
		}(i, wg)
	}

	wg.Wait()

	return time.Since(start)
}

func createFile(folder, name, extension string, content []byte) (err error) {
	fileName := fmt.Sprintf("%s/%s.%s", folder, name, extension)

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err = file.Write(content); err != nil {
		return err
	}

	return nil
}