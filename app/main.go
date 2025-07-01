package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	env := os.Getenv("ENV")
	log.Printf("Current ENV: %s", env)
	if env == "" {
		env = "development"
		log.Printf("ENV not found, use default: development")

	}

	// 僅返回純淨的訊息，不帶顏色碼
	var rawMessage string
	var messageColor string // 可以新增一個欄位來表示建議的顏色

	switch env {
	case "production":
		rawMessage = "Hello from production!"
		messageColor = "red" // 或者 "danger"
		gin.SetMode(gin.ReleaseMode)
	case "staging":
		rawMessage = "Hello from staging!"
		messageColor = "yellow" // 或者 "warning"
		gin.SetMode(gin.ReleaseMode)
	case "development":
		rawMessage = "Hello from development!"
		messageColor = "green" // 或者 "success"
		gin.SetMode(gin.DebugMode)
	default:
		rawMessage = "Hello from unknown environment!"
		messageColor = "blue" // 或者 "info"
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     rawMessage,   // 傳遞原始訊息
			"color":       messageColor, // 傳遞一個表示顏色的提示
			"environment": env,          // 也可以傳遞環境名稱
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
