const fs = require('fs/promises');

module.exports = {
  async getInput(lines = true) {
    const file = await fs.readFile('./input');
    const input = file.toString();
    return lines ? input.split('\n') : input;
  },

  group(arr, opts = {}) {
    const {
      divider = '',
      parse = true,
      base = 10,
    } = opts;

    const out = [[]];

    for (let i = 0; i < arr.length; i += 1) {
      const val = arr[i];

      if (val === divider) {
        out.push([]);
        continue;
      }

      out[out.length - 1].push(parse ? parseInt(val, base) : val);
    }

    return out;
  },

  sum(arr, base = 10) {
    return arr.reduce((memo, val) => memo + (parseInt(val, base) || 0), 0);
  },

  avg(arr, base = 10) {
    return sum(arr) / arr.filter((val) => !isNaN(parseInt(val, base))).length;
  },
};
