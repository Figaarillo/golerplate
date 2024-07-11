package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type EnvVars struct {
	JWT_SECRET_KEY             string        `validate:"required"`
	TEST_DATABASE_NAME         string        `validate:"required"`
	DATABASE_NAME              string        `validate:"required"`
	TEST_DATABASE_PASS         string        `validate:"required"`
	DATABASE_USER              string        `validate:"required"`
	TEST_DATABASE_USER         string        `validate:"required"`
	DATABASE_HOST              string        `validate:"required"`
	DATABASE_PASS              string        `validate:"required"`
	SERVER_PORT                int           `validate:"required,gte=1,lte=65535"`
	DATABASE_PORT              int           `validate:"required,gte=1,lte=65535"`
	JWT_EXPIRES_IN             int           `validate:"required,gte=1"`
	DATABASE_MAX_CONN_LIFETIME time.Duration `validate:"required"`
	TEST_DATABASE_PORT         int           `validate:"required,gte=1,lte=65535"`
	SERVER_READ_TIMEOUT        int           `validate:"required,gte=1"`
	DATABASE_MAX_IDLE_CONNS    int           `validate:"required,gte=1"`
	DATABASE_MAX_CONNS         int           `validate:"required,gte=1"`
}

func NewEnvConf(envPath string) (*EnvVars, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return &EnvVars{}, fmt.Errorf("error loading .env: %w", err)
	}

	env := &EnvVars{
		SERVER_PORT:                atoi(os.Getenv("SERVER_PORT")),
		SERVER_READ_TIMEOUT:        atoi(os.Getenv("SERVER_READ_TIMEOUT")),
		DATABASE_HOST:              os.Getenv("DATABASE_HOST"),
		DATABASE_PORT:              atoi(os.Getenv("DATABASE_PORT")),
		DATABASE_USER:              os.Getenv("DATABASE_USER"),
		DATABASE_PASS:              os.Getenv("DATABASE_PASS"),
		DATABASE_NAME:              os.Getenv("DATABASE_NAME"),
		DATABASE_MAX_CONNS:         atoi(os.Getenv("DATABASE_MAX_CONNS")),
		DATABASE_MAX_IDLE_CONNS:    atoi(os.Getenv("DATABASE_MAX_IDLE_CONNS")),
		DATABASE_MAX_CONN_LIFETIME: time.Duration(atoi(os.Getenv("DATABASE_MAX_CONN_LIFETIME"))),
		JWT_SECRET_KEY:             os.Getenv("JWT_SECRET_KEY"),
		JWT_EXPIRES_IN:             atoi(os.Getenv("JWT_EXPIRES_IN")),
		TEST_DATABASE_PORT:         atoi(os.Getenv("TEST_DATABASE_PORT")),
		TEST_DATABASE_USER:         os.Getenv("TEST_DATABASE_USER"),
		TEST_DATABASE_PASS:         os.Getenv("TEST_DATABASE_PASS"),
		TEST_DATABASE_NAME:         os.Getenv("TEST_DATABASE_NAME"),
	}

	err = validateEnvVars(env)
	if err != nil {
		return &EnvVars{}, err
	}

	return env, nil
}

func validateEnvVars(env *EnvVars) error {
	validate := validator.New()
	if err := validate.Struct(env); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(*env).FieldByName(err.StructField())
			return fmt.Errorf("cannot bind file: %s, error: %s", field.Name, err.Tag())
		}
	}
	return nil
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
