/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
*/

package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	liveapi := "https://newtrackon.com/api/live"
	stableapi := "https://newtrackon.com/api/stable"
	allapi := "https://newtrackon.com/api/all"
	resp, _ := http.Get(liveapi)
	strveri, _ := ioutil.ReadAll(resp.Body)
	go ioutil.WriteFile("Elitlive.txt", strveri, 0666)
	resp, _ = http.Get(stableapi)
	strveri, _ = ioutil.ReadAll(resp.Body)
	go ioutil.WriteFile("Elitstable.txt", strveri, 0666)
	resp, _ = http.Get(allapi)
	strveri, _ = ioutil.ReadAll(resp.Body)

	ioutil.WriteFile("Elitall.txt", strveri, 0666)

}

/*
	stableapi:= "https://newtrackon.com/api/stable"
	allapi:="https://newtrackon.com/api/all"

*/
