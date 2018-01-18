package nap

import (
	"testing"
	"net/http"
)
func TestAPICall(t *testing.T) {
	api := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		return nil
	})
	res := NewResource("/get", "GET", router)
	api.AddResource("get", res)
	if err := api.Call("get", nil); err != nil {
		t.Fail()
	}
}