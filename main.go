package main

import (
	"os"
	"rezza-mock/routers"

	"github.com/joho/godotenv"
)

// NGUYEN VAN A
// 9704 0000 0000 0018
// 03/07
// OTP

func init() {
	godotenv.Load(".env")
}

func main() {
	// fmt.Println(paymentservice.GetMomoPayURL("pay with MoMo", "http://localhost:8081", "http://localhost:8081", "10000", ""))
	router := routers.InitRouter()

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	router.Run(":" + port)
}
