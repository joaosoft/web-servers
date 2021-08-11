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
|test 1|1|100|
|test 2|1|200|
|test 3|10|100|
|test 4|10|200|
|test 5|20|100|
|test 6|20|200|

### Result
|Server|Elapsed time (seconds)|Ranking|
|------|----------------------|-------|
|beego|21.197653|5|
|buffalo|21.348946|7|
|echo|20.856427|2|
|fasthttp & fasthttp-routing|20.549534|3|
|fiber|20.895227|4|
|gin|21.460733|10|
|gocraft|21.475442|9|
|goji|10.773424|1|
|http & mux|21.361887|8|
|http-router|21.306162|6|
|iris|21.869909|12|
|martini & martini-render|21.859482|11|
|revel|-|-|
|web|-|-|

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
