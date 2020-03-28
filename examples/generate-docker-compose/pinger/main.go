package pinger

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Main function of hello-world
func Main(d *typcore.Descriptor) (err error) {
	var resp *http.Response
	if resp, err = http.Get("http://127.0.0.1:7379/ping"); err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return
}
