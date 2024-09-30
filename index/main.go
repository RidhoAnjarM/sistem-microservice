package main

import (
	"encoding/json"
	"fmt"
	"sistem-microservice/author/constants"
	"sistem-microservice/author/controllers"
	"net"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	// Membaca konfigurasi server
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err.Error())
	}

	// Membuat instance baru controllers
	controller := controllers.NewInstance()

	// Start V1
	http.HandleFunc("/api/v1/hello", corsMiddleware(controller.Example))
	http.HandleFunc("/api/v1/users", corsMiddleware(controller.GetAllUsers))

	// Set address dan port tempat server berjalan
	ln, err := net.Listen("tcp", "localhost:"+constants.PORT)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening On localhost:" + constants.PORT)

	// Menjalankan http server
	if err = http.Serve(ln, nil); err != nil {
		panic(err)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, S-App-Authorization, D-App-Authorization, X-App-Authorization, lmsId, AppToken")
		w.Header().Set("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")
		w.Header().Set("Content-Type", "Application/Json")

		if r.Method == "OPTIONS" {
			_ = json.NewEncoder(w).Encode(make(map[string]interface{}))
			return
		}

		defer func(w http.ResponseWriter) {
			if e := recover(); e != nil {
				fmt.Println(time.Now().Format("2006/01/02 15:04:05"), e)
				fmt.Println(string(debug.Stack()))

				_ = json.NewEncoder(w).Encode(map[string]interface{}{
					"status":  false,
					"message": "Terjadi kesalahan dalam memproses permintaan anda", // Maaf, terjadi kesalahan saat memproses data
					"data":    nil,
				})
				return
			}
		}(w)

		next.ServeHTTP(w, r)
	}
}
