package test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	beetest "github.com/astaxie/beego/testing"
)

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

//http://127.0.0.1:8080/shorten?longurl=http://www.beego.me/
//该方法不可行，会报下面的错误No connection could be made because the target machine actively refused it.
func TestShort(t *testing.T) {
	request := beetest.Post("/shorten")
	request.Param("longurl", "http://www.beego.me/")
	response, err := request.Response()
	if err != nil {
		t.Fatal(err)
	}
	//	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	var s ShortResult
	json.Unmarshal(contents, &s)
	if s.UrlShort == "" {
		t.Fatal("shorturl is empty")
	} else {
		t.Log(s.UrlShort)
	}

}

//http://127.0.0.1:8080/expend?shorturl=05ftgI
//该方法不可行，会报下面的错误No connection could be made because the target machine actively refused it.
func TestExpend(t *testing.T) {
	request := beetest.Get("/expand")
	request.Param("shorturl", "5laZF")
	response, err := request.Response()
	if err != nil {
		t.Fatal(err)
	}
	//	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	var s ShortResult
	json.Unmarshal(contents, &s)
	if s.UrlLong == "" {
		t.Fatal("urllong is empty")
	} else {
		t.Log(s.UrlLong)
	}
}
