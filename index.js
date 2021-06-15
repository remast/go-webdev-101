const express = require('express')
const revealRunInTerminal = require('./addons/reveal-run-in-terminal');

const app = express()
const port = 6060

app.use(revealRunInTerminal());
app.use('/', express.static(__dirname + '/'));

app.listen(port, () => {
  console.log(`Presentations app listening at http://localhost:${port}`)
})