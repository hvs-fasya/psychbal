module.exports = {
  devServer: {
    hot: true,
    port: 8082, // CHANGE YOUR PORT HERE!
    https: false,
    inline: true,
    proxy: {
      '/api/v1/':{
        target: 'https://localhost:8080'
      }
    }
  },
  pluginOptions: {
    quasar: {
      theme: 'mat',
      importAll: true
    }
  },
  transpileDependencies: [
    /[\\\/]node_modules[\\\/]quasar-framework[\\\/]/
  ]
}
