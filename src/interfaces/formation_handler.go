package interfaces

import (
	"encoding/json"
	"github.com/gegen07/cartola-university/application"
	"log"
	"net/http"
)

type FormationHandler struct {
	formationApplication application.FormationApplicationInterface
}

func NewFormationHandler(formationApp application.FormationApplicationInterface) *FormationHandler {
	return &FormationHandler{
		formationApplication: formationApp,
	}
}
/**
	GetAll(args ...interface{}) ([]entity.Formation, error)
	GetByID(id uint64) (*entity.Formation, error)
	Insert(formation *entity.Formation) (*entity.Formation, error)
	Update(formation *entity.Formation) (*entity.Formation, error)
	Delete(formation *entity.Formation) error
 */
func (handler *FormationHandler) GetFormations (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	formations, err := handler.formationApplication.GetAll()

	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(400)
	}

	json.NewEncoder(w).Encode(formations)
}