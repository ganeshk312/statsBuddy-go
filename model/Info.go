package model

import (
	"encoding/json"
	"fmt"
	"strconv"
)

/*
"info": {
  "balls_per_over": 6,
  "venue": "Daren Sammy National Cricket Stadium, Gros Islet",
  "dates": [
    "2016-07-26"
  ],
  "event": {
    "name": "Caribbean Premier League",
    "match_number": 24
  },
  "gender": "male",
  "teams": [
    "St Lucia Zouks",
    "Trinbago Knight Riders"
  ],
  "outcome": {
    "winner": "Trinbago Knight Riders",
    "by": {
      "wickets": 3
    }
  },
  "toss": {
    "decision": "field",
    "winner": "Trinbago Knight Riders"
  },
  "player_of_match": [
    "Umar Akmal"
  ],
  "match_type": "T20",
  "overs": 20,
  "team_type": "club",
  "registry": {
    "people": {
      "ADS Fletcher": "1bae756b-ce38-4b99-af93-d34c8dadc379",
      "AP Devcich": "ed90abe4-c676-4501-b7df-ad80935e8e1b",
      "BB McCullum": "b8a55852-72a5-4667-999d-28f845354af9",
      "C Munro": "af2c687b-04d1-4fd0-bcc0-61e7fdbe4ef4",
      "D Ramdin": "d1523761-c316-4e18-92f7-6e048af72422",
      "DE Johnson": "c0babb93-b6a0-43e8-bfe5-00c1a38c3fd3",
      "DJ Bravo": "87e562a9-a160-4e25-bbac-5d70c61b754f",
      "DJG Sammy": "c03c6200-a23c-43ab-8b77-76b71fae379b",
      "Denavon Hayles": "5cbeeb02-fa35-4ee0-8dc8-84873c42044d",
      "GD Elliott": "c03449e0-5c0d-480b-901f-c744a56737e5",
      "J Charles": "09a9d073-9f18-486a-8d67-1ba665a1bae4",
      "JD Cloete": "b0419f8c-68dc-4d18-8727-f51227a9a3be",
      "JE Taylor": "4180d897-7347-482a-a1d2-ac51bf262eae",
      "K Lesporis": "fb496f7a-d8f3-4c59-8ec8-829647161bb9",
      "KK Cooper": "557153ca-f977-4398-a2f0-f47c10b44929",
      "KR Mayers": "73c18486-c4e8-4708-830d-0d7b700361d5",
      "MEK Hussey": "48fd7349-6426-4791-806b-0644a8d8ae5c",
      "NO Miller": "ea3ffd1b-23f2-4490-8796-074cfb5b90ce",
      "PA Gustard": "25bcf808-a2d9-4d8b-8bca-37977305221c",
      "RR Beaton": "684a56df-eaf3-4191-8887-7f4c7cfe55e2",
      "S Shillingford": "d76b0d2d-f791-423b-950f-4a6d022b1db8",
      "SP Narine": "9d430b40-abd3-472f-9dc6-45129c0f41ea",
      "SR Watson": "4329fbb5-8f2d-4fe8-bddb-6f34c9815a4a",
      "Umar Akmal": "fd2bf2a0-1a89-430c-9a7d-f34eaf7b120d",
      "WKD Perkins": "1920c721-f06c-45e9-b8a7-f938774f9fc8",
      "Zahid Bassarath": "ca6bf044-4d96-4a41-94af-9d6f6c6fc46a"
    }
  },
  "city": "St Lucia",
  "officials": {
    "match_referees": [
      "Denavon Hayles"
    ],
    "tv_umpires": [
      "Zahid Bassarath"
    ],
    "umpires": [
      "JD Cloete",
      "PA Gustard"
    ]
  },
  "players": {
    "St Lucia Zouks": [
      "J Charles",
      "ADS Fletcher",
      "SR Watson",
      "MEK Hussey",
      "GD Elliott",
      "DJG Sammy",
      "KR Mayers",
      "DE Johnson",
      "S Shillingford",
      "K Lesporis",
      "JE Taylor"
    ],
    "Trinbago Knight Riders": [
      "WKD Perkins",
      "BB McCullum",
      "C Munro",
      "D Ramdin",
      "Umar Akmal",
      "DJ Bravo",
      "AP Devcich",
      "SP Narine",
      "KK Cooper",
      "NO Miller",
      "RR Beaton"
    ]
  },
  "season": "2016"
}
*/

