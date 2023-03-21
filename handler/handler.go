package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Excute(c *gin.Context) {
	resp := make(map[string]interface{})
	resp["message"] = "suc"
	database := c.Query("database")
	excute := c.Query("excute")
	db, err := sql.Open("mysql", database)
	if err != nil {
		fmt.Println("Error executing connect:", err)
		resp["message"] = err
		return
	}
	res, err := db.Exec(excute)

	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		c.Status(http.StatusNotFound)
	} else {
		resp["resulte"] = res
	}
	c.JSON(http.StatusOK, resp)
	defer db.Close()
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
