package slacksh

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
)

// Handler defines request & response of server
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", fmt.Sprintf("%s-%s", strings.Title(filepath.Base(reflect.TypeOf(response{}).PkgPath())), version))

	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	text := strings.Replace(r.FormValue("text"), "\r", "", -1)
	if text == "" {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}

	log.Println("cmd:", text)

	exec, err := Run(text)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if exec == "" {
		http.Error(w, http.StatusText(http.StatusNoContent), http.StatusNoContent)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	text = fmt.Sprintf("```%s```", exec)
	json, _ := json.Marshal(response{Type: "in_channel", Text: text})

	fmt.Fprintln(w, string(json))
}
