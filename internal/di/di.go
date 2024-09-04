package di

import (
	"basiccrud/internal/config"
	"basiccrud/internal/database"
	"basiccrud/internal/routes"
	"basiccrud/internal/server"
	"basiccrud/internal/user/delivery"
	"basiccrud/internal/user/repository"
	"basiccrud/internal/user/usecase"
)

func Init(conf config.Config) *server.ServerStruct {
	server := server.NewHTTPServer()
	db := db.ConnectPGDB(conf)
	repo := repository.NewUserRepository(db)
	usecase := usecase.UseCase(repo)
	delivery := delivery.UserDelivery(usecase)
	userRoutes := routes.NewUserInit(server, delivery)
	userRoutes.UserRoutes()
	return server
}
