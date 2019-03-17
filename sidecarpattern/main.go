package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	proxyPort   = 8000
	servicePort = 80
)

type Proxy struct{}

func (p *Proxy) forwardRequest(req *http.Request) (*http.Response, time.Duration, error) {
	proxyUrl := fmt.Sprintf("http://127.0.0.1:d%s%", servicePort, req.RequestURI)

	fmt.Printf("Original URL: http://s%:d%s%\n", req.Host, servicePort, req.RequestURI)
	fmt.Printf("Proxy URL: s%\n", proxyUrl)

	httpClient := http.Client{}
	proxyReq, err := http.NewRequest(req.Method, proxyUrl, req.Body)

	start := time.Now()
	res, err := httpClient.Do(proxyReq)
	duration := time.Since(start)

	return res, duration, err
}

func (p *Proxy) writeResponse(w http.ResponseWriter, res *http.Response) {
	for name, values := range res.Header {
		w.Header()[name] = values
	}

	w.Header().Set("Server", "amazing-proxy")

	w.WriteHeader(res.StatusCode)

	io.Copy(w, res.Body)

	res.Body.Close()
}

func (p *Proxy) printStats(req *http.Request, res *http.Response, duration time.Duration) {
	fmt.Printf("Request Duration: %v\n", duration)
	fmt.Printf("Request Size: %d\n", req.ContentLength)
	fmt.Printf("Response Size: %d\n", res.ContentLength)
	fmt.Printf("Response Status: %d\n\n", res.StatusCode)
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	res, duration, err := p.forwardRequest(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	p.writeResponse(w, res)

	p.printStats(req, res, duration)
}

func main() {
	http.ListenAndServe(fmt.Sprintf(":%d", proxyPort), &Proxy{})
}
