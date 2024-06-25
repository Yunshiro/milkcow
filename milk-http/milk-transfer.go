package milk_http

import (
	"fmt"
	"net/http"
)

// MilkTransfer trans the data to
func MilkTransfer(d interface{}) MilkFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fp, err := fmt.Fprint(w, d)
		if err != nil {
			fmt.Printf("Fprint(%v) Err: %v\n", fp, err)
		}
		fmt.Printf("%v path: %v\n", r.Method, r.URL.Path)
	}
}
