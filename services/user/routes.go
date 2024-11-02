package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KengoWada/gorouting/utils"
)

type Handler struct{}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User = []User{
	{ID: 1, Name: "Bumblebee", Email: "bumblebee@autobots.com"},
	{ID: 2, Name: "Optimus Prime", Email: "optimus.prime@autobots.com"},
	{ID: 3, Name: "Ironhide", Email: "ironhide@autobots.com"},
	{ID: 4, Name: "Hot Rod", Email: "hot.rod@autobots.com"},
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /", getUsers)
	r.HandleFunc("POST /", createUser)

	r.HandleFunc("GET /{id}/", getUser)
	r.HandleFunc("DELETE /{id}/", deleteUser)

	return r
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{
		"message": "Done",
		"users":   users,
	}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response := map[string]any{
			"message": "Invalid user id. User id must be of type int.",
		}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	for _, user := range users {
		if user.ID == userID {
			response := map[string]any{
				"message": "Done",
				"user":    user,
			}
			utils.WriteJSONResponse(w, http.StatusOK, response)
			return
		}
	}

	response := map[string]any{"message": fmt.Sprintf("User with id: %d not found", userID)}
	utils.WriteJSONResponse(w, http.StatusNotFound, response)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := utils.ParseJSON(r, &user); err != nil {
		response := map[string]any{"message": "Invalid request"}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	users = append(users, user)
	response := map[string]any{"message": "Done"}
	utils.WriteJSONResponse(w, http.StatusCreated, response)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		response := map[string]any{
			"message": "Invalid user id. User id must be of type int.",
		}
		utils.WriteJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	userIndex := -1
	for index, user := range users {
		if userID == user.ID {
			userIndex = index
			break
		}
	}

	if userIndex == -1 {
		response := map[string]any{"message": fmt.Sprintf("User with id: %d not found", userID)}
		utils.WriteJSONResponse(w, http.StatusNotFound, response)
		return
	}

	users = append(users[:userIndex], users[userIndex+1:]...)
	response := map[string]any{"message": fmt.Sprintf("Deleted user with id: %d", userID)}
	utils.WriteJSONResponse(w, http.StatusNotFound, response)
}
