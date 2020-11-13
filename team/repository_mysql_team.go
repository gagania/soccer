package team

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"test-api/config"
	"test-api/models"
)

const (
	table = "team"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Team, error) {

	var teams []models.Team

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var team models.Team
		// var createdAt, updatedAt string

		if err = rowQuery.Scan(&team.ID,
			&team.Name); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		// team.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		// mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		teams = append(teams, team)
	}

	return teams, nil
}

//insert to team
func Insert(ctx context.Context, team models.Team) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (name) values('%v')", table,
		team.Name)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

//update team
func UpdateTeam(ctx context.Context, team models.Team) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s' where id = '%d'",
		table,
		team.Name,
		team.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete
func DeleteTeam(ctx context.Context, t models.Team) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, t.ID)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	return nil
}
