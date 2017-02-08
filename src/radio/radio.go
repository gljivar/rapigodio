package radio

import (
  "fmt"
  //"syscall"
  //"os"
  "os/exec"
)

func startRadio(name string, streamIpAddress string) {
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
