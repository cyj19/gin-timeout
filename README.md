# gin-timeout

![gin-timeout](https://img.shields.io/github/license/cyj19/gin-timeout)  
a simple and practical gin timeout Middleware



## Installation

1. install package
```
go get -u github.com/cyj19/gin-timeout
```
2. import package
```
import timeout "github.com/cyj19/gin-timeout"
```

## Example

```
func main() {
    r := gin.Defalut()
    timeout := 3 * time.Second
    msg := `{"msg": "handle timeout"}`
    opt := timeout.Option{
        Timeout: &timeout,
        Code: 500,
        Msg: msg
    }
    r.Use(timeout.ContextTimeout(opt))
    r.GET("/ping", func(c *gin.Context) {
		time.Sleep(4 * time.Second)
		c.String(http.StatusOK, "pong")
	})
    r.Run()
}
```
