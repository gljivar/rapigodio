import { h, Component } from 'preact';

let getDomain = url => url && (url.split('/')[ ~url.indexOf('://') ? 2 : 0 ]).replace(/^www\./,'') || null;

export default class StationInfo extends Component {
	render({ stationInfo }) {
		return (
			<div class="stationInfo">
po
				<div class="vitals">
					<div class="name">{stationInfo.Name}</div>
				</div>
			</div>
		);
	}
}
