package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	caCert, err := ioutil.ReadFile(
		"/usr/local/share/ca-certificates/Yandex/YandexInternalRootCA.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	req, _ := http.NewRequest("GET", "https://rc1b-tiokbxm5t53mpmbj.mdb.yandexcloud.net:8443/", nil)
	query := req.URL.Query()
	query.Add("database", "db1")
	query.Add("query", "SELECT now()")

	req.URL.RawQuery = query.Encode()

	req.Header.Add("X-ClickHouse-User", "user1")
	req.Header.Add("X-ClickHouse-Key", "iddqdidkfa")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}
