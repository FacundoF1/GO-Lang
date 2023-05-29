package tuiter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

/*
{
	content: "contenido del tuit",
	user_id: "user_id"
}
*/

type Tuit struct {
	Content string `json:"content"`
	UserId  string `json:"user_id"`
}

func ReceiveTuit(w http.ResponseWriter, req *http.Request) {
	var tuit Tuit
	json.NewDecoder(req.Body).Decode(&tuit)
	writeToFile(tuit)
	ProcessChannel(tuit)
	w.WriteHeader(201)
}

func Timeline(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
	f, _ := os.Open("/tmp/tuits")
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fmt.Fprintf(w, "[")
	for fileScanner.Scan() {
		fmt.Fprintf(w, "%s,", fileScanner.Text())
	}
	fmt.Fprintf(w, "]")
}

func writeToFile(tuit Tuit) {
	f, _ := os.OpenFile("/tmp/tuits", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	json.NewEncoder(f).Encode(tuit)
	f.Sync()
	f.Close()
}

func WriteChannel(ch chan <- Tuit, data Tuit){
	ch <- data
}

func ProcessChannel(data Tuit){
	chnl := make(chan Tuit)
    go WriteChannel(chnl, data)
    fmt.Println(<-chnl)
}