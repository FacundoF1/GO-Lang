package electric_vehicle

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/utils"
)

type ElectricVehicle struct {
	Meta string `json:"meta"`
}

type ListVehicleByBrand struct {}

func ProcessInfo(w http.ResponseWriter, req *http.Request) {
	
	w.Header().Add("content-type", "application/json")
	
	// open file
    fileScanner, err := os.Open("Electric_Vehicle_Population_Data.csv")
    if err != nil {
        log.Fatal(err)
    }

	// remember to close the file at the end of the program
    defer fileScanner.Close()

	// read csv values using csv.Reader
    csvReader := csv.NewReader(fileScanner)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

	// convert records to array of structs
    response := ProcessData(data)
    out := map[string]interface{}{}
    out["electric_vehicle"] = response

    result := utils.ConvertToJsonInString(out)

    // print the array
    fmt.Fprint(w, result)

    w.WriteHeader(200);
}

func ProcessData(data [][]string) map[string]uint16 {
    listMap := map[string]uint16 {}
    for _, v := range data {
        listMap[v[6]] = listMap[v[6]] + 1;
    }
	return listMap;
}

func readData(fileName string) ([][]string, error) {

    f, err := os.Open(fileName)

    if err != nil {
        return [][]string{}, err
    }

    defer f.Close()

    r := csv.NewReader(f)

    // skip first line
    if _, err := r.Read(); err != nil {
        return [][]string{}, err
    }

    records, err := r.ReadAll()

    if err != nil {
        return [][]string{}, err
    }

    return records, nil
}

func WriteChannel(ch chan <- string, data string){
	ch <- data
}

func ProcessChannel(data string){
	chnl := make(chan string)
    go WriteChannel(chnl, data)
    fmt.Println(<-chnl)
}

func InitChannel() chan string {
	chnl := make(chan string)
	fmt.Print("Init Channel")
	return chnl;
}

func CloseChannel(chnl chan string){
    fmt.Print("Close Channel")
    close(chnl)
}

//  chnl := InitChannel();
// 	ProcessChannel("Process Inizialized")
// 	CloseChannel(chnl);