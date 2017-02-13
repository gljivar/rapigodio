import { h, Component } from 'preact';
import StationInfo from './StationInfo';

export default class StationInfos extends Component {

	render({ stationInfos }) {
		return (
			<div>
koko
{stationInfos.map((stationInfo) => ( <div>
<StationInfo stationInfo={stationInfo}/>
X</div> ))}

			</div>
		);
	}
}
