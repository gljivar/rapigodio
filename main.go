package main

import (
  "html/template"
  "net/http"
//  "errors"
  "regexp"
  //"syscall"
  "os"
  "os/exec"
  "fmt"
  "runtime"
  "strconv"
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
var quit chan bool
var radios chan bool 
var commands chan *exec.Cmd
var radioStatus RadioStatus = RadioStatus{}

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

func indexHandler(w http.ResponseWriter, r *http.Request) {

  var stations []StationInfo
  stations = append(stations, StationInfo{
    Id: 1, 
    Name: "Yammat", 
    StreamIpAddress: "http://192.240.102.133:12430/stream;", 
    IconAddress: "https://thumbnailer.mixcloud.com/unsafe/128x128/profile/3/f/6/0/211e-ddbd-422b-9f89-d19ef718bb63.jpg",
    ImageAddress: "http://elelur.com/data_images/articles/happy-dogs-do-you-know-what-makes-them-really-so.jpg"})
  stations = append(stations, StationInfo{
    Id: 2, 
    Name: "Otvoreni", 
    StreamIpAddress: "http://50.7.129.122:8249/;",
    IconAddress: "https://pbs.twimg.com/profile_images/821661264954986496/Uw-QUr9u_reasonably_small.jpg",
    ImageAddress: "http://i.imgur.com/FPiTgfC.jpg"})  
  stations = append(stations, StationInfo{
    Id: 3, 
    Name: "Czworka", 
    StreamIpAddress: "http://stream3.polskieradio.pl:8956/;.mp3",
    IconAddress: "http://player.polskieradio.pl/Content/_img/pr4-logo.png",
    ImageAddress: "https://larrycraven.files.wordpress.com/2011/02/happy-dog1.jpg"})  

  if radioStatus.NowPlaying == (StationInfo{}) {
    radioStatus = RadioStatus{Stations: stations}
  }

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

  go play(station.Name, station.StreamIpAddress)

  radioStatus.NowPlaying = station 
  indexHandler(w, r)
  //http.Redirect(w, r,  "/index/" + string(stationId) + "/" + stationName, http.StatusFound)
}

func play(stationName string, streamIpAddress string) {
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


func stopHandler(w http.ResponseWriter, r *http.Request) {
  quit <- true 
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
  runtime.GOMAXPROCS(2)
  quit = make(chan bool, 3)
  radios = make(chan bool, 3)
  commands = make(chan *exec.Cmd, 3)
  ww := make(chan bool)
  go func() {
  http.HandleFunc("/index/", indexHandler)
  http.HandleFunc("/play/", playHandler)
  http.HandleFunc("/stop/", stopHandler)
  http.ListenAndServe(":8081", nil)
  }()
  <- ww
}
