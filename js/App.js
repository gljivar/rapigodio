import { h, Component } from 'preact';
import StationInfos from './StationInfos';

const API_ORIGIN = '/api';

const asJson = r => r.json();

export default class App extends Component {
  constructor() {
    super();
    this.onClick = this.stopPlaying.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.loadStationInfos = this.loadStationInfos.bind(this);
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
  
  handleChange() {
    this.loadStationInfos();
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
    return (
           <div>
             <nav class="navbar navbar-default">
               <div class="container-fluid">
                 <div class="navbar-header">
                   <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
                     <span class="sr-only">Toggle navigation</span>
                     <span class="icon-bar"></span>
                     <span class="icon-bar"></span>
                     <span class="icon-bar"></span>
                   </button>
                   <a class="navbar-brand" href="#">Rapigodio</a>
                 </div>

                 <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                   <ul class="nav navbar-nav">
                     <li><p class="navbar-text"><strong>{radioStatus.NowPlaying.Name}</strong></p></li>
                   </ul>
                   <ul class="nav navbar-nav navbar-right">
                     <li><button class="btn btn-danger pull-right" type="button" id={radioStatus.NowPlaying.Id} onClick={this.onClick}>Stop</button></li>
                   </ul>
                 </div>
               </div>
             </nav>

             <div class="container-fluid">
               <div class="row"> 
               <StationInfos stationInfos={radioStatus.Stations} handleChange={this.handleChange}/>
               </div>
               <div class="row">
               <img src={radioStatus.NowPlaying.ImageAddress || "http://memeshappen.com/media/created/Hello-good-morning-meme-3201.jpg"} class="img-responsive img-rounded center-block"/>
               </div>
            </div>
          </div>
    );
  }
}
