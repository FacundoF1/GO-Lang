package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"example.com/electric_vehicle"
	"github.com/playsistemico/curso-go/internal/examples"
	"github.com/playsistemico/curso-go/internal/tuiter"
)

var reader = bufio.NewReader(os.Stdin)
var mem1_before, mem2_before, mem1_after, mem2_after runtime.MemStats

func main() {
	http.HandleFunc("/ping", tracer(pingPong))
	http.HandleFunc("/tuit", tuiter.ReceiveTuit)
	http.HandleFunc("/timeline", tuiter.Timeline)
	http.HandleFunc("/vehicles/electrics/info", electric_vehicle.ProcessInfo)
	http.HandleFunc("/monitor", MonitorRuntime)
	http.ListenAndServe(":8080", nil)
	
}

func pingPong(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "pong")
}

func tracer(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		req.Header.Add("trace", "1")
		handler(res, req)
	}
}

func ShowExamples() {
	fmt.Println("Please choose which example to run:")
	fmt.Println("0: Hello world")
	fmt.Println("1: Values")
	fmt.Println("2: Variables")
	fmt.Println("3: Control Structures")
	fmt.Println("4: Arrays")
	fmt.Println("5: Maps")
	fmt.Println("6: Functions")
	fmt.Println("7: Pointers")
	fmt.Println("8: Interfaces")
	fmt.Println("9: Errors")
	fmt.Println("10: Routines")
	fmt.Println("11: Defer")

	option, err := ReadOption(reader)
	if err != nil {
		fmt.Println("Error, invalid option")
		return
	}

	switch option {
	case 0:
		examples.TryHelloWorld()
	case 1:
		examples.TryValues()
	case 2:
		examples.TryVariables()
	case 3:
		examples.TryControlStructures()
	case 4:
		examples.TryArrays()
	case 5:
		examples.TryMaps()
	case 6:
		examples.TryFunctions()
	case 7:
		examples.TryPointers()
	case 8:
		examples.TryInterfaces()
	case 9:
		examples.TryErrors()
	case 10:
		examples.TryRoutines()
	case 11:
		examples.TryDefer()
	default:
		fmt.Println("Error, invalid option")
		return
	}
}

func ReadOption(reader *bufio.Reader) (int, error) {
	var s string
	if s, _ = reader.ReadString('\n'); true {
		s = strings.Trim(s, " \n")
	}
	option, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("The provided option was invalid, exiting...")
		return -1, err
	}

	return option, nil
}

func MonitorRuntime(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("content-type", "application/json")
    m := &runtime.MemStats{}
    // f, err := os.Create(fmt.Sprintf("mmem.csv"))
    // if err != nil {
    //     panic(err)
    // }
	
    // f.WriteString("Allocated;Total Allocated; System Memory;Num Gc;Heap Allocated;Heap System;Heap Objects;Heap Released;\n")
    
	// for { 
    //     f.WriteString(fmt.Sprintf("%d;%d;%d;%d;%d;%d;%d;%d;\n", m.Alloc, m.TotalAlloc, m.Sys, m.NumGC, m.HeapAlloc, m.HeapSys, m.HeapObjects, m.HeapReleased))
    //     time.Sleep(5 * time.Second)
    // }
	runtime.ReadMemStats(m)

	out := map[string]interface{}{}
	delete(out, "Allocated")
	delete(out, "Total Allocated;")
	delete(out, "System Memory")
	delete(out, "Num Gc")
	delete(out, "Heap Allocated")
	delete(out, "Heap System")
	delete(out, "Heap Objects")
	delete(out, "Heap Release")
	
	out["Allocated"] = m.TotalAlloc
	out["Total Allocated"] = m.Alloc 
	out["System Memory"] = m.Sys
	out["Num Gc"] = m.NumGC
	out["Heap Allocated"] = m.HeapAlloc 
	out["Heap System"] =  m.HeapSys
	out["Heap Objects"] = m.HeapObjects 
	out["Heap Release"] = m.HeapReleased

	response, _ := json.MarshalIndent(out, "", "\t")
	fmt.Fprint(w, string(response))
	w.WriteHeader(200);
}
