package api

import (
	"bytes"
	"client/helpers"
	"client/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type UserApi struct {}

func (a *UserApi) CreateUser(name, secondName, fatherName, userName string) (userString string) {
	client := &http.Client{}
	var uArr []models.User
	user := models.User{Name: name, SecondName: secondName, FatherName: fatherName, UserName: userName}
	jsonPost, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/api/users", bytes.NewBuffer(jsonPost))
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req); if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&uArr)
	jsonUser, _ := json.Marshal(uArr)
	userString = string(jsonUser)
	return
}

func (a *UserApi) GetAllUsers() (uArr []models.User) {
	res, err := http.Get("http://localhost:8000/api/users")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&uArr)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (a *UserApi) DeleteMaxIdUser(uArr []models.User) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/api/users/" + strconv.Itoa(helpers.FindMaxId(uArr)), nil)
	if err != nil {
		log.Fatal(err)
	}
	client.Do(req)
	return
}

func (a *UserApi) UpdateMinIdUser(uArr []models.User, name, secondName, fatherName, userName string) {
	client := &http.Client{}
	user := models.User{Name: name, SecondName: secondName, FatherName: fatherName, UserName: userName}
	jsonPut, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/api/users/" + strconv.Itoa(helpers.FindMaxId(uArr)), bytes.NewBuffer(jsonPut))
	if err != nil {
		log.Fatal(err)
	}
	client.Do(req)
}


