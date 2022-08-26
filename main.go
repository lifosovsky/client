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
		userApi.DeleteMaxIdUser()
		//userApi.UpdateMinIdUser("Ekaterina", "Romanova", "Nicknamovna", "she")
		userApi.CreateUser("Boris", "Holov", "Nikolaevich", "ne-popadesh")
		uArr := userApi.GetAllUsers()
		xmlArr, _ := xml.Marshal(uArr)
		w.Write(xmlArr)
	})

	fmt.Println("Server is listening on http://localhost:8081")
	err := http.ListenAndServe(":8081", app)
	if err != nil {
		log.Fatal(err)
	}
}
// [
// {"Id":16,"Name":"max","SecondName":"Egorov","FatherName":"Alexeeyevich","UserName":"max"},
// {"Id":15,"Name":"Vagon","SecondName":"Mops","FatherName":"Ushat","UserName":"editClient"}
// ]