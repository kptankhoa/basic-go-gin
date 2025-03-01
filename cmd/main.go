package main

import (
	"fmt"
	"kptankhoa.dev/basic-go-gin/configs"
	"kptankhoa.dev/basic-go-gin/internal/app"
)

func main() {
	router, err := app.InitializeApp()
	if err != nil {
		return
	}

	err = router.Run(fmt.Sprintf(":%s", configs.PORT))
	if err != nil {
		return
	}

	fmt.Printf("Server is running on port :%s", configs.PORT)
}
