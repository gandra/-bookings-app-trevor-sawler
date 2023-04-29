package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Setup test ...")

	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(http.ResponseWriter, *http.Request) {

}
