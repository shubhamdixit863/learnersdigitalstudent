package services

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == deletes {
		getId := strings.Split(r.URL.RawQuery, equalString)
		if getId[0] == idString {
			id, err := strconv.Atoi(getId[1])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			for i, user := range users {
				if user.ID == id {
					users = append(users[:i], users[i+1:]...)
					w.Write([]byte(fmt.Sprintln(userDeleted, user.Name, user.ID)))
					return
				}
			}
			w.Write([]byte(userNotFound))
			return
		}

	}
	w.WriteHeader(methodNotAllowed)
}
