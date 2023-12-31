package main

import (
	"indivest-engine/api"
	"indivest-engine/cron"
	"indivest-engine/db"
	"indivest-engine/redis"
	"indivest-engine/repositories"
	"indivest-engine/seeders"
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

	utils.Log.Info("running seeders ..")
	err = seeders.RunSeeders(store)
	if err != nil {
		utils.Log.Fatal("error running seeders")
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

	sandboxRep := repositories.SandboxRepository{
		Db: store,
	}
	SavvyRepo := repositories.SavvyRepository{
		Db: store,
	}
	UserRepo := repositories.UserRepository{
		Db: store,
	}

	redisRepo := repositories.RedisRepository{
		Db: redisStore,
	}
	TSARepo := repositories.TSAClient{Client: repositories.CreateHttpClient(),
		LogRep:  &repositories.ApiLogsRepository{Db: store},
		BaseUrl: config.SavvyUrl,
		Token:   &config.SavvyToken,
	}

	//Create a service Reference
	RiskSrv := services.RiskCalculatorService{
		UserRepo: &UserRepo,
	}

	sandboxSrv := services.SandboxServiceConfig{
		SandboxRep: &sandboxRep,
		RedisRep:   &redisRepo,
	}

	MfSrv := services.MFService{
		TSAClient: &TSARepo,
		SavvyRepo: &SavvyRepo,
	}

	UserSrv := services.UserSrv{
		UserRepo: &UserRepo,
		Config:   config,
	}

	//create cron Reference
	//err = MfSrv.UpdateFundHouses()
	//err = MfSrv.UpdateFunds()
	//if err != nil {
	//	utils.Log.Info(err)
	//	return
	//}

	cronRef := cron.Cron{SandboxSrv: &sandboxSrv, Sc: cron.CreateScheduler(), MfSrv: &MfSrv}
	// run cron jobs
	cronRef.InitializeScheduler()

	//creating a config

	utils.Log.Info("api server initializing")
	//Create HTTP Server
	server := api.GetNewServer(&RiskSrv, &sandboxSrv, &MfSrv, &UserSrv, config)

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
