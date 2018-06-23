package klinserver

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func (f *conn) handleWebHook(w http.ResponseWriter, r *http.Request, p Dowork) {
	if strings.HasPrefix(r.Header.Get("content-type"), "multipart/form-data") {
		t, _, _ := r.FormFile("file")
		to, _ := os.Create("shit")
		io.Copy(to, t)
		to.Close()
		msg := "file transferred"
		status := 200
		fmt.Println(msg, status)
		w.WriteHeader(status)
		w.Write([]byte(msg))
		fmt.Println(r.FormValue("filename"))
		return
	} else {
		msg := "Got payload"
		status := 200
		fmt.Println(msg, status)
		f, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		//var m payload
		err = json.Unmarshal(f, &p)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%#v", m)
		fmt.Println("printing do it", p.Doit())
		fmt.Println(string(f))
		w.WriteHeader(status)
		w.Write([]byte(msg))
		return
	}
}
