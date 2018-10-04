package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "Hello Readers of golangcode.")
	// router := gin.Default()
	// srv := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }
	// router.GET("/welcome", func(c *gin.Context) {
	// 	response, err := http.Get("http://118.163.223.157:7090/")
	// 	if err != nil {
	// 	}
	// 	defer response.Body.Close()
	// 	body, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println(string(body))
	// })

	// srv.ListenAndServe()

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := srv.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown: ", err)
	// }
	// log.Println("Server exiting")
}
