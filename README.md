# Simple approach to test Client for 3rd party API in Go with Gin

Using [Gin](https://github.com/gin-gonic/gin) it is quite simple to build a stub server that emulates 3rd 
party server behaviour. 
```Go
func NewServerApiStub() *gin.Engine {
	r := gin.Default()
	r.POST("/articles", stubCreateArticle)
	r.GET("/articles/:id", stubGetArticle)
	return r
}
```

The stub server's lifecycle controlled by tests: 
```Go
func TestMain(m *testing.M) {
	go func() {
		err := NewServerApiStub().Run(":8123")
		if err != nil {
			panic(err)
		}
	}()

	os.Exit(m.Run())
}
```

Like in a case with [Mountebank](http://www.mbtest.org/) the client makes real HTTP requests to the stub 
server during the tests.
