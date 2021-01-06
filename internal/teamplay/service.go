package teamplay

import (
	"context"
)

/* Methods exposed from our microservice */
type Service interface {
	/** Return a list of all registered teams **/
	GetWithID(ctx context.Context, uuid string) (Team, error)

	/** Create the team and return the uuid of the team **/
	Create(ctx context.Context, teamName, playerName string) (string, error)

	/** Add user to team **/
	Add(ctx context.Context, teamID, userID string) error

	/* Returns stats of the service */
	ServiceStatus(ctx context.Context) (int, error)
}
