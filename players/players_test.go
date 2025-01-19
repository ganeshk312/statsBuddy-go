package players

import (
	"database/sql"
	"fmt"
	"testing"
)

func setupTestDB(t *testing.T) *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/players?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func TestSearchPlayers(t *testing.T) {
	// db := setupTestDB(t)
	// defer db.Close()

	tests := []struct {
		name       string
		searchType string
		value      string
		want       interface{}
		wantErr    bool
	}{
		{
			name:       "Find Kohli by ID",
			searchType: "id",
			value:      "ba607b88",
			want:       &Player{ID: "ba607b88", Name: "V Kohli"},
			wantErr:    false,
		},
		{
			name:       "Find Bumrah by ID",
			searchType: "id",
			value:      "462411b3",
			want:       &Player{ID: "462411b3", Name: "JJ Bumrah"},
			wantErr:    false,
		},
		{
			name:       "Find players by name pattern",
			searchType: "name",
			value:      ".*Kohli.*",
			want: []Player{
				{ID: "ba607b88", Name: "V Kohli"},
			},
			wantErr: false,
		},
		{
			name:       "Invalid search type",
			searchType: "invalid",
			value:      "test",
			want:       nil,
			wantErr:    true,
		},
		{
			name:       "ID not found",
			searchType: "id",
			value:      "nonexistent",
			want:       nil,
			wantErr:    false,
		},
		{
			name:       "All name containing khan case insensitive",
			searchType: "name",
			value:      ".*khan.*",
			want: []Player{
				{ID: "462411b3", Name: "JJ Bumrah"},
				{ID: "ba607b88", Name: "V Kohli"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchPlayers(tt.searchType, tt.value)

			if (err != nil) != tt.wantErr {
				t.Errorf("SearchPlayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			switch v := got.(type) {
			case *Player:
				if want, ok := tt.want.(*Player); ok && v != nil {
					if *v != *want {
						t.Errorf("SearchPlayers() got = %v, want %v", v, want)
					}
				}
			case []Player:
				if want, ok := tt.want.([]Player); ok {
					if len(v) < len(want) {
						t.Errorf("SearchPlayers() got len = %v, want len %v", len(v), len(want))
					}
				}
			}
		})
	}
}

func TestSearchPlayersInteractive(t *testing.T) {
	var searchType, value string
	fmt.Println("Enter search type (1 for ID, 2 for Name Pattern):")
	// fmt.Scanln(&searchType)
	searchType = "1"
	if searchType == "1" {
		searchType = "id"
	} else if searchType == "2" {
		searchType = "name"
	} else {
		t.Fatalf("Invalid search type")
	}

	fmt.Println("Enter search value:")
	// fmt.Scanln(&value)
	value = "1bae756b"
	got, err := SearchPlayers(searchType, value)

	fmt.Printf("Search result: %v\n", got)
	fmt.Printf("Err: %v", err)
}
