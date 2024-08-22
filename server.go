package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func getData(c *gin.Context) {
	resp, err := http.Get("https://www.coach.com/api/get-shop/women/handbags/view-all")

	if err != nil {
		fmt.Println("Error:\t", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	c.JSON(200, data)
}

func main() {
	router := gin.Default()
	router.GET("/priceCheck", getData)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7200"
	}

	router.Run("0.0.0.0:" + port)
}
