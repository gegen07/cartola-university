package round

import (
	"context"
	"database/sql"
	"log"
	"fmt"

	"github.com/gegen07/cartola-coltec/app/cartola-coltec-service/models"
)

type pgsqlRoundRepository struct {
	Db *sql.DB
}

func NewPgsqlRoundRepository(Db *sql.DB) Repository {
	return &pgsqlRoundRepository{Db}
}

func (p *pgsqlRoundRepository) GetAll(ctx context.Context, args ...interface{}) ([]*models.Round, error) { //GetAll return an array of rounds
	query := `SELECT r.id, r.round_begin_date, r.round_finish_date FROM round r;`

	rows, err := p.Db.QueryContext(ctx, query, args)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	rounds := make([]*models.Round, 0)

	for rows.Next() {
		r := new(models.Round)
		err := rows.Scan(
			&r.ID,
			&r.RoundBeginDate,
			&r.RoundFinishDate,
		)

		if err != nil {
			log.Fatal(err)
		}

		rounds = append(rounds, r) 
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return rounds, err
}

//GetByID return round w/ match 
func (p *pgsqlRoundRepository) GetByID(ctx context.Context, id int64) (*models.Round, error) {
	roundQuery := `SELECT r.id, r.round_begin_date, r.round_finish_date FROM round r WHERE r.id = ?;`

	row := p.Db.QueryRowContext(ctx, roundQuery, id)
	
	round := new(models.Round)
	err := row.Scan (
		&round.ID,
		&round.RoundBeginDate,
		&round.RoundFinishDate,
	)

	if err != nil {
		log.Fatal(err)
	}

	return round, err
}

func (p *pgsqlRoundRepository) Insert(ctx context.Context, round *models.Round) error {
	query := `INSERT INTO round (round_begin_date, round_finish_date) VALUES (?, ?);`

	stmt, err := p.Db.PrepareContext(ctx, query)
	
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, round.RoundBeginDate, round.RoundFinishDate)

	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (p *pgsqlRoundRepository) Update(ctx context.Context, round *models.Round) error {
	query := `UPDATE round SET round_begin_date = ?, round_finish_date = ?;`
	
	stmt, err := p.Db.PrepareContext(ctx, query)

	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, round.RoundBeginDate, round.RoundFinishDate)

	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if affect != 1 {
		err = fmt.Errorf("Weird behaviour. Total Rows Affected: %d", affect)

		return err
	}

	return nil
}

func (p *pgsqlRoundRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM round WHERE id = ?;`

	stmt, err := p.Db.PrepareContext(ctx, query)

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
		err = fmt.Errorf("Weird behaviour. Total Rows Affected: %d", rowsAffected)

		return err
	}

	return nil
}


