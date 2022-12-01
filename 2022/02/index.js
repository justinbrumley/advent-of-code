const util = require('../../util');

const run = async () => {
  const lines = await util.getInput();
  const nums = util.nums(lines);
  const groups = util.group(lines);
};

run();
