package teamplay

import (
	"context"

	"github.com/AlexIsOpAf/microservice-architecture-example/internal"
)

/* Context enables microservices to access each goroutine request */
type Service interface {
	/** Return a list of all registered teams **/
	Get(ctx context.Context, filter ...internal.Filter) ([]internal.Team, error)

	/** Create the team and return the name of the team **/
	Create(ctx context.Context, team *internal.Team) (string, error)

	/** Add user to team **/
	Add(ctx context.Context, uuid string) error
}
