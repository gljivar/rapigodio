import { h, render } from 'preact';
import App from './App';

render(<App autoreload={false} />, document.getElementById('page'));
