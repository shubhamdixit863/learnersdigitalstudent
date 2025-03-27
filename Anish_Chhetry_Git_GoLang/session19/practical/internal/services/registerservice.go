package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == post {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(bodyNotSupported))
		}
		var re RequestBody
		err = json.Unmarshal(body, &re)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		users = append(users, re)
		//log.Println("Registered: ", re.Name, re.ID)
		w.Write([]byte(fmt.Sprintln(userRegistered, re.Name, re.ID)))
		return
	}
	w.WriteHeader(methodNotAllowed)
}
