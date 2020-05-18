package persistence

import (
	"context"
	"fmt"
	"github.com/gegen07/cartola-university/domain/entity"
	"github.com/gegen07/cartola-university/domain/repository"

	"database/sql"
	)

type FormationRepository struct {
	db *sql.DB
}

func NewFormationRepository(db *sql.DB) *FormationRepository {
	return &FormationRepository{db: db}
}

var _ repository.FormationRepository = &FormationRepository{}

func (r FormationRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]entity.Formation, error) {
	formations := make([]entity.Formation, 0)

	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, args)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		f := entity.Formation{}

		err = rows.Scan(
				&f.ID,
				&f.Goalkeeper,
				&f.Attackers,
				&f.Defenders,
				&f.CreatedAt,
				&f.UpdatedAt,
			)

		if err != nil {
			return nil, err
		}

		formations = append(formations, f)
	}

	return formations, nil
}

func (r FormationRepository) GetAll(ctx context.Context, page int) ([]entity.Formation, error) {
	limit := 10
	offset := (page-1)*limit

	query := `SELECT * from formations LIMIT ? OFFSET ?;`

	formations, err := r.fetch(ctx, query, limit, offset)

	if err != nil {
		return nil, err
	}

	return formations, nil
}

func (r FormationRepository) GetByID(ctx context.Context, id uint64) (*entity.Formation, error) {
	query := `SELECT * from formations WHERE id = ?;`

	formations, err := r.fetch(ctx, query, id)

	if err != nil {
		return nil, err
	}

	if len(formations) > 1 {
		// TODO: else with server error
	}

	return &formations[0], nil
}

func (r FormationRepository) Insert(ctx context.Context, formation *entity.Formation) (*entity.Formation, error) {
	query := `INSERT INTO formations (goalkeepers, attackers, defenders, updated_at, created_at)
				VALUES (goalkeepers=?, attackers=?, defenders=?, updated_at=?, created_at=?);`
	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(ctx,
		formation.Goalkeeper,
		formation.Attackers,
		formation.Defenders,
		formation.UpdatedAt,
		formation.CreatedAt,
		)

	if err != nil {
		return nil, err
	}

	return formation, nil
}

func (r FormationRepository) Update(ctx context.Context, formation *entity.Formation) (*entity.Formation, error) {
	query := `UPDATE formations SET (goalkeepers=?, attackers=?, defenders=?,updated_at=?) WHERE id = ?;`
	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	res, err := stmt.ExecContext(ctx,
		formation.Goalkeeper,
		formation.Attackers,
		formation.Defenders,
		formation.UpdatedAt,
		formation.ID)

	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return nil, err
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Weird Behavior %d", rowsAffected)
		return nil, err
	}

	return formation, nil
}

func (r FormationRepository) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM formations WHERE id =?;`
	stmt, err := r.db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Weird Behavior %d", rowsAffected)
		return err
	}

	return nil
}