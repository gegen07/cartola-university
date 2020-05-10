package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gegen07/cartola-university/application"
	"github.com/gegen07/cartola-university/domain/entity/scout"
	"github.com/gegen07/cartola-university/interfaces/errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PositionHandler struct {
	PositionApplication application.PositionApplicationInterface
}

func NewPositionHandler(positionApplication application.PositionApplicationInterface) *PositionHandler {
	return &PositionHandler{
		PositionApplication: positionApplication,
	}
}

func (handler *PositionHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return errors.NewHTTPError(nil, http.StatusMethodNotAllowed, "Method Not Allowed")
	}

	w.Header().Set("Content-Type", "application/json")

	var positions scout.Positions

	positions, err := handler.PositionApplication.GetAll()

	if err != nil {
		return fmt.Errorf("DB error: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(positions.PublicPositions())

	return nil
}

func (handler *PositionHandler) Insert(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errors.NewHTTPError(nil, http.StatusMethodNotAllowed, "Method Not Allowed")
	}

	w.Header().Set("Content-Type", "application/json")

	var position scout.Position

	err := json.NewDecoder(r.Body).Decode(&position)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, "Invalid JSON")
	}

	p, err := handler.PositionApplication.Insert(&position)

	if err != nil {
		return fmt.Errorf("DB error: %v", err)
	}

	json.NewEncoder(w).Encode(p.PublicPosition())
	w.WriteHeader(http.StatusCreated)

	return nil
}

func (handler *PositionHandler) GetByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		errors.NewHTTPError(nil, http.StatusMethodNotAllowed, "Method Not Allowed")
	}

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, "Invalid JSON")
	}

	position, err := handler.PositionApplication.GetById(id)

	if err != nil {
		return fmt.Errorf("DB Error: %v", err)
	}

	json.NewEncoder(w).Encode(position.PublicPosition())
	w.WriteHeader(http.StatusOK)
	return nil
}

func (handler *PositionHandler) Update(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPut {
		return errors.NewHTTPError(nil, http.StatusMethodNotAllowed, "Method Not Allowed")
	}

	w.Header().Set("Content-Type", "application/json")

	var position scout.Position

	err := json.NewDecoder(r.Body).Decode(&position)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, "Invalid JSON")
	}

	params := mux.Vars(r)

	position.ID, err = strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, "Invalid JSON")
	}

	p, err := handler.PositionApplication.Update(&position)

	if err != nil {
		return fmt.Errorf("DB Error: %v", err)
	}

	json.NewEncoder(w).Encode(p)
	w.WriteHeader(http.StatusOK)

	return nil
}

func (handler PositionHandler) Delete(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodDelete {
		return errors.NewHTTPError(nil, http.StatusMethodNotAllowed, "Method Not Allowed")
	}

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		return errors.NewHTTPError(err, http.StatusBadRequest, "Invalid JSON")
	}

	err = handler.PositionApplication.Delete(id)

	if err != nil {
		return fmt.Errorf("DB Error: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	return nil
}