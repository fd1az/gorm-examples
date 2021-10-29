package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/fd1az/gorm-examples/model"
	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbConfig struct {
	Host      string
	Port      int
	UserName  string
	Password  string
	Name      string
	debugMode bool
}

func main() {
	var cfg DbConfig
	//Seteamos la config de la db
	flag.IntVar(&cfg.Port, "Port", 5432, "Puerto de la base de Postgres")
	flag.StringVar(&cfg.Host, "Host", "localhost", "Host de la base de Postgres")
	flag.StringVar(&cfg.UserName, "UserName", "postgres", "Nombre de usuario para la conexión a Postgres")
	flag.StringVar(&cfg.Password, "Password", "postgres", "Contraseña para la conexión a Postgres")
	flag.StringVar(&cfg.Name, "Name", "gorm_examples", "Nombre la base de Postgres")
	flag.BoolVar(&cfg.debugMode, "debugMode", true, "debug mode de la base de Postgres")
	flag.Parse()

	db := createGomDB(&cfg)
	defer closeGormDBConnection(db)

	//Read Users
	users := GetAllUsers(db, []model.User{})

	for _, u := range users {

		fmt.Println("ID: ", u.ID)
		fmt.Println("Name: ", u.Name)
		fmt.Println("Email", u.Email)
		fmt.Println("CreatedAt", u.CreatedAt)
		fmt.Println("UpdatedAt", u.UpdatedAt)
		fmt.Println("*------------------------*")
	}

	//Create Product
	price, err := decimal.NewFromString("99.99")
	if err != nil {
		panic(err)
	}
	product1 := model.Product{
		Title:       "producto1",
		Description: "el mejor producto del mundo",
		Price:       price,
	}

	err = CreateProduct(db, &product1)

	if err != nil {
		panic(err)
	}
	fmt.Println("ID: ", product1.ID)
	fmt.Println("Title: ", product1.Title)
	fmt.Println("Description", product1.Description)
	fmt.Println("Price", product1.Price)
	fmt.Println("Email", product1.Description)
	fmt.Println("CreatedAt", product1.CreatedAt)
	fmt.Println("UpdatedAt", product1.UpdatedAt)
	fmt.Println("*------------------------*")
}

// createGomDB configuración de acceso a datos y GORM
func createGomDB(cfg *DbConfig) *gorm.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.UserName,
		cfg.Name,
		cfg.Password,
	)

	dbPool := &sql.DB{}

	gormConfig := &gorm.Config{
		Logger:      logger.Discard,
		PrepareStmt: true,
		ConnPool:    dbPool,
	}

	if cfg.debugMode {
		newLogger := logger.New(
			log.Default(), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		)
		gormConfig.Logger = newLogger
	}

	db, err := gorm.Open(postgres.Open(connStr), gormConfig)

	if err != nil {
		log.Fatalf("error trying to connect to DB: %v", err)
	}

	return db
}

// closeGormDBConnection cierra conexiones a DB relacional

func closeGormDBConnection(db *gorm.DB) {
	stmtManger, ok := db.ConnPool.(*gorm.PreparedStmtDB)

	if ok {
		for _, stmt := range stmtManger.Stmts {
			stmt.Close() // close the prepared statement
		}
	}

	dbLocal, err := db.DB()
	if err == nil {
		dbLocal.Close() //CloseDB
	}
}

//CRUD ops

//GetAllUsers al ejecutar esta función podemos ver que el Query que se ejecuta en la DB, tiene la WHERE validando que delete_at sea NULL, esto es porque por defecto GORM utiliza soft delete cuando la delete_at existe en la tabla. Niceeee
func GetAllUsers(db *gorm.DB, users []model.User) []model.User {

	rs := db.Find(&users)

	fmt.Println("RowsAffected", rs.RowsAffected) // returns count of records found
	fmt.Println("Error", rs.Error)               // returns error or nil

	return users
}

func CreateProduct(db *gorm.DB, p *model.Product) error {
	return db.Create(&p).Error
}
