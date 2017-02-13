import { h, Component } from 'preact';
import StationInfo from './StationInfo';

export default class StationInfos extends Component {

  render({ stationInfos }) {
    return (
      <div>
        {stationInfos.map((stationInfo) => ( <div>
         <StationInfo stationInfo={stationInfo}/>
         </div> ))
        }
        <a href="/stop/">Stop</a>
      </div>
     );
  }
}
