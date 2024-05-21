const path = require('path');

module.exports = {
  entry: './public/js/vendor/index.js',
  output: {
    path: path.resolve(__dirname, 'public', 'js', 'vendor'),
    filename: 'bundle.js',
  },
  module: {
    rules: [
      {
        test: /\.(css|sass|scss)$/,
        use: [
          {
            loader: 'style-loader'
          },
          {
            loader: 'css-loader'
          },
          {
            loader: 'sass-loader'
          }
        ]
      }
    ]
  },
  devtool: 'source-map',
};