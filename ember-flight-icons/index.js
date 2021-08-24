'use strict';

module.exports = {
  name: require('./package').name,

  // if your addon uses dynamic import
  options: {
    babel: {
      plugins: [require.resolve('ember-auto-import/babel-plugin')],
    },
  },

  autoImport: {
    webpack: {
      // tbd
      module: {
        rules: [
          {
            test: '/\.svg/',
            type: 'asset/resource'
            type: 'asset/inline'
            type: 'asset/source'
          }
        ]
      }
    }
  }
};
