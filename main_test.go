package main
import (
	"testing"
	"net/http"
	"io/ioutil"
)

func TestHandler(t *testing.T){
	testText := "Bienvenido, testing"
	http.HandleFunc("/", handler)
	go http.ListenAndServe(":80", nil)
	resp, err := http.Get("http://127.0.0.1:80/?key=testing")
	if err != nil{
		t.Error()
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		t.Error()
	}
	if string(responseData) != testText{
		t.Error()
	}
}