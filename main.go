package main

import (
  "html/template"
  "net/http"
//  "errors"
  "regexp"
  //"syscall"
  //"os"
  "os/exec"
  "fmt"
  "runtime"
  "strconv"
)

type stationInfo struct {
  id int
  name string
  streamIpAddress string
}

var templates = template.Must(template.ParseFiles("index.html"))
var playValidPath = regexp.MustCompile("^/(play)/(\\d)/*(.*)$")
var quit chan bool
var radios chan bool 

func startRadio(name string, streamIpAddress string) {
  //binary, _ := exec.LookPath("mplayer")
  //args := []string{"mplayer", streamIpAddress, "-really-quiet"}
  //args := []string{streamIpAddress, "-really-quiet"}

  cmd := exec.Command("mplayer", streamIpAddress + " -really-quiet")
  err := cmd.Run()

  if err != nil {
    fmt.Println(err)
  }
  //env := os.Environ()

  //syscall.Exec(binary, args, env)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "index.html", nil)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func playHandler(w http.ResponseWriter, r *http.Request) {
  matchedPath := playValidPath.FindStringSubmatch(r.URL.Path)
  fmt.Println(matchedPath[2])
  stationId, _ := strconv.ParseInt(matchedPath[2], 0, 64)
  stationName := matchedPath[3]
  var station stationInfo
  if stationId == 2 {
    station = stationInfo{name: "otvoreni", streamIpAddress: "http://50.7.129.122:8249/;"}  
  } else if stationId == 1 {
    station = stationInfo{name: "yammat", streamIpAddress: "http://192.240.102.133:12430/stream;"}  
  } else if stationId == 3 {
    station = stationInfo{name: "cworka", streamIpAddress: "http://stream3.polskieradio.pl:8956/;.mp3"}  
  }
 
  fmt.Println("In play handler")

  go func(name string, streamIpAddress string) {
    started := false
    for {
      select {
      case <- quit:
        fmt.Println("quit 1")
        <- radios
        fmt.Println("quit 2")
        return
      default:
        if !started {
          fmt.Println("default")
          fmt.Println(radios)
          if len(radios) > 0 {
            fmt.Println("quitting " + string(len(radios)))
            quit <- true
          }
          radios <- true 

          fmt.Println("after ii default")
          startRadio(station.name, station.streamIpAddress) 
          fmt.Println("after default")
          //http.Redirect(w, r,  "/index/" + stationId + "/" + stationName, http.StatusFound)
        } 
        started = true
      }
    }
  }("t", "t")

  http.Redirect(w, r,  "/index/" + string(stationId) + "/" + stationName, http.StatusFound)
}

func main() {
  runtime.GOMAXPROCS(2)
  quit = make(chan bool, 3)
  radios = make(chan bool, 3)
  ww := make(chan bool)
  go func() {
  http.HandleFunc("/index/", indexHandler)
  http.HandleFunc("/play/", playHandler)
  http.ListenAndServe(":8080", nil)
  }()
  <- ww
}
