package main

import (
	"log"
	"os"
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	f, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		client := &http.Client{}
		url := fmt.Sprintf(scanner.Text())
		req, _ := http.NewRequest("GET", url, nil)
		req.Proto = "HTTP/1.0"
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Printf("\n%s", url)
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
}
