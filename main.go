package main

import (
  "html/template"
  "net/http"
  "regexp"
  "fmt"
  "runtime"
  "strconv"
  "radio" 
  "flag"
)

type RadioStatus struct {
  NowPlaying radio.StationInfo
  Stations []radio.StationInfo
} 

var templates = template.Must(template.ParseFiles("index.html"))
var playValidPath = regexp.MustCompile("^/(play)/(\\d)/*(.*)$")
var radioStatus RadioStatus = RadioStatus{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "index.html", radioStatus)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func playHandler(w http.ResponseWriter, r *http.Request) {
  matchedPath := playValidPath.FindStringSubmatch(r.URL.Path)
  fmt.Println(matchedPath[2])
  stationId, _ := strconv.Atoi(matchedPath[2])
  station := radio.ChooseById(radioStatus.Stations, stationId)
 
  fmt.Println("In play handler of " + station.Name)

  if (radioStatus.NowPlaying.Id == station.Id) {
    http.Redirect(w, r,  "/index/" + string(station.Id) + "/" + station.Name, http.StatusFound)
    return
  }

  go radio.Play(station.Name, station.StreamIpAddress)

  radioStatus.NowPlaying = station 
  indexHandler(w, r)
  //http.Redirect(w, r,  "/index/" + string(stationId) + "/" + stationName, http.StatusFound)
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
  radio.Quit()
  radioStatus.NowPlaying = radio.StationInfo{}
  http.Redirect(w, r,  "/index/", http.StatusFound)
}

func main() {
  port := flag.Int("p", 8080, "Port") 
  flag.Parse()

  filename := "stations.json"
  radioStatus.Stations = radio.LoadStations(filename)

  runtime.GOMAXPROCS(2)
  radio.Initialize()
  ww := make(chan bool)
  go func() {
  http.HandleFunc("/index/", indexHandler)
  http.HandleFunc("/play/", playHandler)
  http.HandleFunc("/stop/", stopHandler)

  portString := strconv.Itoa(*port)
  fmt.Println("Listening on port ", portString)
  http.ListenAndServe(":" + portString, nil)

  }()
  <- ww
}
