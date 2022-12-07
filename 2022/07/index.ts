const fs = require('fs/promises');

const TOTAL_SPACE = 70e6;
const NEEDED_SPACE = 30e6;

const get = (obj, path) => {
  if (path.length === 0) {
    return obj;
  }

  return get(obj[path[0]], path.slice(1));
};

const buildDirectory = (lines) => {
  // Track where we currently are
  let currentDir = [];

  // Map of directories and files w/sizes
  const directory = {};

  for (let i = 0; i < lines.length; i += 1) {
    const line = lines[i];

    if (line.indexOf('$') === 0) {
      // Process command
      const parts = line.split(' ');
      const [command] = parts.slice(1);

      switch (command) {
        case 'cd': {
          const [dest] = parts.slice(2);

          if (dest === '..') {
            currentDir.pop();
          } else if (dest === '/') {
            // Return to root
            currentDir = [];
          } else {
            currentDir.push(dest);
          }

          break;
        }

        case 'ls': {
          // Loop over next results until you see another command
          for (let j = i + 1; j < lines.length; j += 1) {
            const [info, name] = lines[j].split(' ');

            if (info === '$') {
              // Done processing files, adjust i and break
              i = j - 1;
              break;
            }

            const map = get(directory, currentDir);

            if (info === 'dir') {
              // Push dir to map if not exists already
              map[name] = map[name] || {};
            } else {
              // Add file to current directory with file size
              map[name] = parseInt(info, 10);
            }
          }

          break;
        }
      }
    }
  }

  return directory;
};

// Recursively get size of given directory
const getSize = (directory) => {
  return Object.keys(directory).reduce((memo, key) => {
    const val = directory[key];

    if (typeof val === 'number') {
      return memo + val;
    }

    return memo + getSize(val);
  }, 0);
};

const getSumOfSizes = (directory, maxSize) => {
  const size = getSize(directory);
  let sum = 0;

  if (size <= maxSize) {
    sum += size;
  }

  // Run same function for children that are directories
  Object.keys(directory).forEach((key) => {
    if (typeof directory[key] === 'object') {
      sum += getSumOfSizes(directory[key], maxSize);
    }
  });

  return sum;
};

const getDirectorySizes = (directory, sizes = []) => {
  // Store size of current directory first
  sizes.push(getSize(directory));

  // Then recursirvely store size of nested directories
  Object.keys(directory).forEach((key) => {
    if (typeof directory[key] === 'object') {
      getDirectorySizes(directory[key], sizes);
    }
  });

  return sizes;
};

const run = async () => {
  const file = await fs.readFile('./input');
  const input = file.toString();
  const lines = input.split('\n').filter(Boolean);

  // Build directory with files and folders
  const directory = buildDirectory(lines);

  // Part 1
  // const sum = getSumOfSizes(directory, 100e3);
  // console.log('Sum of directories under 100e3:', sum);

  // Part 2
  const totalSpaceUsed = getSize(directory);
  const spaceAvailable = TOTAL_SPACE - totalSpaceUsed;
  const spaceNeeded = NEEDED_SPACE - spaceAvailable;

  // The trick is to get a list of all the sizes,
  // filter out ones that are too small,
  // then grab the smallest one remaining
  const sizes = getDirectorySizes(directory).filter((size) => size >= spaceNeeded);
  const minSize = Math.min(...sizes);
  console.log('Smallest directory that meets requirement:', minSize);
};

run();
