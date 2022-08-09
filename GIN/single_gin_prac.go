// comment
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Sum    int  `json:"summation"`
	Status bool `json:"stat"`
}
type request struct {
	Ab int `form:"num1"`
	Cd int `form:"num2"`
}

/*
type response struct{
    sum int `json:"sum"`
    status bool `json:"status"`
}*/
func add(c *gin.Context) {

	//a, err := strconv.Atoi(c.Query("num1"))
	//b, err1 := strconv.Atoi(c.Query("num2"))
	var res response
	var req request
	err := c.BindQuery(&req)
	fmt.Println(err)
	//sum=a+b,status=true,}
	res.Sum = req.Ab + req.Cd
	res.Status = true
	// //sum += b
	// if err != nil || err1 != nil {
	// 	res.Status = false
	// 	c.JSON(404, gin.H{
	// 		"error":  "occurred",
	// 		"status": res.Status,
	// 	})
	// } else {
	c.JSON(200, res)
	// c.JSON(200, gin.H{
	// 	"sum":    res.Sum,
	// 	"status": res.Status,
	// })
}

func main() {
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.JSON(
			http.StatusOK, gin.H{"message": "Hello World"})
	})
	r.GET("/api/add", add)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
