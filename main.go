package main

import (
  "syscall"
  "os"
  "os/exec"
)

type stationInfo struct {
  name string
  streamIpAddress string
}

func startRadio(name string, streamIpAddress string) {
  binary, _ := exec.LookPath("mplayer")
  args := []string{"mplayer", streamIpAddress}

  env := os.Environ()

  syscall.Exec(binary, args, env)
}

func main() {
  station := stationInfo{name: "otvoreni", streamIpAddress: "http://50.7.129.122:8249/;"} 

  startRadio(station.name, station.streamIpAddress)

}
