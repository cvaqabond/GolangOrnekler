package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, _ := http.Get("https://ixakbulut.com")
	alis, _ := ioutil.ReadAll(resp.Body)
	re := regexp.MustCompile(`http([^"]*)`)
	veri := re.FindAll(alis, -1)
	for _, a := range veri {
		fmt.Println(string(a))
	}

}
