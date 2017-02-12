import { h, Component } from 'preact';
import StationInfos from './StationInfos';

const API_ORIGIN = 'https://hacker-news.firebaseio.com';

const asJson = r => r.json();

export default class App extends Component {
	state = { items: [] };

	loadItems() {

          items =  [
  {
    "Id": 1,
    "Name": "Yammat",
    "StreamIpAddress": "http://192.240.102.133:12430/stream;",
    "IconAddress": "https://thumbnailer.mixcloud.com/unsafe/128x128/profile/3/f/6/0/211e-ddbd-422b-9f89-d19ef718bb63.jpg",
    "ImageAddress": "http://elelur.com/data_images/articles/happy-dogs-do-you-know-what-makes-them-really-so.jpg"
  },]

		//fetch(`${API_ORIGIN}/v0/topstories.json`).then(asJson)
		//	.then( items => Promise.all( items.slice(0, 19).map(
		//		item => fetch(`${API_ORIGIN}/v0/item/${item}.json`).then(asJson)
		//	)) )
		//	.then( items => this.setState({ items }) );
	}

	componentDidMount() {
		this.loadItems();
		if (this.props.autoreload=='true') {
			setInterval(::this.loadItems, 4000);
		}
	}

	render({ }, { items }) {
		return (<div>
				<StationInfos items={items} />
			   </div>);
	}
}
