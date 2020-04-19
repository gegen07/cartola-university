package interfaces

import (
	"encoding/json"
	"github.com/gegen07/cartola-university/application"
	"github.com/gegen07/cartola-university/domain/entity"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type FormationHandler struct {
	formationApplication application.FormationApplicationInterface
}

func NewFormationHandler(formationApp application.FormationApplicationInterface) *FormationHandler {
	return &FormationHandler{
		formationApplication: formationApp,
	}
}

func (handler *FormationHandler) GetAllFormations (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	formations, err := handler.formationApplication.GetAll()

	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(formations)
}

func (handler *FormationHandler) GetFormationById (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	formationId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	formation, err := handler.formationApplication.GetByID(formationId)

	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(formation.PublicFormation())
}

func (handler *FormationHandler) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var formation entity.Formation
	_ = json.NewDecoder(r.Body).Decode(&formation)

	f, err := handler.formationApplication.Insert(&formation)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(f.PublicFormation())
}

func (handler *FormationHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var formation entity.Formation
	var err error

	params := mux.Vars(r)
	formation.ID, err = strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&formation)

	f, err := handler.formationApplication.Update(&formation)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(f.PublicFormation())
}

func (handler *FormationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	formationId, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.formationApplication.Delete(formationId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}