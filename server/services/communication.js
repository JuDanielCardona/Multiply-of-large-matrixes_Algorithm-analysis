// routes/microservices.js
const axios = require('axios');

async function InitExecutions(test, matrixA, matrixB) {
  try {
    let respuesta;

    respuesta = await axios.post('http://localhost:8081/', {
      TestID: test,
      MatrixA: matrixA,
      MatrixB: matrixB
    });
    console.log('Python:', respuesta.data);

    respuesta = await axios.post('http://localhost:8082/', {
      TestID: test,
      MatrixA: matrixA,
      MatrixB: matrixB
    });
    console.log('Golang:', respuesta.data);

  } catch (error) {
    console.error('Error al enviar los datos');
  }
}

module.exports = {
  InitExecutions,
};
