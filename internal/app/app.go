package app

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/config"
	v1 "github.com/Mersock/project-timesheet-backend/internal/controller/http/v1"
	"github.com/Mersock/project-timesheet-backend/internal/usecase"
	"github.com/Mersock/project-timesheet-backend/internal/usecase/repo"
	"github.com/Mersock/project-timesheet-backend/internal/utils"
	"github.com/Mersock/project-timesheet-backend/pkg/httpserver"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/Mersock/project-timesheet-backend/pkg/mysql"
	"github.com/Mersock/project-timesheet-backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"os"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.LOG.Level)

	//Repository
	db, err := mysql.NewMysqlConn(cfg.MYSQL.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - RUN - mysql.NewMysqlConn: %w", err))
	}
	defer db.Close()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("iso8601date", utils.IsISO8601Date)
	}

	tokenMaker, err := token.NewJWTMaker(cfg.KEY.TokenSymmetric)
	if err != nil {
		l.Fatal(fmt.Errorf("app - RUN - token.NewJWTMaker: %w", err))
	}
	//Use case
	rolesUseCase := usecase.NewRolesUseCase(
		repo.NewRolesRepo(db),
	)
	userUseCase := usecase.NewUsersUseCase(
		repo.NewUsersRepo(db),
	)
	authUseCase := usecase.NewAuthUseCase(
		repo.NewUsersRepo(db),
		tokenMaker,
		cfg,
	)
	projectUseCase := usecase.NewProjectsUseCase(
		repo.NewProjectRepo(db),
		repo.NewDutiesRepo(db),
		repo.NewWorkTypesRepo(db),
	)
	workTypeUseCase := usecase.NewWorkTypesUseCase(
		repo.NewWorkTypesRepo(db),
	)
	statusUseCase := usecase.NewStatusUseCase(
		repo.NewStatusRepo(db),
	)

	//HTTP server
	handler := gin.New()
	v1.NewRouter(
		handler,
		l,
		tokenMaker,
		rolesUseCase,
		userUseCase,
		authUseCase,
		projectUseCase,
		workTypeUseCase,
		statusUseCase,
	)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	//signal interrupt
	interrupt := make(chan os.Signal, 1)
	select {
	case s := <-interrupt:
		l.Info("app - RUN - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown httpserver
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
