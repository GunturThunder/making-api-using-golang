package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(appconfig AppConfig, dbconfig DBConfig) {
	fmt.Println("Welcome", appconfig.AppName)
	server.initializeDB(dbconfig)
	server.Router = mux.NewRouter()
}

func (server *Server) initializeDB(dbconfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbconfig.DBHost, dbconfig.DBUser, dbconfig.DBPassword, dbconfig.DBName, dbconfig.DBPort)

	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connecting to the database server")
	}

	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}

	}
	fmt.Println("Database Migrated Successfully.")
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s ", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoTokoWeb")
	appConfig.AppEnv = getEnv("APP_ENV", "Dev")
	appConfig.AppPort = getEnv("App_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "mungkung")
	dbConfig.DBName = getEnv("DB_NAME", "db_belajar")
	dbConfig.DBPort = getEnv("DB_PORT", "8080")
	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
