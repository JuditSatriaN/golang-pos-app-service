package personalia

import (
	"encoding/json"
	"net/http"

	personaliactrl "github.com/golang_pos_app_service/controller/services/personalia"
)

func ServicesGetUser(w http.ResponseWriter, r *http.Request) {

	users, err := personaliactrl.GetAllUserController()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	usersBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersBytes)
}
