package players

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
	table = "player"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Players, error) {

	var players []models.Players

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
		var player models.Players
		// var createdAt, updatedAt string

		if err = rowQuery.Scan(&player.ID,
			&player.Name); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		// Player.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		// mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		players = append(players, player)
	}

	return players, nil
}

//insert to Player
func InsertPlayer(ctx context.Context, players models.Players) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (team_id,name) values('%d','%v')", table,
		players.TeamID, players.Name)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

//update Player
func UpdatePlayer(ctx context.Context, players models.Players) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set name ='%s' where id = '%d'",
		table,
		players.Name,
		players.ID,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete
func DeletePlayer(ctx context.Context, t models.Players) error {

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
