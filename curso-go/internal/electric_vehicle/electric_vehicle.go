package electric_vehicle

import (
	"fmt"
	"net/http"

	"example.com/utils"
	"example.com/utils/file_system"
)

type ElectricVehicle struct {
	Meta string `json:"meta"`
}

type ListVehicleByBrand struct {}

func ProcessInfo(w http.ResponseWriter, req *http.Request) {
	
	w.Header().Add("content-type", "application/json")

	// convert records to array of structs
    data := file_system.ReadFile("Electric_Vehicle_Population_Data.csv")
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

//  chnl := InitChannel();
// 	ProcessChannel("Process Inizialized")
// 	CloseChannel(chnl);