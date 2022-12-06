/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
    './app.{vue,js,ts,jsx,tsx}',
    './components/**/*.{vue,js,ts,jsx,tsx}',
    './layouts/**/*.{vue,js,ts,jsx,tsx}',
    './pages/**/*.{vue,js,ts,jsx,tsx}',
    './plugins/**/*.{js,ts}',
  ],
  theme: {
    extend: {},
    fontFamily: {
      body: [
        "'メイリオ', 'Meiryo','ＭＳ ゴシック','Hiragino Kaku Gothic ProN','ヒラギノ角ゴ ProN W3',sans-serif"
      ]
    }
  },
  plugins: [require('daisyui')],
}
