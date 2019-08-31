package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var comp = getComputer(16)

type computerState struct {
	ROM       []string `json:"rom"`
	RAM       []string `json:"ram"`
	PC        string   `json:"pc"`
	Aregister string   `json:"aRegister"`
	Dregister string   `json:"dRegister"`
}

type programState struct {
	ROM []string `json:"rom"`
}

func guiHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "gui.html")
}

func loadProgram(w http.ResponseWriter, r *http.Request) {
	comp.loadProgram("test.hack")
	json.NewEncoder(w).Encode(programState{ROM: dumpROM(&comp.program)})
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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal server error"))
		return
	}
	var setram map[string]interface{}
	err = json.Unmarshal(body, &setram)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal server error"))
		return
	}
	index := setram["index"].(float64)
	value := setram["bits"].(string)
	comp.data.tick(strToBool(value), intToBools(int(index), 15), true)
	json.NewEncoder(w).Encode(getState())
}

func getState() computerState {
	return computerState{
		RAM:       dumpRAM(&comp.data),
		PC:        boolToStr(comp.processor.count.read()),
		Aregister: boolToStr(comp.processor.aRegister.out),
		Dregister: boolToStr(comp.processor.dRegister.out),
	}
}

func main() {
	http.HandleFunc("/", guiHandler)
	http.HandleFunc("/load", loadProgram)
	http.HandleFunc("/tick", doTick)
	http.HandleFunc("/reset", resetComputer)
	http.HandleFunc("/set-ram", setRAM)
	fmt.Print("Running server")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
