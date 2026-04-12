package statisticstransporthttp

import (
	"context"
	"net/http"
	"time"

	"github.com/977ADAM/golang-todoapp-project/internal/core/domain"
	corehttpserver "github.com/977ADAM/golang-todoapp-project/internal/core/transport/http/server"
)

type StatisticsHTTPHandler struct {
	statisticsService StatisticsService
}

type StatisticsService interface {
	GetStatistics(
		ctx context.Context,
		userID *int,
		from *time.Time,
		to *time.Time,
	) (domain.Statistics, error)
}

func NewStatisticsHTTPHandler(statisticsService StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		statisticsService: statisticsService,
	}
}

func (h *StatisticsHTTPHandler) Routes() []corehttpserver.Route {
	return []corehttpserver.Route{
		{
			Method:  http.MethodGet,
			Path:    "/statistics",
			Handler: h.GetStatistics,
		},
	}
}
