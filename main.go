package main

import (
  "html/template"
  "net/http"
//  "errors"
  "regexp"
  "syscall"
  "os"
  "os/exec"
  "fmt"
)

type stationInfo struct {
  id int
  name string
  streamIpAddress string
}

var templates = template.Must(template.ParseFiles("index.html"))
var playValidPath = regexp.MustCompile("^/(play)/(\\d)/*(.*)$")
var quit chan bool

func startRadio(name string, streamIpAddress string) {
  binary, _ := exec.LookPath("mplayer")
  args := []string{"mplayer", streamIpAddress, "-really-quiet"}
 
  env := os.Environ()

  syscall.Exec(binary, args, env)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "index.html", nil)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func playHandler(w http.ResponseWriter, r *http.Request) {
  matchedPath := playValidPath.FindStringSubmatch(r.URL.Path)
  stationId := matchedPath[2]
  stationName := matchedPath[3]

  station := stationInfo{name: "otvoreni", streamIpAddress: "http://50.7.129.122:8249/;"}  
 
  fmt.Println("In play handler")

  go func(name string, streamIpAddress string) {
    started := false
    for {
      select {
      case <- quit:
        fmt.Println("quit")
        return
      default:
        if !started {
          fmt.Println("quit")
          startRadio(station.name, station.streamIpAddress) 
        } 
        started = true
      }
    }
  }("t", "t")

  http.Redirect(w, r,  "/index/" + stationId + "/" + stationName, http.StatusFound)
}

func main() {
  http.HandleFunc("/index/", indexHandler)
  http.HandleFunc("/play/", playHandler)
  http.ListenAndServe(":8080", nil)
}
