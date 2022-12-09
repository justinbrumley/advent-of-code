string[] lines = System.IO.File.ReadAllLines("./input");

Point head = new Point(0, 0, false);

/* Part One */
Point tail = new Point(0, 0, true);

for (int i = 0; i < lines.Length; i++) {
  string line = lines[i];
  string[] parts = line.Split(' ');

  string direction = parts[0];
  int amount = int.Parse(parts[1]);

  for (int j = 0; j < amount; j++) {
    head.Move(direction);
    tail.MoveTowards(head);
  }
}

// tail.PrintHistory();
Console.WriteLine($"Part One - Unique Tail Positions: {tail.History.Count}");

/* Part Two */
head = new Point(0, 0, false);
List<Point> knots = new List<Point>();

// Initialize all 9 knots
for (int i = 0; i < 9; i++) {
  // Only need to track history on last knot
  knots.Add(new Point(0, 0, i == 8));
}

for (int i = 0; i < lines.Length; i++) {
  string line = lines[i];
  string[] parts = line.Split(' ');

  string direction = parts[0];
  int amount = int.Parse(parts[1]);

  for (int j = 0; j < amount; j++) {
    head.Move(direction);

    // Don't show my old professors this loop in a loop in a loop
    for (int k = 0; k < knots.Count; k++) {
      Point h = k == 0 ? head : knots[k - 1];
      knots[k].MoveTowards(h);
    }
  }
}

tail = knots[knots.Count - 1];
// tail.PrintHistory();
Console.WriteLine($"Part Two - Unique Tail Positions: {tail.History.Count}");

class Point {
  public int X;
  public int Y;

  // Unique positions this point has been in
  public List<Point> History = new List<Point>();

  public Point(int x, int y, bool withHistory) {
    X = x;
    Y = y;

    if (withHistory) {
      History.Add(new Point(X, Y, false));
    }
  }

  /**
   * Check if point is "touching" provided point
   */
  public bool IsTouching(Point point) {
    int xDist = X - point.X;
    int yDist = Y - point.Y;

    return (xDist <= 1 && xDist >= -1) && (yDist <= 1 && yDist >= -1);
  }

  /**
   * Move one space in the provided direction
   */
  public void Move(string direction) {
    switch (direction) {
      case "U":
        Y += 1;
        break;

      case "D":
        Y -= 1;
        break;

      case "L":
        X -= 1;
        break;

      case "R":
        X += 1;
        break;
    }
  }

  /**
   * Update X and Y to move towards point.
   * Also track history of this point.
   */
  public void MoveTowards(Point point) {
    if (IsTouching(point)) {
      return;
    }

    int xDist = Math.Abs(point.X - X);
    int yDist = Math.Abs(point.Y - Y);

    if ((xDist >= 2 && yDist >= 1) || (xDist >= 1 && yDist >= 2)) {
      // Diagonal Movement
      X += (point.X > X) ? 1 : -1;
      Y += (point.Y > Y) ? 1 : -1;
    } else if (xDist > 1) {
      // Horizontal Movement
      X += (point.X > X) ? 1 : -1;
    } else if (yDist > 1) {
      // Vertical Movement
      Y += (point.Y > Y) ? 1 : -1;
    }

    TrackHistory();
  }

  /**
   * Add current position to history list,
   * if that point hasn't already been visited
   */
  public void TrackHistory() {
    for (int i = 0; i < History.Count; i++) {
      Point point = History[i];

      if (point.X == X && point.Y == Y) {
        return;
      }
    }

    History.Add(new Point(X, Y, false));
  }

  /**
   * Print history of points in X,Y list
   * For debugging
   */
  public void PrintHistory() {
    for (int i = 0; i < History.Count; i++) {
      Console.WriteLine($"{History[i].X}, {History[i].Y}");
    }
  }
}
