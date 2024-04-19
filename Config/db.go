package Config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

var DB *gorm.DB

type DBConfig struct {
	Host            string
	Port            int
	User            string
	DBName          string
	Password        string
	MaxIdleConns    int
	MaxOpenConns    int
	MaxIdleLifeTime string
	MaxConnLifeTime string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:            os.Getenv("DB_HOST"),
		Port:            intConv(os.Getenv("DB_PORT")),
		User:            os.Getenv("DB_USER"),
		Password:        os.Getenv("DB_PASSWORD"),
		DBName:          os.Getenv("DB_NAME"),
		MaxIdleConns:    intConv(os.Getenv("DB_MAX_IDLE_CONN")),
		MaxOpenConns:    intConv(os.Getenv("DB_MAX_OPEN_CONN")),
		MaxIdleLifeTime: os.Getenv("DB_MAX_IDLE_LIFE_TIME"),
		MaxConnLifeTime: os.Getenv("DB_MAX_CONN_LIFE_TIME"),
	}
	return &dbConfig
}

// convert string to int
func intConv(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(s)
	}
	return i
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func DatabaseInit() {
	dbConfig := BuildDBConfig()
	dbURL := DbURL(dbConfig)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{Logger: dbLogger})

	/* Database Connection Pooling */

	sqlDB, err := DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)

	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	maxIdleLifeTime, _ := time.ParseDuration(dbConfig.MaxIdleLifeTime)
	sqlDB.SetConnMaxIdleTime(time.Minute * maxIdleLifeTime)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	maxConnLifeTime, _ := time.ParseDuration(dbConfig.MaxConnLifeTime)
	sqlDB.SetConnMaxLifetime(time.Minute * maxConnLifeTime)

	/* Database Connection Pooling */

	if err != nil {
		panic("Failed to connect to database")
	}
}

func GetDB() *gorm.DB {
	return DB
}
