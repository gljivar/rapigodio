import { h, Component } from 'preact';

let getDomain = url => url && (url.split('/')[ ~url.indexOf('://') ? 2 : 0 ]).replace(/^www\./,'') || null;
const API_ORIGIN = '/api';
const asJson = r => r.json();

export default class StationInfo extends Component {
  constructor() {
    super();
    this.onClick = this.handleClick.bind(this);
  }

  handleClick(event) {
    const {id} = event.target;
    console.log(id);

    fetch(`/play/` + id + "/Yammat")
    .then(asJson);
    //.then(radioJson => this.setState({ radioStatus: radioJson }) );
    
  }


	render({ stationInfo }) {
		return (
			<div class="stationInfo">
                       
<h3 id={stationInfo.Id} onClick={this.onClick}>
          {stationInfo.Name}
        </h3>
 
                        <a href={"/play/" + stationInfo.Id + "/" + stationInfo.Name}>{stationInfo.Name}
                          <br /> <img src={stationInfo.IconAddress} /></a>

			</div>
		);
	}
}
