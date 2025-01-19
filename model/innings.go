package model

/*
{
  "innings": [
    {
      "team": "Ireland",
      "overs": [
        {
          "over": 0,
          "deliveries": [
            {
              "batter": "WTS Porterfield",
              "bowler": "IK Pathan",
              "extras": {
                "wides": 1
              },
              "non_striker": "JP Bray",
              "runs": {
                "batter": 0,
                "extras": 1,
                "total": 1
              }
            },
            {
              "batter": "WTS Porterfield",
              "bowler": "IK Pathan",
              "non_striker": "JP Bray",
              "runs": {
                "batter": 0,
                "extras": 0,
                "total": 0
              }
            }
          ]
        }
      ]
    },
    {
      "team": "India",
      "overs": [
        {
          "over": 0,
          "deliveries": [
            {
              "batter": "G Gambhir",
              "bowler": "WB Rankin",
              "non_striker": "RG Sharma",
              "runs": {
                "batter": 4,
                "extras": 0,
                "total": 4
              }
            }
          ]
        }
      ]
    }
  ]
}
*/

// Inning represents a single innings in a cricket match.
type Inning struct {
	Team            string           `json:"team"`
	Overs           []Over           `json:"overs"`
	AbsentHurt      []string         `json:"absent_hurt,omitempty"`
	PenaltyRuns     *PenaltyRuns     `json:"penalty_runs,omitempty"`
	Declared        bool             `json:"declared,omitempty"`
	Forfeited       bool             `json:"forfeited,omitempty"`
	PowerPlays      []PowerPlay      `json:"powerplays,omitempty"`
	MiscountedOvers *MiscountedOvers `json:"miscounted_overs,omitempty"`
	Target          *Target          `json:"target,omitempty"`
	SuperOver       bool             `json:"super_over,omitempty"`
}

// Over represents an over within an innings.
type Over struct {
	Over       int        `json:"over"`
	Deliveries []Delivery `json:"deliveries"`
}

// Delivery represents a single delivery in an over.
type Delivery struct {
	Batter       string        `json:"batter"`
	Bowler       string        `json:"bowler"`
	NonStriker   string        `json:"non_striker"`
	Runs         *Runs         `json:"runs"`
	Wickets      []Wicket      `json:"wickets,omitempty"`
	Extras       *Extras       `json:"extras,omitempty"`
	Replacements *Replacements `json:"replacements,omitempty"`
	Review       *Review       `json:"review,omitempty"`
}

// Runs represents runs scored on a delivery.
type Runs struct {
	Batter      int  `json:"batter"`
	Extras      int  `json:"extras"`
	Total       int  `json:"total"`
	NonBoundary bool `json:"non_boundary,omitempty"`
}

// Wicket represents a wicket taken during a delivery.
type Wicket struct {
	Kind      string    `json:"kind"`
	PlayerOut string    `json:"player_out"`
	Fielders  []Fielder `json:"fielders,omitempty"`
}
type Fielder struct {
	Name string `json:"name"`
	// Substitute bool   `json:"substitute,omitempty"`
}

// Extras represents the extras conceded on a delivery.
type Extras struct {
	Byes    int `json:"byes,omitempty"`
	LegByes int `json:"legbyes,omitempty"`
	NoBalls int `json:"noballs,omitempty"`
	Penalty int `json:"penalty,omitempty"`
	Wides   int `json:"wides,omitempty"`
}

// PenaltyRuns represents penalty runs added to the innings.
type PenaltyRuns struct {
	Pre  int `json:"pre,omitempty"`
	Post int `json:"post,omitempty"`
}

// PowerPlay represents a powerplay period within an innings.
type PowerPlay struct {
	From float64 `json:"from"`
	To   float64 `json:"to"`
	Type string  `json:"type"`
}

// MiscountedOvers represents miscounted overs in an innings.
type MiscountedOvers map[int]MiscountedOver

// MiscountedOver represents details of a single miscounted over.
type MiscountedOver struct {
	Balls  int    `json:"balls"`
	Umpire string `json:"umpire,omitempty"`
}

// Target represents the target score and overs for an innings.
type Target struct {
	Overs int `json:"overs"`
	Runs  int `json:"runs"`
}

// Replacements represents player replacements in a delivery.
type Replacements struct {
	Match []MatchReplacement `json:"match,omitempty"`
	Role  []RoleReplacement  `json:"role,omitempty"`
}

// MatchReplacement represents a replacement of a player in the match.
type MatchReplacement struct {
	In     string `json:"in"`
	Out    string `json:"out"`
	Reason string `json:"reason"`
	Team   string `json:"team"`
}

// RoleReplacement represents a replacement of a specific role in the game.
type RoleReplacement struct {
	In     string `json:"in"`
	Out    string `json:"out,omitempty"`
	Reason string `json:"reason"`
	Role   string `json:"role"`
}

// Review represents details of a review during a delivery.
type Review struct {
	Batter      string `json:"batter"`
	By          string `json:"by"`
	Decision    string `json:"decision"`
	Umpire      string `json:"umpire,omitempty"`
	UmpiresCall bool   `json:"umpires_call,omitempty"`
}
