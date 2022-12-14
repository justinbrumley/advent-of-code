#include <fstream>
#include <iostream>
#include <list>
#include <string>

using namespace std;

struct ListOrInt {
  bool isList;
  int value;
  string original;
  list<ListOrInt> values;
};

int min(int x, int y) {
  if (x < y) {
    return x;
  }

  return y;
}

ListOrInt get(list<ListOrInt> items, int index) {
  int i = 0;
  for (const ListOrInt& item : items) {
    if (index == i) {
      return item;
    }

    i++;
  }

  return ListOrInt();
}

// Compare two ListOrInt values. It assumes a and b are both lists.
int compare(ListOrInt a, ListOrInt b) {
  int minSize = min(a.values.size(), b.values.size());

  for (int i = 0; i < minSize; i++) {
    ListOrInt aVal = get(a.values, i);
    ListOrInt bVal = get(b.values, i);

    // First check if both values are numbers
    // If they are, just compare them
    if (!aVal.isList && !bVal.isList) {
      if (aVal.value < bVal.value) {
        return -1;
      } else if (aVal.value > bVal.value) {
        return 1;
      }

      // If numbers are equal, move on to the next list item
      continue;
    }

    // If values are mixed types,
    // convert the one that is a number to a list
    // with that number as the only item in the list
    if (aVal.isList && !bVal.isList) {
      ListOrInt newItem;
      newItem.isList = false;
      newItem.value = bVal.value;

      bVal.isList = true;
      bVal.values.push_back(newItem);
    } else if (!aVal.isList && bVal.isList) {
      ListOrInt newItem;
      newItem.isList = false;
      newItem.value = aVal.value;

      aVal.isList = true;
      aVal.values.push_back(newItem);
    }

    // If both values are lists, recursively compare them.
    if (aVal.isList && bVal.isList) {
      int comparison = compare(aVal, bVal);
      if (comparison != 0) {
        return comparison;
      }
    }
  }

  // Final check, compare the lengths of arrays
  if (a.values.size() < b.values.size()) {
    return -1;
  } else if (a.values.size() > b.values.size()) {
    return 1;
  }

  // If we reach here, then the lists are identical
  return 0;
}

int main() {
  ifstream input_file("input");

  if (!input_file.is_open()) {
    cerr << "Error: Unable to open input file" << endl;
    return 1;
  }

  list<string> lines;
  string line;
  while (getline(input_file, line)) {
    lines.push_back(line);
  }

  input_file.close();

  list<ListOrInt> rows;

  // For part two, add the separators
  lines.push_back("[[2]]");
  lines.push_back("[[6]]");

  for (const string& line : lines) {
    // Skip blank lines
    if (line == "") { continue; }

    // Each row is itself a list
    ListOrInt item;
    item.isList = true;
    item.original = line;

    // Use a list to track which parent we are looking at
    list<ListOrInt*> stack;
    stack.push_back(&item);

    // Build out the row by checking each character
    // Skipping the first and last brackets
    for (int i = 1; i < line.size() - 1; i++) {
      char ch = line.at(i);

      // Do nothing, just move on to next character
      if (ch == ',') { continue; }

      switch (ch) {
        case '[': {
          // Start of a new nested list
          ListOrInt* currItem = stack.back();

          ListOrInt newItem;
          newItem.isList = true;

          // Add list to current list
          currItem->values.push_back(newItem);

          // Then switch pointer to new list
          stack.push_back(&(currItem->values.back()));

          break;
        }

        case ']': {
          // End of current list, jump back to parent
          stack.pop_back();
          break;
        }

        default: {
          // Parse number up to next comma or close bracket
          string s = line.substr(i);
          string num = s.substr(0, min(s.find(","), s.find("]")));

          ListOrInt newItem;
          newItem.isList = false;
          newItem.value = stoi(num);

          ListOrInt* currItem = stack.back();

          currItem->values.push_back(newItem);

          // Skip over the next n - 1 characters
          i += num.size() - 1;
        }
      }
    }

    rows.push_back(item);
  }

  // Part One
  int idxSum = 0;
  int idx = 1;
  for (int i = 0; i < rows.size(); i += 2, idx++) {
    ListOrInt a = get(rows, i);
    ListOrInt b = get(rows, i + 1);
    if (compare(a, b) == -1) { idxSum += idx; }
  }

  cout << "Sum of indices of correct pairs: " << idxSum << endl;

  // Part Two
  // Start by sorting the list
  rows.sort([](ListOrInt a, ListOrInt b) { return compare(a, b) < 0; });

  // Now find the indices of the two separate elements
  // and multiply them together
  int i = 1;
  int key = 0;
  for (const ListOrInt& item : rows) {
    if (item.original == "[[2]]" || item.original == "[[6]]") {
      if (key == 0) {
        key = i;
      } else {
        key *= i;
        break;
      }
    }

    i++;
  }

  cout << "Decoder Key: " << key << endl;

  return 0;
}

