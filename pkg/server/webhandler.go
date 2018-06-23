package klinserver

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func (f *conn) handleWebHook(w http.ResponseWriter, r *http.Request) {
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
		var m f.payload
		err = json.Unmarshal(f, &m)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m.C)
		//fmt.Printf("%#v", m)
		fmt.Println(string(f))
		w.WriteHeader(status)
		w.Write([]byte(msg))
		return
	}
}
