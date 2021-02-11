OMEx open api v3 go sdk usage :
-----

### 1.Downloads or updates OMEX code's dependencies, in your command line:

```
go get -u github.com/omexapp/omexchain-go-sdk-api
```
### 2.Write the go file. warm tips: test go file, must suffix *_test.go, eg: omex_open_api_v3_test.go
```
package gotest

import (
	"fmt"
	"github.com/omexapp/omexchain-go-sdk-api"
	"testing"
)

func TestOMExServerTime(t *testing.T) {
	serverTime, err := NewOMExClient().GetServerTime()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("OMEx's server time: ", serverTime)
}

func NewOMExClient() *omex.Client {
	var config omex.Config
	config.Endpoint = "https://www.omex.app/"
	config.ApiKey = ""
	config.SecretKey = ""
	config.Passphrase = ""
	config.TimeoutSecond = 45
	config.IsPrint = true
	config.I18n = omex.ENGLISH

	client := omex.NewClient(config)
	return client
}
```
### 3. run test go:
```
go test -v -run TestOMExServerTime omex_open_api_v3_test.go
```
