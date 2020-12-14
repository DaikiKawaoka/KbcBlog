module.exports = {
  presets: [
    '@vue/cli-plugin-babel/preset'
  ],
  pages: {
    index: {
      entry: 'src/main.js', // 必須パラメータ
      title: 'ページタイトル',
    }
  }
}
