package teamplay

import "context"

type Team struct {
	TeamID    string   `json:"ID"`
	CreatedBy string   `json:"createdby"`
	Name      string   `json:"name"`
	Members   []string `json:"members"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type User struct {
	ID         string `json:"ID"`
	PlayerName string `json:"playerName,omitempty"`
}

/** Helps us interface with our database **/
type Repository interface {
	CreateTeam(ctx context.Context, t Team) error
	GetTeam(ctx context.Context, uuid string) (Team, error)
	AddMember(ctx context.Context, uuidTeam string, member User) error
}
