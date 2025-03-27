package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == put {

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

		getId := strings.Split(r.URL.RawQuery, equalString)
		if getId[0] == idString {

			id, err := strconv.Atoi(getId[1])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			for i, user := range users {
				if user.ID == id {
					users[i].Name = re.Name
					w.Write([]byte(fmt.Sprintln(userUpdated, re.Name, user.ID)))
					return
				}
			}

			w.Write([]byte(userNotFound))
			return
		}
	}
	w.WriteHeader(methodNotAllowed)
}
