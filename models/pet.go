package main

type Pet struct {
	Id           int      `json:Id`
	Name         string   `json:Name`
	BichoType    string   `json:BichoType` //Perro o Gato
	Gender       string   `json:Gender`    //Macho o Hembra
	Description  string   `json:Description`
	PetImage     []string `json:PetImage`
	DateAdded    string   `json:DateAdded`
	DateRescued  string   `json:DAteRescued`
	DateAdopted  string   `json:DateAdopted`
	AgeAtAdded   int      `json:AgeAtAdded`
	AgeAtAdopted int      `json:AgeAtAdopted`
	Size         string   `json:Size`
	Status       string   `json:Status`
	// Treatment []
	// PrivateInfo []
	// Responsible
	Link string `json:Link`
}

type Pets []Pet

// type Treatment struct{
// 	Id
// 	Name
// 	Description
// 	DateInit
// 	Duration
// }
