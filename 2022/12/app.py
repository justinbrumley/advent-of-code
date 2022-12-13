import sys
import os
import copy
import time
import math

sys.setrecursionlimit(10000)

def clear_console():
    os.system('clear')

class Node:
    def __init__(self, x, y, value):
        self.x = x
        self.y = y
        self.value = ord(value)
        self.distance = float('inf')
        self.visited = False
        self.parent = None
        self.is_start = False
        self.is_end = False

        # Set initial distance to 0 if starting position
        if self.value == 83:
            self.value = ord("a")
            self.distance = 0
            self.is_start = True
        elif self.value == 69:
            self.value = ord("z")
            self.is_end = True


class Grid:
    def __init__(self, nodes):
        self.max_x = 0
        self.max_y = 0

        self.nodes = copy.deepcopy(nodes)

        for node in self.nodes:
            if node.x > self.max_x:
                self.max_x = node.x
            if node.y > self.max_y:
                self.max_y = node.y

            if node.is_start:
                self.current_node = node
            elif node.is_end:
                self.target_node = node

        # Store nodes in grid for easy reference
        self.arr = []
        for y in range(self.max_y + 1):
            self.arr.append([])
            for x in range(self.max_x + 1):
                self.arr[y].append(None)

        for node in self.nodes:
            self.arr[node.y][node.x] = node


    # Set start node using x, y (since nodes could be clones and can't be trusted to pass by reference)
    # We need to find the current starting node and remove the distance and is_start properties
    def set_start_node(self, x, y):
        self.current_node = self.arr[y][x]
        for node in self.nodes:
            if node.x == x and node.y == y:
                node.is_start = True
                node.distance = 0
            elif node.is_start == True:
                node.is_start = False
                node.distance = float('inf')


    # Return a list of neighbors of the provided node
    # Filters out visited nodes
    # Filters out nodes that can't be reached
    def get_neighbors(self):
        x = self.current_node.x
        y = self.current_node.y

        neighbors = []

        if y > 0:
            node = self.arr[y - 1][x]
            neighbors.append(node)

        if y < len(self.arr) - 1:
            node = self.arr[y + 1][x]
            neighbors.append(node)

        if x > 0:
            node = self.arr[y][x - 1]
            neighbors.append(node)

        if x < len(self.arr[y]) - 1:
            node = self.arr[y][x + 1]
            neighbors.append(node)

        # Filter out visited neighbors
        neighbors = filter(lambda node: node.visited == False, neighbors)

        # Filter out neighbors with a height larger than current height + 1
        neighbors = filter(lambda node: node.value <= self.current_node.value + 1, neighbors)

        return neighbors


    # (Dijkstra's Kinda)
    # Grab neighbors, filter out neighbors with value too high,
    # then assign distance to them
    def calc_distances(self):
        # Set current node as visited
        self.current_node.visited = True

        if self.current_node == self.target_node:
            # We are done!
            return

        # Loop over neighbors and assign new distances to them
        neighbors = self.get_neighbors()
        for node in neighbors:
            # Hardcoded distance of 1 between all nodes on the grid
            if node.distance > self.current_node.distance + 1:
                node.distance = self.current_node.distance + 1
                node.parent = self.current_node

        # Find the next smallest unvisited node and move there
        # Then calculate distances again
        current_node = None
        for node in self.nodes:
            if node.visited == False and (current_node == None or node.distance < current_node.distance):
                current_node = node

        # self.draw_grid(None)

        if current_node != None and not math.isinf(current_node.distance):
            self.current_node = current_node
            self.calc_distances()


    def get_steps_to_target(self):
        # Will add distance values to nodes
        grid.calc_distances()

        steps = 0
        parent = grid.target_node.parent
        path = set()

        if parent == None:
            # Path not found
            return float('inf'), path

        while parent:
            steps += 1
            path.add(parent)
            parent = parent.parent

        return steps, path

    def draw_grid(self, path):
        clear_console()

        for row in self.arr:
            line = ""
            for node in row:
                is_in_path = False
                if path != None:
                    out = ""
                    for n in path:
                        if n.x == node.x and n.y == node.y:
                            is_in_path = True
                            break

                if node.is_start:
                    line += 'S'
                elif node.is_end:
                    line += 'E'
                elif is_in_path:
                    line += "■"
                elif node.visited == True:
                    line += '-'
                else:
                    line += "."

            print(line)


# Build the initial grid of nodes first
nodes = []
with open("input", "r") as file:
    y = 0
    for line in file:
        x = 0
        for char in list(line.strip()):
            node = Node(x, y, char);
            nodes.append(node)
            x += 1
        y += 1


# Part One
# Initialize the Grid object, which will determine starting and ending positions
# grid = Grid(nodes)
# steps, path = grid.get_steps_to_target()
# grid.draw_grid(path)
# print("Steps to reach target: %d | path: %d" % (steps, len(path)))

# Part Two (brute force)
# We need a new grid for every possible starting spot (a)
# Then get number of steps to target to find least
grids = []
lowest_steps = float('inf')
code = ord('a')

start = time.time()
for node in nodes:
    if node.value == code:
        grid = Grid(nodes)
        grid.set_start_node(node.x, node.y)
        steps, path = grid.get_steps_to_target()
        if math.isinf(steps):
            continue
        if steps < lowest_steps:
            lowest_steps = steps

end = time.time()
print("Elapsed time is  {}".format(end-start))
print("Lowest possible steps for path starting at 'a' is %d" % lowest_steps)
