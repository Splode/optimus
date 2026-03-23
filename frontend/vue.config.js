let cssConfig = {
  loaderOptions: {
    css: {
      // Disable URL resolution in css-loader; Wails v2 embeds and serves dist/ as static files
      url: false
    }
  }
}

if (process.env.NODE_ENV === 'production') {
  cssConfig = {
    ...cssConfig,
    extract: {
      filename: '[name].css',
      chunkFilename: '[name].css'
    }
  }
}

module.exports = {
  css: cssConfig,
  configureWebpack: {
    output: {
      filename: '[name].js'
    },
    optimization: {
      splitChunks: false
    }
  },
  devServer: {
    allowedHosts: 'all'
  }
}
