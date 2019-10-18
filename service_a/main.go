package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	_, _ = fmt.Fprintf(w, "Calling Service B\n")

	// Todo: Create a request to service B
	req, err := http.NewRequest("GET", "<ip to contact service b>", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}

// For tracing:
//	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
//	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
//	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
//	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
//	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
//	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
//	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	_, _ = fmt.Fprintf(w, string(body))
	_, _ = fmt.Fprintf(w, "Hello from service A\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
