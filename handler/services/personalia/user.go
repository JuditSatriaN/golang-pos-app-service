package personalia

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang_pos_app_service/pkg"

	personaliactrl "github.com/golang_pos_app_service/controller/services/personalia"

	personaliaentity "github.com/golang_pos_app_service/entity/personalia"
)

func ServicesGetUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := personaliactrl.GetAllUser(ctx)
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

func ServicesInsertUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user personaliaentity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Password != "" {
		user.Password, _ = pkg.HashPassword(user.Password)
	}

	err = personaliactrl.InsertUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, "Data user berhasil disimpan")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServicesUpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user personaliaentity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userDB, found, err := personaliactrl.GetUserByUserID(ctx, user.UserID)
	if !found {
		_, err = fmt.Fprintf(w, "User dengan nama : %s tidak ditemukan.", user.UserName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// replace to existing data
	if user.Password == "" {
		user.Password = userDB.Password
	} else {
		user.Password, _ = pkg.HashPassword(user.Password)
	}

	err = personaliactrl.UpdateUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, "Data user berhasil diubah")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServicesDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user personaliaentity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, found, err := personaliactrl.GetUserByUserID(ctx, user.UserID)
	if !found {
		_, err = fmt.Fprintf(w, "User dengan user_id : %s tidak ditemukan.", user.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = personaliactrl.DeleteUser(ctx, user.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, "Data user berhasil dihapus")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServicesUpsertUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var user personaliaentity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userDB, found, err := personaliactrl.GetUserByUserID(ctx, user.UserID)

	// replace to existing data
	if user.Password == "" {
		user.Password = userDB.Password
	} else {
		user.Password, _ = pkg.HashPassword(user.Password)
	}

	if !found {
		err = personaliactrl.InsertUser(ctx, user)
	} else {
		err = personaliactrl.UpdateUser(ctx, user)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprintf(w, "Data user berhasil disimpan")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