// MatchInfo represents all the data about a single cricket match.
type MatchInfo struct {
	BallsPerOver    int                 `json:"balls_per_over"`              // The number of balls per over, typically 6
	BowlOut         []*BowlOut          `json:"bowl_out,omitempty"`          // Bowl-out data, if any
	City            string              `json:"city"`                        // City where the match took place
	Dates           []string            `json:"dates"`                       // Dates when the match took place
	Event           *Event              `json:"event,omitempty"`             // Event information
	Gender          string              `json:"gender"`                      // Gender of the players (male, female)
	MatchType       string              `json:"match_type"`                  // Type of the match (Test, ODI, T20, etc.)
	MatchTypeNumber *int                `json:"match_type_number,omitempty"` // Match number for this type (optional)
	Missing         interface{}         `json:"missing,omitempty"`           // List of missing data fields
	Officials       *Officials          `json:"officials,omitempty"`         // Information about match officials
	Outcome         *Outcome            `json:"outcome,omitempty"`           // Outcome of the match
	Overs           int                 `json:"overs"`                       // Number of overs per side
	PlayerOfMatch   []string            `json:"player_of_match,omitempty"`   // Players who were player(s) of the match
	Players         map[string][]string `json:"players"`                     // Players in each team
	Registry        *Registry           `json:"registry,omitempty"`          // Mapping of player names to IDs
	// Season          string              `json:"season"`                      // Season of the match (e.g. 2018)
	Season    Season     `json:"season"`              // Season of the match (e.g. 2018)
	SuperSubs *SuperSubs `json:"supersubs,omitempty"` // Super-subs used in the match
	TeamType  string     `json:"team_type"`           // Type of teams (international, club)
	Teams     []string   `json:"teams"`               // Names of teams in the match
	Toss      *Toss      `json:"toss,omitempty"`      // Toss details (winner and decision)
	Venue     string     `json:"venue"`               // Venue of the match
}
type Season string

func (s *Season) UnmarshalJSON(data []byte) error {
	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	switch v := temp.(type) {
	case string:
		*s = Season(v)
	case float64:
		*s = Season(strconv.Itoa(int(v)))
	default:
		return fmt.Errorf("unexpected type for Season: %T", v)
	}
	return nil
}

// BowlOut represents a single delivery in a bowl-out.
type BowlOut struct {
	Bowler  string `json:"bowler"`  // Name of the bowler
	Outcome string `json:"outcome"` // Outcome of the bowl-out (hit/miss)
}

// Event contains details about the event the match is part of.
type Event struct {
	Name        string `json:"name"`         // Name of the event
	MatchNumber int    `json:"match_number"` // Match number in the event
	// Group       string `json:"group"`        // Group (if any)
	Group   Group  `json:"group,omitempty"` // Group (if any)
	Stage   string `json:"stage"`           // Stage of the event (e.g. Final)
	SubName string `json:"sub_name,omitempty"`
}

type Group string

func (g *Group) UnmarshalJSON(data []byte) error {
	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	switch v := temp.(type) {
	case string:
		*g = Group(v)
	case float64:
		*g = Group(strconv.Itoa(int(v)))
	default:
		return fmt.Errorf("unexpected type for Group: %T", v)
	}
	return nil
}

// Officials stores the details of the match officials.
type Officials struct {
	MatchReferees  []string `json:"match_referees,omitempty"`  // Match referees
	ReserveUmpires []string `json:"reserve_umpires,omitempty"` // Reserve umpires
	TVUmpires      []string `json:"tv_umpires,omitempty"`      // TV umpires
	Umpires        []string `json:"umpires,omitempty"`         // Umpires for the match
}

// Outcome stores the outcome of the match.
type Outcome struct {
	Winner     string   `json:"winner"`               // Winner of the match
	By         *Victory `json:"by,omitempty"`         // Margin of victory (wickets, runs, etc.)
	BowlOut    string   `json:"bowl_out,omitempty"`   // Bowl-out winner (if applicable)
	Eliminator string   `json:"eliminator,omitempty"` // Winner of the super-over eliminator
	Method     string   `json:"method,omitempty"`     // Method used (e.g. D/L, VJD)
	Result     string   `json:"result"`               // Result of the match (draw, no result, tie)
}

// Victory describes the margin of victory.
type Victory struct {
	Innings int `json:"innings,omitempty"` // If won by innings, the value will be 1
	Runs    int `json:"runs,omitempty"`    // Runs by which the match was won
	Wickets int `json:"wickets,omitempty"` // Wickets by which the match was won
}

// Registry maps player names to unique identifiers.
type Registry struct {
	People map[string]string `json:"people"` // Mapping of player name to their unique ID
}

// SuperSubs represents any supersub used by a team.
// type SuperSubs struct {
// 	Teams map[string]string `json:"teams"` // Team name and supersub player
// }

type SuperSubs map[string]string

// Toss stores the result of the toss.
type Toss struct {
	Decision    string `json:"decision,omitempty"`    // The decision made (bat/field)
	Winner      string `json:"winner,omitempty"`      // The team that won the toss
	Uncontested bool   `json:"uncontested,omitempty"` // Whether the toss was uncontested (for some formats)
}
