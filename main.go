package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var comp = getComputer(16)

type computerState struct {
	ROM []string `json:"rom"`
	RAM []string `json:"ram"`
	PC  string   `json:"pc"`
}

func guiHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "gui.html")
}

func loadProgram(w http.ResponseWriter, r *http.Request) {
	comp.loadProgram("test.hack")
	json.NewEncoder(w).Encode(getState())
}

func doTick(w http.ResponseWriter, r *http.Request) {
	comp.tick(false)
	json.NewEncoder(w).Encode(getState())
}

func resetComputer(w http.ResponseWriter, r *http.Request) {
	comp.tick(true)
	json.NewEncoder(w).Encode(getState())
}

func setRAM(w http.ResponseWriter, r *http.Request) {
	// TODO: implement
}

func getState() computerState {
	return computerState{
		ROM: dumpROM(&comp.program),
		RAM: dumpRAM(&comp.data),
		PC:  boolToStr(comp.processor.count.read()),
	}
}

func main() {
	http.HandleFunc("/", guiHandler)
	http.HandleFunc("/load", loadProgram)
	http.HandleFunc("/tick", doTick)
	http.HandleFunc("/reset", resetComputer)
	fmt.Print("Running server")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
