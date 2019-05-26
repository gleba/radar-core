package gates

import (
	"fmt"
	"github.com/gleba/radar-core/ux"
	"github.com/jmoiron/sqlx"
	"github.com/kshvakov/clickhouse"
	"log"
	"os"
)

var SqlX *sqlx.DB

func OpenClickHose() {
	//caCert, err := ioutil.ReadFile(
	//	"/usr/local/share/ca-certificates/Yandex/YandexInternalRootCA.crt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(caCert)
	//&tls.Config{
	//	RootCAs: caCertPool,
	//}

	var err error
	//tls?.RequestClientCert
	SqlX, err = sqlx.Open("clickhouse", os.Getenv("CLICKHOUSE"))

	ux.Safe(err)
	if err := SqlX.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}

	for _, table := range tablies {
		_, err = SqlX.Exec(table)
		ux.Safe(err)
	}
	log.Println("open gate: ClickHouse")
}
