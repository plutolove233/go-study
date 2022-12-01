package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("fetch error, ", err)
			return
		}

		h, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("fetch reading ", url, ": ", err)
			return
		}

		fmt.Println(string(h))
	}
}
