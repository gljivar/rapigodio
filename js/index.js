import { h, render } from 'preact';
import App from './App';

render(<App autoreload={true} />, document.getElementById('page'));
