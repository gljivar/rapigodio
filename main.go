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
  "log"
)

var templates = template.Must(template.ParseFiles("index.html"))
var playValidPath = regexp.MustCompile("^/(play)/(\\d)/*(.*)$")

func indexHandler(entryPoint string) func(w http.ResponseWriter, r *http.Request) {
  fn := func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, entryPoint)
  }
  return http.HandlerFunc(fn)
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
  //indexHandler(w, r)
  //http.Redirect(w, r,  "/index/" + string(stationId) + "/" + stationName, http.StatusFound)
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
  radio.Quit()
  radio.Status.NowPlaying = radio.StationInfo{}
  http.Redirect(w, r,  "/index/", http.StatusFound)
}

func radioHandler(w http.ResponseWriter, r *http.Request) {
  data, err := json.Marshal(radio.Status)
  if err != nil {
    http.Error(w, err.Error(), 400)
  }
  w.Write(data)
}

func main() {
 
  entry := flag.String("entry", "./index.html", "Entrypoint")
 // static := flag.String("static", ".", "Directory to serve static files from")
  port := flag.Int("port", 8080, "Port") 
  flag.Parse()

  radio.Initialize()

  filename := "stations.json"
  radio.Status.Stations = radio.LoadStations(filename)

  runtime.GOMAXPROCS(2)
  ww := make(chan bool)
  go func() {

  //var chttp = http.NewServeMux()

  //chttp.Handle("/", http.FileServer(http.Dir("./")))

 // http.Handle("/dist/", http.FileServer(http.Dir("./dist/")))
  //http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))
  http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))

  http.HandleFunc("/api/v1/radio", radioHandler);
  http.HandleFunc("/index/", indexHandler(*entry))
  http.HandleFunc("/play/", playHandler)
  http.HandleFunc("/stop/", stopHandler)

  portString := strconv.Itoa(*port)
  fmt.Println("Listening on port ", portString)
  log.Fatal(http.ListenAndServe(":" + portString, nil))

  }()
  <- ww
}
