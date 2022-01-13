# Test with Gin (Golang)

This Gin project is used for basic framework review with performance analysis that can be used as reference to decide on the framework to use in microservices. To know more information, visit [Performance Result](https://github.com/samueltan3972/framework-review)

## RUN
Install go before running the project. [https://go.dev/](https://go.dev/)
```
go get .
go run .
```

## Docker 
```
docker build -t seanhao1233/test-with-gin .	
docker run -p 8080:8080 seanhao1233/test-with-gin
```

## View Result

View the result at http://localhost:8080
it comes with 3 API end point
- /hello : return a simple hello world message
- /database : perform all database CRUD operation
- /fibonacci : compute and return first 5,000 fibonacci numbers


## Help

* The [Gin official website](https://gin-gonic.com/).
* The [Gin documentation](https://gin-gonic.com/docs/).
* The [Gin github](https://github.com/gin-gonic/gin).
* The [Gin sample apps](https://github.com/gin-gonic/examples).
