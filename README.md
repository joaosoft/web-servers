# web-servers
[![Build Status](https://travis-ci.com/joaosoft/web-servers.svg?branch=master)](https://travis-ci.com/joaosoft/web-servers) | [![codecov](https://codecov.io/gh/joaosoft/web-servers/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/web-servers) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/web-servers)](https://goreportcard.com/report/github.com/joaosoft/web-servers) | [![GoDoc](https://godoc.org/github.com/joaosoft/web-servers?status.svg)](https://godoc.org/github.com/joaosoft/web-servers)

Web servers example:
* [beego](https://github.com/beego/beego)
* [buffalo](https://github.com/gobuffalo/buffalo) 
* [echo](https://github.com/labstack/echo)
* [fasthttp](https://github.com/valyala/fasthttp) & [fasthttp-routing](https://github.com/qiangxue/fasthttp-routing)
* [fiber](https://github.com/gofiber/fiber)
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
>#### Tested with x go routines, each one doing y requests.
(`GET` to `/v1/persons/:id_person/addresses/:id_address`)

### Tests
|Tests|Go Routines|Requests|
|------|----------------------|-------|
|test 1|1|200|
|test 2|1|400|
|test 3|10|200|
|test 4|10|400|
|test 5|20|200|
|test 6|20|400|

### Result
|Server|Elapsed time (seconds)|Ranking|
|------|----------------------|-------|
|beego|41.040017|8|
|buffalo|41.535441|10|
|echo|41.219579|9|
|fasthttp & fasthttp-routing|40.874413|7|
|fiber|40.568224|6|
|gin|39.849965|3|
|gocraft|39.942582|4|
|goji|20.456166|1|
|http & mux|41.573545|11|
|http-router|21.485230|2|
|iris|40.341649|5|
|martini & martini-render|42.398068|12|
|revel|-|-|
|web|-|-|

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
