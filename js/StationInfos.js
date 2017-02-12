import { h, Component } from 'preact';
import StationInfo from './StationInfos';

export default class StationInfos extends Component {

	render({ stationInfos }) {
		return (
			<div>
koko
				<div class="station-infos">
					{ stationInfos.map( (stationInfo, i) => (
						<StationInfo key={stationInfo.Id} rank={i} item={stationInfo} />
					)) }
				</div>
			</div>
		);
	}
}
