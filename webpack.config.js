const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const { VueLoaderPlugin } = require('vue-loader')


module.exports = {
  mode: "development",
  plugins: [
    new VueLoaderPlugin(),
    new HtmlWebpackPlugin({
      title: "Nuts Demo registry admin",
      template: "./web/src/index.html"
    }),
  ],
  devtool: 'inline-source-map',
  entry: {
    index: './web/src/index.js',
    print: './web/src/print.js',
  },
  output: {
    filename: '[name].bundle.js',
    path: path.resolve(__dirname, 'web/dist'),
    clean: true,
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: 'vue-loader'
      },
      {
        test: /\.css$/,
        use: [
          'vue-style-loader',
          'css-loader'
        ]
      }
    ]
  }
};