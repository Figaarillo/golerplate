package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm(env *EnvVars) *gorm.DB {
	DB_HOST := env.DATABASE_HOST
	DB_PORT := env.DATABASE_PORT
	DB_USER := env.DATABASE_USER
	DB_PASS := env.DATABASE_PASS
	DB_NAME := env.DATABASE_NAME
	DB_MAX_CONNS := env.DATABASE_MAX_CONNS
	DB_MAX_IDLE_CONNS := env.DATABASE_MAX_IDLE_CONNS
	DB_MAX_CONN_LIFETIME := env.DATABASE_MAX_CONN_LIFETIME

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT)

	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		time.Sleep(time.Duration(i) * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(DB_MAX_IDLE_CONNS)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(DB_MAX_CONNS)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(DB_MAX_CONN_LIFETIME)

	log.Println("Successfully connected to database 🤟")

	return db
}
