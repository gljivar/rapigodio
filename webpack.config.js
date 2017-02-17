var path = require('path');
var webpack = require('webpack');

module.exports = {
  devtool: 'source-map',
  entry: [
    //'webpack-dev-server/client?http://localhost:3000',
    'bootstrap-loader',
    './js/index'
  ],
  output: {
    path: path.join(__dirname, 'dist', 'static'),
    filename: 'bundle.js',
    publicPath: '/dist/'
  },
  module: {
    loaders: [
      { test: /\.js$/, loaders: ['babel-loader'], include: path.join(__dirname, 'js') },
      { test: /\.css$/, loader: 'style-loader!css-loader!sass-loader', include: path.join(__dirname, 'css') },
      { test:/bootstrap-sass[\/\\]assets[\/\\]javascripts[\/\\]/, loader: 'imports-loader?jQuery=jquery' },
      { test: /\.(woff2?|svg)$/, loader: 'url-loader?limit=10000' },
      { test: /\.(ttf|eot)$/, loader: 'file-loader' },
    ]
  }
};

