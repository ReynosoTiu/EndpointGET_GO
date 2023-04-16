package main

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Carnet string `json:"carnet"`
	Nombre string `json:"nombre"`
}

func hojaTrabajo4(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := Payload{
		Carnet: "201345126",
		Nombre: "José Luis Reynoso Tiu",
	}
	json.NewEncoder(w).Encode(p)
}
