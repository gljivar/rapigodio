import { h, Component } from 'preact';
import StationInfos from './StationInfos';

//const API_ORIGIN = 'https://hacker-news.firebaseio.com';

const asJson = r => r.json();

export default class App extends Component {
	state = { stationInfos: [] }

	loadStationInfos() {


          var infos =  [
  {
    "Id": 1,
    "Name": "Yammat",
  }
];
this.setState({stationInfos: infos});

		//fetch(`${API_ORIGIN}/v0/topstories.json`).then(asJson)
		//	.then( items => Promise.all( items.slice(0, 19).map(
		//		item => fetch(`${API_ORIGIN}/v0/item/${item}.json`).then(asJson)
		//	)) )
		//	.then( items => this.setState({ items }) );
	}

	componentDidMount() {
		this.loadStationInfos();
		if (this.props.autoreload=='true') {
			setInterval(::this.stationInfos, 4000);
		}
	}
        render({ }, { stationInfos}) {
                return (<div>
                           <StationInfos stationInfos={stationInfos} />
                           </div>);
        }

}
