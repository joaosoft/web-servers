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
Tested with 100 requests at the same time 
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time|
|------|----|------------|
|http & mux|8081|0.016552|
|gin|8082|0.012126|
|beego|8083|0.013656|
|echo|8084|0.010677|
|martini & martini-render|8085|0.016708|
|fasthttp & fasthttp-routing|8086|0.011158|
|iris|8087|0.017330|
|revel|8088|0.115028|
|buffalo|8089|0.020912|
|goji|8090|0.010759|
|gocraft|8091|0.014490|
|web|8091|0|

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
