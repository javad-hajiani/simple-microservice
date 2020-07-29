package bootstrap

import (
	"fmt"
	"net/http"
	"users/controllers"
)

func BootApplication() {
	fmt.Println("Bootstraping the application...")
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
