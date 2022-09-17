package Migration
import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"sweaty/Database"
)

func Migration(c echo.Context) error {
	db := Database.ConnectToDB().Database
	err := db.AutoMigrate(&Sweater{})
	if err != nil {
		log.Fatal(err)
	}
	return c.String(http.StatusOK , "migrated")

}
