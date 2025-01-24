const path = require('path');

// https://webpack.js.org/
// https://webpack.js.org/guides/installation/
// https://webpack.js.org/concepts/
module.exports = {
  entry: './index.js',
  target: 'web',
  // https://webpack.js.org/configuration/mode/
  mode: 'development',
  mode: 'production',
  output: {
    path: path.resolve(__dirname, '.'),
    filename: 'bundle.js',
    // https://webpack.js.org/configuration/output/
    library: {
      name: 'ReactAriaComponents_1_6_0',
      type: 'window',
    },
  },
};
