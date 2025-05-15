package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkOutHandler struct {
}

func NewWorkoutHandler() *WorkOutHandler {
	return &WorkOutHandler{}
}

func (wh *WorkOutHandler) HandleWorkoutById(w http.ResponseWriter, r *http.Request) {

	paramsWorkOutId := chi.URLParam(r, "id")

	if paramsWorkOutId == "" {
		http.NotFound(w, r)
		return
	}

	workOutID, err := strconv.ParseInt(paramsWorkOutId, 10, 64)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "this the given workout id %d\n", workOutID)

}

func (wh *WorkOutHandler) CreateWorkOut(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Workout has been created\n")
}
