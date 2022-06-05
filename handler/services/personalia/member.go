package personalia

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	personaliaentity "github.com/golang_pos_app_service/entity/personalia"
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

func ServicesInsertMember(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var member personaliaentity.Member

	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = personaliactrl.InsertMember(ctx, member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Data member berhasil disimpan")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServicesUpdateMember(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var member personaliaentity.Member

	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, found, err := personaliactrl.GetMemberByID(ctx, member.ID)
	if !found {
		_, err = fmt.Fprintf(w, "Member dengan nama : %s tidak ditemukan.", member.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = personaliactrl.UpdateMember(ctx, member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Data member berhasil diubah")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServicesDeleteMember(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var member personaliaentity.Member

	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, found, err := personaliactrl.GetMemberByID(ctx, member.ID)
	if !found {
		_, err = fmt.Fprintf(w, "Member dengan id : %s tidak ditemukan.", member.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = personaliactrl.DeleteMember(ctx, member.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = fmt.Fprintf(w, "Data Member berhasil dihapus")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServicesUpsertMember(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var member personaliaentity.Member

	err := json.NewDecoder(r.Body).Decode(&member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, found, err := personaliactrl.GetMemberByID(ctx, member.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !found {
		err = personaliactrl.InsertMember(ctx, member)
	} else {
		err = personaliactrl.UpdateMember(ctx, member)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	memberBytes, err := json.Marshal(member)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(memberBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
