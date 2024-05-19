// services/fileReader.js
const fs = require('fs');

function getMatriz(path) {
  return new Promise((resolve, reject) => {
    fs.readFile(path, 'utf8', (err, data) => {
      if (err) {
        reject(err);
      } else {
        const lines = data.trim().split('\n');
        const matrix = lines.map(line => line.trim().split(' ').map(Number));
        resolve(matrix);
      }
    });
  });
}

module.exports = {
    getMatriz,
};
