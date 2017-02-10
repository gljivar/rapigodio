package radio

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
)

func LoadStations(filename string) ([]StationInfo) {
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

func ChooseById(ss []StationInfo, id int) (StationInfo) {
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
 
