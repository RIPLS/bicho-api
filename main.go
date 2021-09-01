package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type bichoType int16

const (
	Dog bichoType = 0
	Cat           = 1
)

type gender int16

const (
	Male   gender = 0
	Female        = 1
)

type petsize int16

const (
	small      petsize = 0
	medium             = 1
	large              = 2
	extralarge         = 3
)

type petStatus int16

const (
	inTreatment      petStatus = 0
	almostReady                = 1
	readyforAdoption           = 2
	inContact                  = 3
	inAdaptation               = 4
)

type Pet struct {
	Id           int       `json:Id`
	Name         string    `json:Name`
	BichoType    bichoType `json:BichoType` //Perro o Gato
	Gender       gender    `json:Gender`    //Macho o Hembra
	Description  string    `json:Description`
	PetImage     []string  `json:PetImage`
	DateAdded    string    `json:DateAdded`
	DateRescued  string    `json:DAteRescued`
	DateAdopted  string    `json:DateAdopted`
	AgeAtAdded   int       `json:AgeAtAdded`
	AgeAtAdopted int       `json:AgeAtAdopted`
	Size         petsize   `json:Size`
	Status       petStatus `json:Status`
	Link         string    `json:Link`
	// Treatment []
	// PrivateInfo []
	// Responsible
}

type Pets []Pet

//Mocks

var bichos = Pets{
	{
		Id:           1,
		Name:         "Boniato",
		BichoType:    1,
		Gender:       0,
		Description:  "test",
		PetImage:     []string{"test1", "test2"},
		DateAdded:    "",
		DateRescued:  "",
		DateAdopted:  "",
		AgeAtAdded:   72,
		AgeAtAdopted: -1,
		Size:         3,
		Status:       0,
		Link:         "",
	},
}

func getPets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bichos)
}

func newPet(w http.ResponseWriter, r *http.Request) {
	var newPet Pet
	resp, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserte una mascota válida")
		return
	}

	json.Unmarshal(resp, &newPet)

	newPet.Id = len(bichos) + 1
	bichos = append(bichos, newPet)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPet)
}

func getPetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	petId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID inválido")
		return
	}

	for _, pet := range bichos {
		if pet.Id == petId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(pet)
		}
	}

}

func deletePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	petId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID inválido")
		return
	}

	for i, pet := range bichos {
		if pet.Id == petId {
			bichos = append(bichos[:i], bichos[i+1:]...)
			fmt.Fprintf(w, "La mascota de ID %v ha sido eliminado satisfactoriamente", petId)
		}
	}

}

func updatePet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	petId, err := strconv.Atoi(vars["id"])

	var updatedPet Pet

	if err != nil {
		fmt.Fprintf(w, "ID inválido")
		return
	}

	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte datos válidos")
		return
	}
	json.Unmarshal(resp, &updatedPet)

	for i, pet := range bichos {
		if pet.Id == petId {
			bichos = append(bichos[:i], bichos[i+1:]...)
			updatedPet.Id = petId
			bichos = append(bichos, updatedPet)
			fmt.Fprintf(w, "La mascota de ID %v ha sido actualizada satisfactoriamente", petId)
		}
	}

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a la BICHO API")

}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/bichos", getPets).Methods("GET")
	router.HandleFunc("/bichos", newPet).Methods("POST")
	router.HandleFunc("/bichos/{id}", getPetById).Methods("GET")
	router.HandleFunc("/bichos/{id}", deletePet).Methods("DELETE")
	router.HandleFunc("/bichos/{id}", updatePet).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))

}
