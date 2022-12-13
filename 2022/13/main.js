const fs = require('fs/promises');

const isCorrect = (a, b) => {
  for (let i = 0; i < Math.min(a.length, b.length); i += 1) {
    let aVal = a[i];
    let bVal = b[i];

    if (typeof aVal === 'number' && typeof bVal === 'number') {
      if (aVal < bVal) {
        // CORRECT ORDER
        return true;
      } else if (aVal > bVal) {
        // INCORRECT ORDER
        return false;
      }

      // Same number, just continue
      continue;
    }

    if (typeof aVal === 'number' && Array.isArray(bVal)) {
      // Convert a to list
      aVal = [aVal];
    }

    if (Array.isArray(aVal) && typeof bVal === 'number') {
      // Convert b to list
      bVal = [bVal];
    }

    if (Array.isArray(aVal) && Array.isArray(bVal)) {
      const isCorrectArray = isCorrect(aVal, bVal);
      if (isCorrectArray !== null) { return isCorrectArray; }
    }
  }

  // Finally, check which array ran out first
  if (a.length < b.length) {
    return true;
  } else if (a.length > b.length) {
    return false;
  }

  // Undetermined
  return null;
};

const run = async () => {
  const file = await fs.readFile('./input');
  const input = file.toString();
  const lines = input.split('\n');

  /* Part One */
  /*
  const correctIndices = [];
  let idx = 0;
  for (let i = 0; i < lines.length; i += 3, idx += 1) {
    const a = JSON.parse(lines[i]);
    const b = JSON.parse(lines[i + 1]);

    if (isCorrect(a, b)) {
      correctIndices.push(idx + 1);
    }
  }

  console.log('Sum of correct indices:', correctIndices.reduce((memo, val) => memo + val, 0));
  */

  const filteredLines = lines.filter(Boolean);

  // Add separator lines
  filteredLines.push('[[2]]');
  filteredLines.push('[[6]]');

  const sortedLines = filteredLines.sort((a, b) => (
    isCorrect(JSON.parse(a), JSON.parse(b)) ? -1 : 1
  ));

  console.log('Decoder Key:', (sortedLines.indexOf('[[2]]') + 1) * (sortedLines.indexOf('[[6]]') + 1));
};

run();
