package app

import (
	"github.com/rs/zerolog/log"
	"svc-user/internal/repository"
	"svc-user/internal/router"
	"svc-user/internal/service"
	"svc-user/internal/store"
)

func Run(port string) error {
	//db := store.NewSQLLite()
	db := store.NewPostgreSQLDb().Connect()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	err := router.NewRouter(port, userService).Listen()
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
		return err
	}

	return nil
}
