// app.js
const { getMatriz } = require('./services/matrixReader');
const { InitExecutions } = require('./services/communication');

const testTag = 'Test2';
const fileA = 'sources/matrizA.txt';
const fileB = 'sources/matrizB.txt';

Promise.all([getMatriz(fileA), getMatriz(fileB)])
  .then(([matrixA, matrixB]) => {
    InitExecutions(testTag, matrixA, matrixB);
  })
  .catch(error => {
    console.error('Error al leer los archivos:', error);
  });
