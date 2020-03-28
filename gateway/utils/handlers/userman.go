package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/spaceuptech/space-cloud/gateway/modules"
	"github.com/spaceuptech/space-cloud/gateway/utils"
)

// HandleProfile returns the handler for fetching single user profile
func HandleProfile(modules *modules.Modules) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the path parameters
		vars := mux.Vars(r)
		projectID := vars["project"]
		dbAlias := vars["dbAlias"]
		id := vars["id"]

		userManagement, err := modules.User(projectID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// Create a context of execution
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()
		defer utils.CloseTheCloser(r.Body)

		// Get the JWT token from header
		token := utils.GetTokenFromHeader(r)

		status, result, err := userManagement.Profile(ctx, token, dbAlias, projectID, id)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if err != nil {
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"user": result})
	}
}

// HandleProfiles returns the handler for fetching all user profiles
func HandleProfiles(modules *modules.Modules) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the path parameters
		vars := mux.Vars(r)
		projectID := vars["project"]
		dbAlias := vars["dbAlias"]

		userManagement, err := modules.User(projectID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// Create a context of execution
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		// Get the JWT token from header
		token := utils.GetTokenFromHeader(r)
		defer utils.CloseTheCloser(r.Body)

		w.Header().Set("Content-Type", "application/json")
		status, result, err := userManagement.Profiles(ctx, token, dbAlias, projectID)
		if err != nil {
			w.WriteHeader(status)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(result)
	}
}

// HandleEmailSignIn returns the handler for email sign in
func HandleEmailSignIn(modules *modules.Modules) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projectID := vars["project"]
		dbAlias := vars["dbAlias"]

		userManagement, err := modules.User(projectID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// Create a context of execution
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		// Load the request from the body
		req := map[string]interface{}{}
		_ = json.NewDecoder(r.Body).Decode(&req)
		defer utils.CloseTheCloser(r.Body)

		status, result, err := userManagement.EmailSignIn(ctx, dbAlias, projectID, req["email"].(string), req["pass"].(string))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if err != nil {
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		_ = json.NewEncoder(w).Encode(result)
	}
}

// HandleEmailSignUp returns the handler for email sign up
func HandleEmailSignUp(modules *modules.Modules) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the path parameters
		vars := mux.Vars(r)
		projectID := vars["project"]
		dbAlias := vars["dbAlias"]

		userManagement, err := modules.User(projectID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// Create a context of execution
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		// Load the request from the body
		req := map[string]interface{}{}
		_ = json.NewDecoder(r.Body).Decode(&req)
		defer utils.CloseTheCloser(r.Body)

		status, result, err := userManagement.EmailSignUp(ctx, dbAlias, projectID, req["email"].(string), req["name"].(string), req["pass"].(string), req["role"].(string))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if err != nil {
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		_ = json.NewEncoder(w).Encode(result)
	}
}

// HandleEmailEditProfile returns the handler for edit profile
func HandleEmailEditProfile(modules *modules.Modules) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the path parameters
		vars := mux.Vars(r)
		projectID := vars["project"]
		dbAlias := vars["dbAlias"]
		id := vars["id"]

		userManagement, err := modules.User(projectID)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// Create a context of execution
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		// Get the JWT token from header
		token := utils.GetTokenFromHeader(r)

		// Load the request from the body
		req := map[string]interface{}{}
		_ = json.NewDecoder(r.Body).Decode(&req)
		defer utils.CloseTheCloser(r.Body)

		status, result, err := userManagement.EmailEditProfile(ctx, token, dbAlias, projectID, id, req["email"].(string), req["name"].(string), req["pass"].(string))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if err != nil {
			_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		_ = json.NewEncoder(w).Encode(result)
	}
}
