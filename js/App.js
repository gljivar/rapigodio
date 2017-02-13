import { h, Component } from 'preact';
import StationInfos from './StationInfos';

const API_ORIGIN = '/api';

const asJson = r => r.json();

export default class App extends Component {
  constructor() {
    super();
    this.onClick = this.stopPlaying.bind(this);
  }

  state = { radioStatus: {
              NowPlaying: "",
              Stations: []
            }
          }

  stopPlaying(event) {
    fetch(`/stop/`)
    .then(asJson)
    .then(radioJson => this.setState({ radioStatus: radioJson }) );
  }

  loadStationInfos() {
    fetch(`${API_ORIGIN}/v1/radio`)
    .then(asJson)
    .then(radioJson => this.setState({ radioStatus: radioJson }) );
  }

  componentDidMount() {
    this.loadStationInfos();
    if (this.props.autoreload=='true') {
      setInterval(::this.loadStationInfos, 4000);
    }
  }

  render({ }, {radioStatus}) {
    return (<div>
               <div>Now playing: {radioStatus.NowPlaying.Name}</div>
               <h3 id={radioStatus.NowPlaying.Id} onClick={this.onClick}>
                 Stop 
               </h3>
   
               <StationInfos stationInfos={radioStatus.Stations} />

               <img src={radioStatus.NowPlaying.ImageAddress || "http://memeshappen.com/media/created/Hello-good-morning-meme-3201.jpg"} />
            </div>

);
    }
}
