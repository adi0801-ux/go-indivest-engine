package main

import (
	"indivest-engine/api"
	"indivest-engine/db"
	"indivest-engine/redis"
	"indivest-engine/repositories"
	"indivest-engine/services"
	"indivest-engine/utils"
	"log"
)

func main() {
	err := utils.InitialiseLogger()
	if err != nil {
		log.Fatalln(err)
	}
	utils.Log.Info("logger initialized")

	//loading config from env file
	//**live loading for config added **
	utils.Log.Info("config loading...")
	config, err := utils.LoadConfig(".")
	if err != nil {
		utils.Log.Fatal(err)
		return
	}

	utils.Log.Info("config loaded")

	utils.Log.Info("database connecting....")

	//Make DB Connection
	store, err := makeDBConnection(config)
	if err != nil {
		utils.Log.Fatal(err)
		return
	}
	// Make Migrations
	err = store.RunMigrations()

	if err != nil {
		utils.Log.Fatal("error creating migrations")
	}

	utils.Log.Info("database connected")

	// Make Redis Connection
	utils.Log.Info("redis connecting....")
	redisStore, err := makeRedisConnection(config)

	if err != nil {
		utils.Log.Fatal(err)
		return
	}
	utils.Log.Info("redis connected")

	//create repository references

	//Create a Repository Reference
	userRep := repositories.UserDetailsRepository{
		Db: store,
	}
	sandboxRep := repositories.SandboxRepository{
		Db: store,
	}

	redisRepo := repositories.RedisRepository{
		Db: redisStore,
	}

	//Create a service Reference
	Srv := services.ServiceConfig{
		UserRep: &userRep,
	}

	sandboxSrv := services.SandboxServiceConfig{
		SandboxRep: &sandboxRep,
		RedisRep:   &redisRepo,
	}

	//creating a config

	utils.Log.Info("api server initializing")
	//Create HTTP Server
	server := api.GetNewServer(&Srv, &sandboxSrv, config)

	err = server.StartServer(config.ServerAddress)
	if err != nil {
		utils.Log.Fatal("cannot start server: ", err)
	}

	utils.Log.Info("api server initialized")

	// Results:
	// Name: addUser, Method: GET
	// Name: destroyUser, Method: DELETE

}

func makeDBConnection(config *utils.Config) (*db.Database, error) {
	dbConfig := &db.ConnectionConfig{
		DSN: config.DSN,
	}

	database, err := db.ConnectToDB(dbConfig)
	return database, err
}

func makeRedisConnection(config *utils.Config) (*redis.Client, error) {
	redisConfig := &redis.ConnectionConfig{
		Address:  config.RedisAddress,
		Password: config.RedisPassword,
		Username: config.RedisUserName,
		DBName:   config.RedisDb,
	}
	client, err := redis.ConnectToRedis(redisConfig)

	return client, err
}
