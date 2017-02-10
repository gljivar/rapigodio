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

var templates = template.Must(template.ParseFiles("index.html"))
var playValidPath = regexp.MustCompile("^/(play)/(\\d)/*(.*)$")

func indexHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "index.html", radio.Status)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func playHandler(w http.ResponseWriter, r *http.Request) {
  matchedPath := playValidPath.FindStringSubmatch(r.URL.Path)
  fmt.Println(matchedPath[2])
  stationId, _ := strconv.Atoi(matchedPath[2])
  station := radio.ChooseById(radio.Status.Stations, stationId)
 
  fmt.Println("In play handler of " + station.Name)

  if (radio.Status.NowPlaying.Id == station.Id) {
    http.Redirect(w, r,  "/index/" + string(station.Id) + "/" + station.Name, http.StatusFound)
    return
  }

  go radio.Play(station.Name, station.StreamIpAddress)

  radio.Status.NowPlaying = station 
  indexHandler(w, r)
  //http.Redirect(w, r,  "/index/" + string(stationId) + "/" + stationName, http.StatusFound)
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
  radio.Quit()
  radio.Status.NowPlaying = radio.StationInfo{}
  http.Redirect(w, r,  "/index/", http.StatusFound)
}

func main() {
  port := flag.Int("p", 8080, "Port") 
  flag.Parse()

  radio.Initialize()

  filename := "stations.json"
  radio.Status.Stations = radio.LoadStations(filename)

  runtime.GOMAXPROCS(2)
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
