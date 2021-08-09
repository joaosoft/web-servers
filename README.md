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
Tested with 2 go routines, each one doing 100 requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

|Server|Port|Elapsed time|
|------|----|------------|
|http & mux|8081|7.505765|
|gin|8082|7.448347|
|beego|8083|7.521267|
|echo|8084|7.542187|
|martini & martini-render|8085|7.556153|
|fasthttp & fasthttp-routing|8086|7.590783|
|iris|8087|7.443879|
|revel|8088|-|
|buffalo|8089|0.020912|
|goji|8090|4.191412|
|gocraft|8091|7.639343|
|web|8091|-|

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
