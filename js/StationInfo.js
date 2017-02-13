import { h, Component } from 'preact';

let getDomain = url => url && (url.split('/')[ ~url.indexOf('://') ? 2 : 0 ]).replace(/^www\./,'') || null;

export default class StationInfo extends Component {
	render({ stationInfo }) {
		return (
			<div class="stationInfo">
					<div class="name">{stationInfo.Name}</div>
                        <a href={"/play/" + stationInfo.Id + "/" + stationInfo.Name}>{stationInfo.Name}
                          <br /> <img src={stationInfo.IconAddress} /></a>
                       <p><a href="/play/{{stationInfo.Id}}/{{stationInfo.Name}}">{stationInfo.Name} <br /> <img src="{{stationInfo.IconAddress}}" /></a></p>

			</div>
		);
	}
}
