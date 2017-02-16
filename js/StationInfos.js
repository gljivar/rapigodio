import { h, Component } from 'preact';
import StationInfo from './StationInfo';

export default class StationInfos extends Component {
  constructor() {
    super();
    this.handleChange = this.handleChange.bind(this);
  }


  handleChange() {
    this.props.handleChange();
  }

  render({ stationInfos }) {
    return (
      <div>
        {stationInfos.map((stationInfo) => ( <div>
         <StationInfo stationInfo={stationInfo} handleChange={this.handleChange}/>
         </div> ))
        }
      </div>
     );
  }
}
