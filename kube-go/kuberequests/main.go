package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const apiServerEndpoint string = "https://localhost:41707"
const caPath string = "/home/unthinkable/kube-ca.crt"
const tokenPath string = "/home/unthinkable/kube-token"

func main() {
	fmt.Println("Connecting to Kubernetes...!!")
	fmt.Println("")
	fmt.Println("Enter the endpoint path : ")

	var path string
	buffer := bufio.NewReader(os.Stdin)

	path, err := buffer.ReadString('\n')

	errNotNil(err)

	GetReuest(path)

}

func GetReuest(path string) {
	const url string = apiServerEndpoint
	completeURL := url + path

	fmt.Println("Requesting GET on : ", completeURL)

	client := httpsClient(caPath)

	req := addHeaders(completeURL)

	res, err := client.Do(req)

	errNotNil(err)

	defer res.Body.Close()

	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content Length: ", res.ContentLength)

	jsondata, err := ioutil.ReadAll(res.Body)
	errNotNil(err)

	fmt.Println("Respose is : ", string(jsondata))

}

func addHeaders(completeURL string) *http.Request {
	req, _ := http.NewRequest("GET", strings.TrimSpace(completeURL), nil)
	token, err := ioutil.ReadFile(tokenPath)
	errNotNil(err)
	authValue := "Bearer " + string(token)
	fmt.Println("Token is ", authValue)
	req.Header.Set("Authorization", strings.TrimSpace(authValue))
	return req
}

func httpsClient(caPath string) *http.Client {
	caCert, err := ioutil.ReadFile(caPath)
	errNotNil(err)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}
	return client
}

func errNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
