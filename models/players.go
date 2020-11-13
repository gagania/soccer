package models

type (
	//players
	Players struct {
		ID     int    `json:"id"`
		TeamID int    `json:"team_id"`
		Name   string `name:"name"`
	}
)
