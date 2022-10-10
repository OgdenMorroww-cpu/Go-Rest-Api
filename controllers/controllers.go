package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type UserModel struct {
	ID            string `json:"id"`
	FirstName     string `json:"first_name"`
	MiddleName    string `json:"middle_name"`
	LastName      string `json:"last_name"`
	Age           string `json:"age"`
	Email         string `json:"email"`
	MaritalStatus string `json:"marital_status"`
}

type getAllUsers []UserModel

var users = getAllUsers{
	{
		ID:            "1",
		FirstName:     "Elon",
		MiddleName:    "Reeve",
		LastName:      "Musk",
		Age:           "50",
		Email:         "Elon@spaceX.com",
		MaritalStatus: "false",
	},
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
}

func CreateNewUsers(w http.ResponseWriter, r *http.Request) {
	var newUser UserModel
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with your full info")
	}
	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func GetNewUserID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	for _, singleUser := range users {
		if singleUser.ID == userID {
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var updateUsers UserModel
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter your full info")
	}
	json.Unmarshal(reqBody, &updateUsers)

	for i, singleUser := range users {
		if singleUser.ID == userID {
			singleUser.FirstName = updateUsers.FirstName
			singleUser.MiddleName = updateUsers.MiddleName
			singleUser.LastName = updateUsers.LastName
			singleUser.Age = updateUsers.Age
			singleUser.Email = updateUsers.Email
			singleUser.MaritalStatus = updateUsers.MaritalStatus
			users = append(users[:i], singleUser)
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func DeleteUsers(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]

	for i, singleUser := range users {
		if singleUser.ID == userID {
			users = append(users[:i], users...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", userID)
		}
	}
}
