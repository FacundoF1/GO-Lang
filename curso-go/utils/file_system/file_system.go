package file_system

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)


func ReadFile(fileName string) ([][]string) {

    // open file
    fileScanner, err := os.Open(fileName)
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

    return data;
}

func WriteFile(fileName string) ([][]string, error) {
	// open file
    fileScanner, err := os.Open(fileName)
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

    return data, nil;
}

func processCSV(rc io.Reader) (ch chan []string) {
    ch = make(chan []string, 10)
    go func() {
        r := csv.NewReader(rc)
        if _, err := r.Read(); err != nil { //read header
            log.Fatal(err)
        }
        defer close(ch)
        for {
            rec, err := r.Read()
            if err != nil {
                if err == io.EOF {
                    break
                }
                log.Fatal(err)

            }
            ch <- rec
        }
    }()
    return
}