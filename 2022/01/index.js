const util = require('./util');

const run = async () => {
  const lines = await util.getInput();
  const groups = util.group(lines);

  const sums = groups
    .map((values) => util.sum(values))
    .sort((a, b) => a > b ? -1 : 1);

  console.log(util.sum(sums.slice(0, 3)));
};

run();
