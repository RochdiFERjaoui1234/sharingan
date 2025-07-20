#!/usr/bin/env node
const fs = require('fs');

const api = `
const express = require('express');
const helmet = require('helmet');
const rateLimit = require('express-rate-limit');
const app = express();

app.use(helmet());
app.use(express.json());
app.use(rateLimit({ windowMs: 15*60*1000, max: 100 }));

let items = [];

app.get('/items', (req, res) => res.json(items));
app.post('/items', (req, res) => {
  const { name } = req.body;
  if (!name) return res.status(400).json({ error: 'Name required' });
  const item = { id: Date.now(), name };
  items.push(item);
  res.status(201).json(item);
});

const port = process.env.PORT || 3000;
app.listen(port, () => console.log(\`Secure API listening on \${port}\`));
`;

fs.writeFileSync('secure-api.js', api);
console.log('secure-api.js created');





