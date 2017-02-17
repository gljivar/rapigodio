import { h, Component } from 'preact';

const API_ORIGIN = '/api';
const asJson = r => r.json();

export default class StationInfo extends Component {
  constructor() {
    super();
    this.onClick = this.handleClick.bind(this);
  }

  handleChange() {
    this.props.handleChange();
  } 

  handleClick(event) {
    const {id} = event.target;
    console.log(id);

    fetch(`/play/` + id + "/Yammat")
    .then(asJson)
    .then(this.handleChange());
    //.then(radioJson => this.setState({ radioStatus: radioJson }) );
  }

  render({ stationInfo }) {
    return (
      <div class="stationInfo">
        <h3 id={stationInfo.Id} onClick={this.onClick}>{stationInfo.Name}</h3>
        <img class="img-thumbnail" src={stationInfo.IconAddress} id={stationInfo.Id} onClick={this.onClick} />
      </div>
    );
  }
}
