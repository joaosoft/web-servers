# web-servers
[![Build Status](https://travis-ci.org/joaosoft/web-servers.svg?branch=master)](https://travis-ci.org/joaosoft/web-servers) | [![codecov](https://codecov.io/gh/joaosoft/web-servers/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/web-servers) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/web-servers)](https://goreportcard.com/report/github.com/joaosoft/web-servers) | [![GoDoc](https://godoc.org/github.com/joaosoft/web-servers?status.svg)](https://godoc.org/github.com/joaosoft/web-servers)

Web servers example:
* [http](https://github.com/golang/go/blob/master/src/net/http) & [mux](https://github.com/gorilla/mux)
* [gin](https://github.com/gin-gonic/gin)
* [beego](https://github.com/beego/beego)
* [echo](https://github.com/labstack/echo)
* [martini](https://github.com/go-martini/martini) & [martini-render](https://github.com/martini-contrib/render)
* [fasthttp](https://github.com/valyala/fasthttp) & [fasthttp-routing](https://github.com/qiangxue/fasthttp-routing)
* [iris](https://github.com/kataras/iris) 
* [revel](https://github.com/revel/revel) 
* [buffalo](https://github.com/gobuffalo/buffalo) 
* [goji](https://github.com/goji/goji) 
* [gocraft](https://github.com/gocraft/web) 
* [httprouter](github.com/julienschmidt/httprouter) 
* [web](github.com/joaosoft/web) 

## API
- `GET` to `/v1/persons/:id_person?age=30`
- `GET` to `/v1/persons/:id_person/addresses/:id_address`
- `GET` to `/v1/errors?id_error=200`

## Benchmark
>#### Tested with 10 go routines, each one doing 100 requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time (seconds)|
|------|----|------------|
|beego|8081|2.308835|
|buffalo|8082|2.229512|
|echo|8083|2.202042|
|fasthttp & fasthttp-routing|8084|2.178394|
|gin|8085|2.207624|
|gocraft|8086|2.177916|
|goji|80987|1.100797|
|http & mux|8088|2.216122|
|httprouter|8089|1.095665|
|iris|8090|2.261207|
|martini & martini-render|8091|2.260218|
|revel|8092|-|
|web|8093|-|

>#### Tested with 50 go routines, each one doing 100 requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time (seconds)|
|------|----|------------|
|beego|8081|2.269001|
|buffalo|8082|2.208634|
|echo|8083|2.149836|
|fasthttp & fasthttp-routing|8084|2.159298|
|gin|8085|2.175926|
|gocraft|8086|2.166956|
|goji|80987|1.105030|
|http & mux|8088|2.185361|
|httprouter|8089|1.177891|
|iris|8090|2.226463|
|martini & martini-render|8091|2.275814|
|revel|8092|-|
|web|8093|-|

>#### Tested with 10 go routines, each one doing 1000 requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time (seconds)|
|------|----|------------|
|beego|8081|22.180687|
|buffalo|8082|22.492784|
|echo|8083|22.552805|
|fasthttp & fasthttp-routing|8084|22.745346|
|gin|8085|21.815680|
|gocraft|8086|22.259329|
|goji|80987|11.592532|
|http & mux|8088|23.594827|
|httprouter|8089|11.647801|
|iris|8090|23.368297|
|martini & martini-render|8091|22.966796|
|revel|8092|-|
|web|8093|-|

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
