module.exports = {
  
  "transpileDependencies": [
    "vuetify"
  ],
  devServer: {
    proxy: 'http://localhost:8081'
    // {
    //   '/api': {
    //     target: 'http://localhost:8081',
    //     ws: true,
    //     changeOrigin: true
    //   }
    // }
  }
}