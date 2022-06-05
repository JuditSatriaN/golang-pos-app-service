package personalia

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	personaliactrl "github.com/golang_pos_app_service/controller/services/personalia"
)

func ServicesGetMembers(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	members, err := personaliactrl.GetAllMember(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	membersBytes, err := json.Marshal(members)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(membersBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
