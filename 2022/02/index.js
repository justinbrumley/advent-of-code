const util = require('../../util');

const Rock = 1;
const Paper = 2;
const Scissors = 3;

const Loss = 0;
const Draw = 3;
const Win = 6;

const run = async () => {
  const lines = await util.getInput();

  let points = 0;

  lines.forEach((line) => {
    const [opponent, me] = line.split(' ');

    switch (opponent) {
      // Rock
      case 'A': {
        switch (me) {
          // Lose
          case 'X': {
            points += 0 + Scissors;
            break;
          }

          // Draw
          case 'Y': {
            points += 3 + Rock;
            break;
          }

          // Win
          case 'Z': {
            points += 6 + Paper;
            break;
          }
        }

        break;
      }

      // Paper
      case 'B': {
        switch (me) {
          // Lose
          case 'X': {
            points += 0 + Rock;
            break;
          }

          // Draw
          case 'Y': {
            points += 3 + Paper;
            break;
          }

          // Win
          case 'Z': {
            points += 6 + Scissors;
            break;
          }
        }

        break;
      }

      // Scissors
      case 'C': {
        switch (me) {
          // Lose
          case 'X': {
            points += 0 + Paper;
            break;
          }

          // Draw
          case 'Y': {
            points += 3 + Scissors;
            break;
          }

          // Win
          case 'Z': {
            points += 6 + Rock;
            break;
          }
        }

        break;
      }
    }
  });

  console.log('Score:', points);
};

run();
