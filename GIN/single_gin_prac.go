package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type response struct {
	Sum    int  `json:"summation"`
	Status bool `json:"stat"`
}
type request struct {
	A int `form:"num1"`
	B int `form:"num2"`
}

/*
type response struct{
    sum int `json:"sum"`
    status bool `json:"status"`
}*/
func add(c *gin.Context) {

	var res response
	var req request
	var salary int
	err := c.BindQuery(&req)
	if err != nil {
		fmt.Println(err)
	}
	db, err_db := sql.Open("mysql", "root:1234@123@tcp(127.0.0.1:3306)/dotDB")
	if err_db != nil {
		fmt.Println("error accessing database")
		c.JSON(
			http.StatusOK, gin.H{"message": "error for db"})

	}
	defer db.Close()
	rows, err_row := db.Query("SELECT salary FROM dotDB.new_users WHERE id=?", 1)
	if err_row != nil {
		fmt.Println(err)
		c.JSON(
			http.StatusOK, gin.H{"salary": "error for rows"})
	}
	defer rows.Close()
	rows.Next()
	err_col := rows.Scan(&salary)
	if err_col != nil {
		fmt.Println("error while accessing row values")
		c.JSON(http.StatusOK, gin.H{"salary": salary, "error message": err_col})

	}

	res.Sum = (req.A + req.B) * salary
	//res.Sum = (req.A + req.B)

	c.JSON(200, res)

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
