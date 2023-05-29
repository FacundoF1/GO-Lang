package electric_vehicle

import "fmt"

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
