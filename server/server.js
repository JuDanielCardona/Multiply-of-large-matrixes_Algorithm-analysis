const express = require('express');
const { Pool } = require('pg');
const cors = require('cors');

const app = express();
const port = 3000;

app.use(cors());
app.use(express.static('public'));

const pool = new Pool({
  user: 'admin',
  host: 'localhost',
  database: 'executiontimes_db',
  password: '12345',
  port: 5432,
});

app.get('/execution_times/golang/:test', async (req, res) => {
  const test = req.params.test;
  try {
    const result = await pool.query('SELECT algorithm, time FROM execution_times WHERE language = $1 AND test = $2', ['Golang', test]);
    res.json(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Error al obtener los datos' });
  }
});

app.get('/execution_times/python/:test', async (req, res) => {
  const test = req.params.test;
  try {
    const result = await pool.query('SELECT algorithm, time FROM execution_times WHERE language = $1 AND test = $2', ['Python', test]);
    res.json(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Error al obtener los datos' });
  }
});

app.get('/tests/golang', async (req, res) => {
  try {
    const result = await pool.query('SELECT DISTINCT test FROM execution_times WHERE language = $1', ['Golang']);
    const tests = result.rows.map(row => row.test);
    res.json(tests);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Error al obtener los tests' });
  }
});

app.get('/tests/python', async (req, res) => {
  try {
    const result = await pool.query('SELECT DISTINCT test FROM execution_times WHERE language = $1', ['Python']);
    const tests = result.rows.map(row => row.test);
    res.json(tests);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Error al obtener los tests' });
  }
});

app.get('/bar_chart_data/python/:test', async (req, res) => {
  const test = req.params.test;
  try {
    const result = await pool.query('SELECT algorithm, time FROM execution_times WHERE language = $1 AND test = $2', ['Python', test]);
    res.json(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Error al obtener los datos para el gráfico de barras' });
  }
});

app.get('/bar_chart_data/golang/:test', async (req, res) => {
  const test = req.params.test;
  try {
    const result = await pool.query('SELECT algorithm, time FROM execution_times WHERE language = $1 AND test = $2', ['Golang', test]);
    res.json(result.rows);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Error al obtener los datos para el gráfico de barras' });
  }
});

app.get('/comparison_data/:test', async (req, res) => {
  const test = req.params.test;
  try {
    const resultPython = await pool.query('SELECT algorithm, time FROM execution_times WHERE language = $1 AND test = $2', ['Python', test]);
    const resultGolang = await pool.query('SELECT algorithm, time FROM execution_times WHERE language = $1 AND test = $2', ['Golang', test]);

    // Merge data from Python and Golang
    const comparisonData = resultPython.rows.map((rowPython, index) => {
      const rowGolang = resultGolang.rows[index];
      return {
        algorithm: rowPython.algorithm,
        timePython: rowPython.time,
        timeGolang: rowGolang ? rowGolang.time : null
      };
    });

    res.json(comparisonData);
  } catch (err) {
    console.error(err);
    res.status(500).json({ error: 'Error al obtener los datos de comparación' });
  }
});

app.listen(port, () => {
  console.log(`Servidor escuchando en http://localhost:${port}`);
});
