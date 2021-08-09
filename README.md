# web-servers
[![Build Status](https://travis-ci.com/joaosoft/web-servers.svg?branch=master)](https://travis-ci.com/joaosoft/web-servers) | [![codecov](https://codecov.io/gh/joaosoft/web-servers/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/web-servers) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/web-servers)](https://goreportcard.com/report/github.com/joaosoft/web-servers) | [![GoDoc](https://godoc.org/github.com/joaosoft/web-servers?status.svg)](https://godoc.org/github.com/joaosoft/web-servers)

Web servers example:
* [beego](https://github.com/beego/beego)
* [buffalo](https://github.com/gobuffalo/buffalo) 
* [echo](https://github.com/labstack/echo)
* [fasthttp](https://github.com/valyala/fasthttp) & [fasthttp-routing](https://github.com/qiangxue/fasthttp-routing)
* [gin](https://github.com/gin-gonic/gin)
* [gocraft](https://github.com/gocraft/web) 
* [goji](https://github.com/goji/goji) 
* [http](https://github.com/golang/go/blob/master/src/net/http) & [mux](https://github.com/gorilla/mux)
* [http-router](github.com/julienschmidt/httprouter) 
* [iris](https://github.com/kataras/iris) 
* [martini](https://github.com/go-martini/martini) & [martini-render](https://github.com/martini-contrib/render)
* [revel](https://github.com/revel/revel) 
* [web](github.com/joaosoft/web) 

## API
- `GET` to `/v1/persons/:id_person?age=30`
- `GET` to `/v1/persons/:id_person/addresses/:id_address`
- `GET` to `/v1/errors?id_error=200`

## Benchmark
>#### Tested with 10 go routines, each one doing 100 requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time (seconds)|Ranking|
|------|----|------------|-----------------|
|beego|8081|2.308835|11|
|buffalo|8082|2.229512|8|
|echo|8083|2.202042|5|
|fasthttp & fasthttp-routing|8084|2.178394|4|
|gin|8085|2.207624|6|
|gocraft|8086|2.177916|3|
|goji|80987|1.100797|2|
|http & mux|8088|2.216122|7|
|http-router|8089|1.095665|1|
|iris|8090|2.261207|10|
|martini & martini-render|8091|2.260218|9|
|revel|8092|-|-|
|web|8093|-|-|

>#### Tested with 50 go routines, each one doing 100 requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time (seconds)|Ranking|
|------|----|------------|-----------------|
|beego|8081|2.269001|10|
|buffalo|8082|2.208634|8|
|echo|8083|2.149836|3|
|fasthttp & fasthttp-routing|8084|2.159298|6|
|gin|8085|2.175926|5|
|gocraft|8086|2.166956|4|
|goji|80987|1.105030|1|
|http & mux|8088|2.185361|7|
|http-router|8089|1.177891|2|
|iris|8090|2.226463|9|
|martini & martini-render|8091|2.275814|11|
|revel|8092|-|-|
|web|8093|-|-|

>#### Tested with 10 go routines, each one doing 1000 requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time (seconds)|Ranking|
|------|----|------------|-----------------|
|beego|8081|22.180687|4|
|buffalo|8082|22.492784|6|
|echo|8083|22.552805|7|
|fasthttp & fasthttp-routing|8084|22.745346|8|
|gin|8085|21.815680|3|
|gocraft|8086|22.259329|5|
|goji|80987|11.592532|1|
|http & mux|8088|23.594827|11|
|http-router|8089|11.647801|2|
|iris|8090|23.368297|10|
|martini & martini-render|8091|22.966796|9|
|revel|8092|-|-|
|web|8093|-|-|

>#### Tested with 25 go routines, each one doing 2000 requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time (seconds)|Ranking|
|------|----|------------|-----------------|
|beego|8081|58.257302|6|
|buffalo|8082|56.215957|3|
|echo|8083|59.959902|7|
|fasthttp & fasthttp-routing|8084|60.870985|8|
|gin|8085|57.990198|5|
|gocraft|8086|61.508768|9|
|goji|80987|40.771783|1|
|http & mux|8088|63.721881|10|
|http-router|8089|46.644245|2|
|iris|8090|77.142292|11|
|martini & martini-render|8091|57.308490|4|
|revel|8092|-|-|
|web|8093|-|-|

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
