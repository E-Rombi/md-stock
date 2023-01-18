package infrastructure

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	application "md-stock/internal/application/product/create"
	api "md-stock/internal/infrastructure/api"
	infrastructure "md-stock/internal/infrastructure/product"
)

type Application struct {
	server *echo.Echo
}

func NewApplication(server *echo.Echo) *Application {
	return &Application{
		server: server,
	}
}

func (app *Application) Start() {
	configuration := app.buildConfig()

	db := app.buildDB(configuration)

	app.setUpProduct(app.server, db)
}

func (app *Application) buildConfig() *Configuration {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	config := &Configuration{}
	err = viper.Unmarshal(config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func (app *Application) buildDB(config *Configuration) *gorm.DB {
	DBConfig := config.Database
	sc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBConfig.User,
		DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.DBName)

	db, err := gorm.Open(mysql.Open(sc), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}

	return db
}

func (app *Application) setUpProduct(server *echo.Echo, db *gorm.DB) {
	gateway := infrastructure.NewProductMySQLGateway(db)

	create := application.NewDefaultCreateProductUseCase(gateway)

	handler := api.NewProductApi(create)

	server.POST("/products", handler.Create)
}
