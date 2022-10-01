package Migration
import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sweaty/Database"
	"sweaty/Model/Sweater"
)

func Migration(c echo.Context) error {
	db := Database.ConnectToDB().Database
	Sweater.Migrate(db)
	return c.String(http.StatusOK , "migrated")

}
