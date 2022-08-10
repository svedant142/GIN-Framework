// comment
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
type request_update struct {
	Salary int `form:"salary"`
}
type response_update struct {
	Status bool `json:"status"`
}

/*
type response struct{
    sum int `json:"sum"`
    status bool `json:"status"`
}*/
func initialize_database() *sql.DB {
	db, err_db := sql.Open("mysql", "root:1234@123@tcp(127.0.0.1:3306)/dotDB")
	if err_db != nil {
		fmt.Println("error accessing database")
		// c.JSON(
		// 	http.StatusOK, gin.H{"message": "error for db"})
	}
	return db
}
func add(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var res response
		var req request
		var salary int
		err := c.BindQuery(&req)
		if err != nil {
			fmt.Println(err)
		}
		// db, err_db := sql.Open("mysql", "root:1234@123@tcp(127.0.0.1:3306)/dotDB")
		// if err_db != nil {
		// 	fmt.Println("error accessing database")
		// 	c.JSON(
		// 		http.StatusOK, gin.H{"message": "error for db"})

		// }

		rows := db.QueryRow("SELECT salary FROM dotDB.new_users WHERE id=?", 1)
		// if err_row != nil {
		// 	fmt.Println(err)
		// 	c.JSON(
		// 		http.StatusOK, gin.H{"salary": "error for rows"})
		// }
		// defer rows.Close()
		// rows.Next()

		err_col := rows.Scan(&salary)
		if err_col != nil {
			fmt.Println("error while accessing row values")
			c.JSON(http.StatusOK, gin.H{"salary": salary, "error message": err_col})

		}

		res.Sum = (req.A + req.B) * salary
		res.Status = true
		//res.Sum = (req.A + req.B)

		c.JSON(200, res)
	}
	return gin.HandlerFunc(fn)
}
func update(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var req request_update
		var res response_update
		err := c.BindQuery(&req)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		}

		stmt, err := db.Prepare("UPDATE dotDB.new_users SET salary = ? where id = 1")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		}
		_, err1 := stmt.Exec(&req.Salary)
		if err1 != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		}
		res.Status = true
		c.JSON(200, res)
	}
	return gin.HandlerFunc(fn)
}

func main() {
	var db *sql.DB = initialize_database()
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		c.JSON(
			http.StatusOK, gin.H{"message": "Hello World"})
	})
	r.GET("/api/add", add(db))
	r.POST("/api/update", update(db))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	defer db.Close()
}
