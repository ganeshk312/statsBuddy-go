package model

type CricketMatch struct {
	Meta    MetaInfo  `json:"meta"`
	Info    MatchInfo `json:"info"`
	Innings []Inning  `json:"innings"`
}
