package main

import (
	"github.com/rs/cors"
	"net/http"
	"userServer/handle"
)

func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.Write([]byte("{\"hello\": \"world\"}"))
	//})

	mux.HandleFunc("/dev-api/user/login", handle.UserLogin)
	mux.HandleFunc("/dev-api/user/logout", handle.UserLogout)
	mux.HandleFunc("/dev-api/user/register", handle.UserRegister)
	mux.HandleFunc("/dev-api/user/change-password", handle.UserUpdatePassword)
	mux.HandleFunc("/dev-api/user/info", handle.UserInfo)
	mux.HandleFunc("/dev-api/role/fetch-roles", handle.FetchRoles)
	mux.HandleFunc("/dev-api/user/fetch-users", handle.FetchUsers)
	mux.HandleFunc("/dev-api/user/update-user-roles", handle.UpdateUserRoles)
	mux.HandleFunc("/dev-api/role/update-role", handle.UpdateRole)


	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	c := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://localhost:9531"},
		AllowedOrigins:   []string{"http://localhost:*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	// Insert the middleware
	handler := c.Handler(mux)
	http.ListenAndServe(":9090", handler)
}
