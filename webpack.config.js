const HtmlWebPackPlugin = require("html-webpack-plugin");
const publicPath = this.mode === "production" ? "/dist" : "/";

module.exports = {
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
        },
      },
      {
        test: /\.html$/,
        use: [
          {
            loader: "html-loader",
          },
        ],
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"],
      },
    ],
  },
  output: {
    path: __dirname + "/dist",
    filename: "main.js",
    publicPath: process.env.NODE_ENV === "production" ? "/dist" : "/",
  },
  plugins: [
    new HtmlWebPackPlugin({
      template: "./src/index.ejs",
      filename: "./index.html",
      templateParameters: {
        host: process.env.NODE_ENV === "production" ? "" : "http://biturl.top",
      },
    }),
  ],
  devServer: {
    historyApiFallback: true,
  },
};
