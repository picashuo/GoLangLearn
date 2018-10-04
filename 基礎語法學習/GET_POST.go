package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// func main() {
// 	client := &http.Client{}

// 	req, err := http.NewRequest("POST", "https://tw.yahoo.com/", strings.NewReader("name=test"))
// 	if err != nil {
// 		// handle error
// 	}

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	req.Header.Set("Cookie", "name=test")

// 	resp, err := client.Do(req)

// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		// handle error
// 	}

// 	fmt.Println(string(body))
// }
