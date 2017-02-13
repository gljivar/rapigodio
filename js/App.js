import { h, Component } from 'preact';
import StationInfos from './StationInfos';

const API_ORIGIN = '/api';

const asJson = r => r.json();

export default class App extends Component {
  state = { radioStatus: {
              NowPlaying: "",
              Stations: []
            }
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
               <StationInfos stationInfos={radioStatus.Stations} />
            </div>);
    }
}
