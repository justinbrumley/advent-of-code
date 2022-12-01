const fs = require('fs/promises');
const util = require('../util');

const run = async () => {
  const file = await fs.readFile('./input');
  const input = file.toString();
  const lines = input.split('\n');
  const nums = lines.map((v) => parseInt(v));
};

run();
