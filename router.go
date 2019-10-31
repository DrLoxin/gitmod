package gitmod

import "github.com/gorilla/mux"

func GetRouter() *mux.Router{
	return mux.NewRouter()
}