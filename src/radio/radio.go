package radio

import (
  "fmt"
  //"syscall"
  "os"
  "os/exec"
)

func startRadio(name string, streamIpAddress string, commands chan *exec.Cmd) {
  //binary, _ := exec.LookPath("mplayer")
  //args := []string{"mplayer", streamIpAddress, "-really-quiet"}
  //args := []string{streamIpAddress, "-really-quiet"}

  cmd := exec.Command("mplayer", streamIpAddress + " -really-quiet")
  err := cmd.Start()

  if err != nil {
    fmt.Println(err)
  }

  commands <- cmd

  //env := os.Environ()

  //syscall.Exec(binary, args, env)
}

func Play(stationName string, streamIpAddress string, quit chan bool, radios chan bool, commands chan *exec.Cmd) {
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
        startRadio(stationName, streamIpAddress, commands)

        fmt.Println("Radio " + stationName + " started succesfully")
      }
      started = true
    }
  }
}

