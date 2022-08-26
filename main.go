package main

import (
	"client/api"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

func main() {
	app := http.NewServeMux()
	app.HandleFunc("/do", func(w http.ResponseWriter, r *http.Request) {
		userApi := api.UserApi{}
		uArr := userApi.GetAllUsers()
		userApi.DeleteMaxIdUser(uArr)
		userApi.UpdateMinIdUser(uArr, "Ekaterina", "Romanova", "Nicknamovna", "she")
		userApi.CreateUser("Boris", "Holov", "Nikolaevich", "ne-popadesh")
		xmlArr, _ := xml.Marshal(userApi.GetAllUsers())
		w.Write(xmlArr)
	})

	fmt.Println("Server is listening on http://localhost:8081")
	err := http.ListenAndServe(":8081", app)
	if err != nil {
		log.Fatal(err)
	}
}
