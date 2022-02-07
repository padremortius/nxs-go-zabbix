# nxs-go-zabbix

This Go package provides access to Zabbix API v5.4.
Also see older versions in other branches.

At the time not all Zabbix API methods are implemented, but work in progress.

## Install

```
go get github.com/nixys/nxs-go-zabbix/v5
```

## Example of usage

*You may find more examples in unit-tests in this repository*

**Get all hosts on Zabbix server:**

```go
package main

import (
	"fmt"
	"os"

	"github.com/padremortius/nxs-go-zabbix/v5"
)

func main() {

	var z zabbix.Context

	/* Get variables from environment to login to Zabbix server */
	zbxHost := os.Getenv("ZABBIX_HOST")
	zbxAPIKey := os.Getenv("ZABBIX_APIKEY")

	if zbxHost == "" || zbxAPIKey == "" {
		fmt.Println("Login error: make sure environment variables `ZABBIX_HOST`, `ZABBIX_APIKEY` are defined")
		os.Exit(1)
	}

	/* Set connect params to zabbix server */
	z.SessionKey = zbxAPIKey
	z.Host = zbxHost

	/* Get all hosts */
	hObjects, _, err := z.HostGet(zabbix.HostGetParams{
		GetParameters: zabbix.GetParameters{
			Output: zabbix.SelectExtendedOutput,
		},
	})
	if err != nil {
		fmt.Println("Hosts get error:", err)
		return
	}

	/* Print names of retrieved hosts */
	fmt.Println("Hosts list:")
	for _, h := range hObjects {
		fmt.Println("-", h.Host)
	}
}
```

Run:

```
ZABBIX_HOST="https://zabbix.yourdomain.com/api_jsonrpc.php" ZABBIX_USERNAME="Admin" ZABBIX_PASSWORD="PASSWORD" go run main.go
```
