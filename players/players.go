package players

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	db   *sql.DB
	once sync.Once
)

func getDB() *sql.DB {
	once.Do(func() {
		connStr := "postgres://postgres:postgres@localhost:5432/players?sslmode=disable"
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
	})
	return db
}

func GetPlayerByID(id string) (*Player, error) {
	db := getDB()
	player := &Player{}
	err := db.QueryRow("SELECT identifier, name FROM players WHERE identifier = $1", id).Scan(&player.ID, &player.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return player, err
}

func GetPlayersByNamePattern(namePattern string) ([]Player, error) {
	db := getDB()
	players := []Player{}
	rows, err := db.Query("SELECT identifier, name FROM players WHERE name ~* $1", namePattern)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.ID, &player.Name); err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

func SearchPlayers(searchType string, value string) (interface{}, error) {
	switch searchType {
	case "id":
		return GetPlayerByID(value)
	case "name":
		return GetPlayersByNamePattern(value)
	default:
		return nil, fmt.Errorf("invalid search type: %s", searchType)
	}
}
