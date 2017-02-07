package main

import (
  "syscall"
  "os"
  "os/exec"
  "fmt"
  "bufio"
)

type stationInfo struct {
  name string
  streamIpAddress string
}

func startRadio(name string, streamIpAddress string) {
  //binary, _ := exec.LookPath("mplayer")
  //args := []string{"mplayer", streamIpAddress, "-really-quiet"}
 
  binary, _ := exec.LookPath("ls") 
  args := []string{"ls"}
  env := os.Environ()

  syscall.Exec(binary, args, env)
}

func main() {
  station := stationInfo{name: "otvoreni", streamIpAddress: "http://50.7.129.122:8249/;"} 

  quit := make(chan bool)
  
  go func(name string, streamIpAddress string) {
    started := false
    for {
      select {
      case <- quit:
        return
      default:
        if !started {
          startRadio(station.name, station.streamIpAddress) 
        } 
        started = true
      }
    }
  }("t", "t")
  //go startRadio(station.name, station.streamIpAddress)
 
  scanner := bufio.NewScanner(os.Stdin)
    var text string
    for text != "q" {  // break the loop if text == "q"
        fmt.Print("Enter your text: ")
        scanner.Scan()
        text = scanner.Text()

            fmt.Println("Your text was: ", text)
        if text == "s" {
            quit <- true
            fmt.Println("Your text was: ", text)
        }
    }
}
