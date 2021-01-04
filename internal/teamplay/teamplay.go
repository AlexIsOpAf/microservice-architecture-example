package teamplay

import (
	"context"

	"github.com/AlexIsOpAf/microservice-architecture-example/internal"
)

type teamPlayService struct{}

/** NewService returns a new instance of the service that can perform db operations **/
func NewService() Service { return &teamPlayService{} }

/** Get queries the database for all teams **/
func (tp *teamPlayService) Get(ctx context.Context, filter ...internal.Filter) ([]internal.Team, error) {
	// Query database and return all teams - for now we'll just mock it until we sync everything up

	team := []internal.Team{
		CreatedBy: "Alexander",
		Name:      "Aje Raiders",
		Members:   []string{"John", "Damon", "Dukedom"},
	}

	return []internal.Team{team}, nil
}
