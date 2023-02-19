package app

import (
	"fmt"
	"github.com/Mersock/project-timesheet-backend/config"
	"github.com/Mersock/project-timesheet-backend/pkg/logger"
	"github.com/Mersock/project-timesheet-backend/pkg/mysql"
	"github.com/gin-gonic/gin"
	"os"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.LOG.Level)

	//Repository
	sql, err := mysql.NewMysqlConn(cfg.MYSQL.URL)
	if err != nil {
		l.Fatal(fmt.Errorf("app - RUN - mysql.NewMysqlConn: %w", err))
	}
	defer sql.Close()

	//HTTP server
	gin.New()

	//signal interrupt
	interrupt := make(chan os.Signal, 1)
	select {
	case s := <-interrupt:
		l.Info("app - RUN -signal: " + s.String())
	}
}
