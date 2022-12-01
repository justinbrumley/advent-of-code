const fs = require('fs/promises');

const run = async () => {
  const file = await fs.readFile('./input.txt');
  const input = file.toString();

  const lines = input.split('\n');

  const counts = [];

  for (let i = 0; i < lines.length; i += 1) {
    const val = lines[i];

    if (val.trim() === '') {
      counts.push(0);
      continue;
    }

    // Add to the total for the current elf
    counts[counts.length - 1] += parseInt(val, 10);
  }

  // Find the max value
  // console.log('Highest Calory Count:', Math.max(...counts));

  // Sort elves by count
  counts.sort((a, b) => a > b ? -1 : 1);

  console.log('Total of top three:', counts.slice(0, 3).reduce((memo, count) => memo + count, 0));
};

run();
