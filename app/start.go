package app

// Use this file to have anything start just before the http webserver is
// started. Anything within 'func Start()' will fire just before-hand. This
// cound include migrations, last min config etc...
import (
	"fmt"

	"leyra/app/models"
)

// Before fires just before the application's web server is started.
func Before() {
	fmt.Printf("Starting your application...\n")

	S.DB.AutoMigrate(&model.Post{})
}

// After fires just after the application's web server has been started.
func After() {

}
