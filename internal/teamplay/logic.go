package teamplay

import (
	"context"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/lithammer/shortuuid/v3"
)

/** Empty struct that implements the service interface - Struct implementation hidden: So realistically we could add config to this **/
type teamPlayService struct {
	repo   Repository
	logger log.Logger
}

/** NewService returns a new instance of the service that can perform db operations **/
func NewService(rep Repository, log log.Logger) Service {
	return &teamPlayService{
		repo:   rep,
		logger: log,
	}
}

/** Get queries the database for specific team **/
func (tp *teamPlayService) GetWithID(ctx context.Context, uuid string) (Team, error) {

	logger := log.With(tp.logger, "method", "Get")

	team, err := tp.repo.GetTeam(ctx, uuid)

	if err != nil {
		level.Error(logger).Log("err", err)
		return Team{}, err
	}

	logger.Log("GetWithID", team.TeamID)

	return team, nil
}

func (tp *teamPlayService) Create(ctx context.Context, teamName, playerName string) (string, error) {
	// Query the database and insert new team - for now we'll just mock it until we sync everything up

	logger := log.With(tp.logger, "method", "Create")

	team := Team{
		TeamID:    shortuuid.New(),
		CreatedBy: playerName,
		Name:      teamName,
		Members:   []string{playerName},
	}

	if err := tp.repo.CreateTeam(ctx, team); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create team", team.TeamID)
	return team.TeamID, nil
}

func (tp *teamPlayService) Add(ctx context.Context, teamID, userID string) error {
	// query db to add player's ID to team's ID
	return nil
}

func (tp *teamPlayService) ServiceStatus(ctx context.Context) (int, error) {
	/* Maybe something a little more sexy if we were running a real service in production */
	logger.Log("Service health..")
	return http.StatusOK, nil
}

func init() {
	/* Simple logger that is safe for concurrent use by multiple go routines outputs error to POSIX stderr */
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
