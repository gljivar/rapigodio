package radio

import (
  "fmt"
  //"syscall"
  "os"
  "os/exec"
)

type StationInfo struct {
  Id int
  Name string
  StreamIpAddress string
  IconAddress string
  ImageAddress string
}

var radios chan bool
var quit chan bool
var commands chan *exec.Cmd

func Initialize() {
  radios = make(chan bool, 3)
  quit = make(chan bool, 3)
  commands = make(chan *exec.Cmd, 3)
}

func startRadio(name string, streamIpAddress string) {
  //binary, _ := exec.LookPath("mplayer")
  //args := []string{"mplayer", streamIpAddress, "-really-quiet"}
  //args := []string{streamIpAddress, "-really-quiet"}

  fmt.Println("Streaming from ", streamIpAddress)

  cmd := exec.Command("mplayer", streamIpAddress) // + " -really-quiet")
  err := cmd.Start()

  if err != nil {
    fmt.Println(err)
  }

  commands <- cmd

  //env := os.Environ()

  //syscall.Exec(binary, args, env)
}

func Play(stationName string, streamIpAddress string) {
  started := false
  for {
    select {
    case <- quit:
      fmt.Println("Quitting " + stationName)
      <- radios
      comm := <- commands
      comm.Process.Signal(os.Kill)
      fmt.Println("Quit successful for " + stationName)
      return
    default:
      if !started {
        fmt.Println("Starting " + stationName)
        if len(radios) > 0 {
          fmt.Println("Sending quit signal. Currently running radios: " + string(len(radios)))
          quit <- true
        }
        radios <- true

        fmt.Println("Starting to run radio after sending quit signals")
        startRadio(stationName, streamIpAddress)

        fmt.Println("Radio " + stationName + " started succesfully")
      }
      started = true
    }
  }
}

func Quit() {
  quit <- true
}
