package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Figaarillo/golerplate/internal/setup"
	"github.com/Figaarillo/golerplate/internal/shared/config"

	_ "github.com/Figaarillo/golerplate/docs" // load API Docs files (Swagger)
	_ "github.com/joho/godotenv/autoload"     // load .env file automatically
)

// @title golerplate API
// @version 1.0
// @description GOlerplate is boilerpalte for GO
// @termsOfService http://swagger.io/terms/
// @contact.name Figarillo
// @contact.email axel.leonardi.22@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @host localhost:8080
func main() {
	env, err := config.NewEnvConf(".env")
	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	PORT := env.SERVER_PORT

	db := config.InitGorm(env)
	router := config.InitRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	setup.NewCategory(router, db)
	setup.NewClient(router, db)
	setup.NewProduct(router, db)
	setup.NewSwagger(router)

	log.Printf("Server is running!🔥 Go to http://localhost:%d/\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), router)
}
