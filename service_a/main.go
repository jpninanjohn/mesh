package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	invokeServiceB(w, r)
	invokeDobby(w, r)
	_, _ = fmt.Fprintf(w, "\nHello from service A\n")
}

func invokeServiceB(w http.ResponseWriter, r *http.Request) {
	// Call service B
	_, _ = fmt.Fprintf(w, "Calling Service B\n")
	req, err := http.NewRequest("GET", "http://service-b/", nil)
	if err != nil {
		fmt.Printf("%s", err)
		_, _ = fmt.Fprintf(w, "error creating request: %s\n", err.Error())
		return
	}

	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		_, _ = fmt.Fprintf(w, "error making request: %s\n", err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		_, _ = fmt.Fprintf(w, "error reading response: %s\n", err.Error())
		return
	}
	_, _ = fmt.Fprintf(w, string(body))
}

func invokeDobby(w http.ResponseWriter, r *http.Request) {
	// Call dobby service
	_, _ = fmt.Fprintf(w, "Calling Dobby Service\n")
	req, err := http.NewRequest("GET", "http://dobby/health", nil)
	if err != nil {
		fmt.Printf("%s", err)
		_, _ = fmt.Fprintf(w, "error creating request: %s\n", err.Error())
		return
	}

	req.Header.Add("x-request-id", r.Header.Get("x-request-id"))
	req.Header.Add("x-b3-traceid", r.Header.Get("x-b3-traceid"))
	req.Header.Add("x-b3-spanid", r.Header.Get("x-b3-spanid"))
	req.Header.Add("x-b3-parentspanid", r.Header.Get("x-b3-parentspanid"))
	req.Header.Add("x-b3-sampled", r.Header.Get("x-b3-sampled"))
	req.Header.Add("x-b3-flags", r.Header.Get("x-b3-flags"))
	req.Header.Add("x-ot-span-context", r.Header.Get("x-ot-span-context"))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		_, _ = fmt.Fprintf(w, "error making request: %s\n", err.Error())
		return
	}

	_, _ = fmt.Fprintf(w, "status code from dobby: %d\n", resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		_, _ = fmt.Fprintf(w, "error reading response: %s\n", err.Error())
		return
	}
	_, _ = fmt.Fprintf(w, "response from dobby: %s", string(body))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
