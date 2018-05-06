package lib

import (
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	config "github.com/sermilrod/wailingbot/lib/configuration"
	"github.com/sermilrod/wailingbot/lib/migrations"
	"github.com/sermilrod/wailingbot/lib/persistence"
	"github.com/sermilrod/wailingbot/lib/quotes/handlers"
)

// Start and configure the webserver
func Start() {
	cfg := config.Configuration{}
	log.SetLevel(log.INFO)
	log.SetOutput(os.Stdout)
	err := config.Parse(&cfg)
	if err != nil {
		log.Errorf("Unable to parse configuration: %s", err.Error())
	}

	dbConn, err := pg.Client(&cfg)
	if err != nil {
		log.Errorf("Unable to establish database connection: %s", err)
	} else {
		if err := dbConn.Ping(); err != nil {
			dbConn.Close()
			log.Errorf("postgresql: Could not establish a good connection: %v", err)
		}
		defer dbConn.Close()

		// Run migrations
		migrations.Run(dbConn)
		log.Printf("Migrations completed successfully")

		// Echo instance
		api := echo.New()
		api.HideBanner = true
		api.Server.Addr = ":" + cfg.Port

		// Middleware
		api.Use(middleware.Logger())
		api.Use(middleware.Recover())

		// Routes
		api.POST("/quote", quote.QuoteHandler(&cfg, dbConn))
		api.POST("/random", quote.RandomHandler(&cfg, dbConn))

		// Start server and configure gracehttp to gracefully shut it down
		api.Logger.SetLevel(log.INFO)
		api.Logger.Fatal(gracehttp.Serve(api.Server))
	}
}
