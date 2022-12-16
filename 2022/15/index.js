const fs = require('fs/promises');

const regex = /^.*=(-?\d+).*=(-?\d+).*=(-?\d+).*=(-?\d+)$/;

const sensors = [];
const beacons = {};

let min = Infinity;
let max = -Infinity;

const getDistance = (a, b) => {
  const distance = Math.abs(a.x - b.x) + Math.abs(a.y - b.y);
  return distance;
};

const run = async () => {
  const file = await fs.readFile('./input');
  const lines = file.toString().split('\n');

  lines.filter(Boolean).forEach((l) => {
    const parts = l.match(regex);
    const [line, sensorX, sensorY, beaconX, beaconY] = parts;

    const sensor = {
      x: parseInt(sensorX, 10),
      y: parseInt(sensorY, 10),
    };

    sensor.key = `${sensor.x} ${sensor.y}`;

    const beacon = {
      x: parseInt(beaconX, 10),
      y: parseInt(beaconY, 10),
    };

    beacon.key = `${beacon.x} ${beacon.y}`;

    // Calculate the distance between the sensor and the beacon
    const distance = getDistance(sensor, beacon);
    sensor.range = distance;

    // Store location of sensor and beacon
    sensors.push(sensor);
    beacons[beacon.key] = beacon;

    min = Math.min(sensor.x - distance, min);
    max = Math.max(sensor.x + distance, max);
  });

  // Part One
  // Find number of empty points at row 2,000,000
  /*
  let count = 0;
  for (let i = min; i <= max; i += 1) {
    const point = { x: i, y: 2000000 };

    // First check if the point is a beacon
    if (beacons[`${point.x} ${point.y}`]) { continue; }

    // Then ensure that point is within range of a sensor
    let inRange = false;
    for (let j = 0; j < sensors.length; j += 1) {
      const sensor = sensors[j];

      if (getDistance(point, sensor) <= sensor.range) {
        inRange = true;
        break;
      }
    }

    if (inRange) {
      count += 1;
    }
  }

  console.log(`Found ${count} empty points`);
  */

  // Part Two
  const findHiddenPoint = () => {
    // Loop over grid and find true value
    for (let y = 0; y <= 4e6; y += 1) {
      for (let x = 0; x <= 4e6; x += 1) {
        let inRange = false;

        for (let i = 0; i < sensors.length; i += 1) {
          const distance = getDistance({ x, y }, sensors[i]);

          if (distance <= sensors[i].range) {
            // Skip to other side of the sensor
            inRange = true;
            if (sensors[i].x > x) {
              x += (sensors[i].x - x) * 2;
            } else {
              // Already on the right side of the sensor, so just jump to the end of it
              x += (sensors[i].range - distance);
            }
            break;
          }
        }

        if (!inRange) {
          return { x, y };
        }
      }
    }
  };

  const hiddenPoint = findHiddenPoint();
  console.log('Found hidden point at:', hiddenPoint);
};

run().then(() => {
  process.exit();
});
