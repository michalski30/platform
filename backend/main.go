package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://localhost:5173",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // Erlaubt Cookies und Auth-Header (z.B. Bearer Tokens)
		MaxAge:           12 * time.Hour, // Wie lange der Browser die CORS-Antwort cachen darf
	}

	// Middleware global für alle Routen aktivieren
	router.Use(cors.New(config))

	router.GET("/call", func(c *gin.Context) {
		fmt.Println("asaaaaaaaaaa")
		c.JSON(200, gin.H{"status": "ppppppppppppppp!"})
	})

	router.Run(":8080")
}
