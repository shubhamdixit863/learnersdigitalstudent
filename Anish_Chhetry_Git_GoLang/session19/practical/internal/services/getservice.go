package services

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == get {
		getId := strings.Split(r.URL.RawQuery, equalString)
		if getId[0] == idString {
			id, err := strconv.Atoi(getId[1])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			for _, user := range users {
				if user.ID == id {
					data, _ := json.Marshal(user)
					w.Write(data)
					return
				}
			}

			w.Write([]byte(userNotFound))
			return
		}
		data, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(data)
		return

	}
	w.WriteHeader(methodNotAllowed)
}
