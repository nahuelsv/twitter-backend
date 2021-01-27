package routers

import (
	"encoding/json"
	"net/http"

	"github.com/nahuelsv/twitter-backend/bd"
	"github.com/nahuelsv/twitter-backend/models"
)

/*Register is a function to insert our user into the database*/
func Register(w http.ResponseWriter, r *http.Request) {

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid received data: "+err.Error(), 400)
		return
	}

	if len(user.Email) < 8 {
		http.Error(w, "Email required", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Password required. At least 6 chars as minimun", 400)
		return
	}

	_, exist, _ := bd.CheckUser(user.Email)
	if exist == true {
		http.Error(w, "User already registered with this email", 400)
		return
	}

	_, success, err := bd.InsertRegister(user)
	if err != nil {
		http.Error(w, "An error has occurred while tried to insert the user: "+err.Error(), 400)
		return
	}

	if success == false {
		http.Error(w, "Failed to insert user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
