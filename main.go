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
  "encoding/json"
  "io/ioutil"
)

type StationInfo struct {
  Id int
  Name string
  StreamIpAddress string
  IconAddress string
  ImageAddress string
}

type RadioStatus struct {
  NowPlaying StationInfo
  Stations []StationInfo
} 

var templates = template.Must(template.ParseFiles("index.html"))
var playValidPath = regexp.MustCompile("^/(play)/(\\d)/*(.*)$")
var radioStatus RadioStatus = RadioStatus{}

func loadStations(filename string) ([]StationInfo) {
  fileContent, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err) 
  }
  var stations []StationInfo 
  if err := json.Unmarshal(fileContent, &stations); err != nil {
    panic(err)
  }
  fmt.Println(stations)

  return stations 
}

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
  station := chooseById(radioStatus.Stations, stationId)
 
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
  radioStatus.NowPlaying = StationInfo{}
  http.Redirect(w, r,  "/index/", http.StatusFound)
}

func chooseById(ss []StationInfo, id int) (StationInfo) {
  allSatisfyingCondition := choose(ss, func(s StationInfo) bool { return s.Id == id })
  return allSatisfyingCondition[0]
}

func choose(ss []StationInfo, test func(StationInfo) bool) (ret []StationInfo) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}

func main() {
  port := flag.Int("p", 8080, "Port") 
  flag.Parse()

  filename := "stations.json"
  radioStatus.Stations = loadStations(filename)

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
