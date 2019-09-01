package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

type softwareLibrary struct {
	Programs []string `json:"programs"`
}

func guiHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "gui.html")
}

func getPrograms(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("software")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 Internal server error"))
		return
	}
	programs := make([]string, 0)
	for _, f := range files {
		programs = append(programs, f.Name())
	}
	json.NewEncoder(w).Encode(softwareLibrary{Programs: programs})
}

func loadProgram(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["program"]
	if len(keys[0]) < 1 {
		log.Println("URL query 'program' is missing")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request"))
		return
	}
	comp = getComputer(16)
	comp.loadProgram("software/" + keys[0])
	json.NewEncoder(w).Encode(programState{ROM: dumpROM(&comp.program)})
}

func doTick(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()["ticks"]
	if len(keys) == 0 || len(keys[0]) == 0 {
		comp.tick(false)
	} else {
		amount, err := strconv.Atoi(keys[0])
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 Bad Request"))
			return
		}
		for i := 0; i < amount; i++ {
			comp.tick(false)
		}
	}

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
	http.HandleFunc("/get-programs", getPrograms)
	fmt.Print("Running server")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
