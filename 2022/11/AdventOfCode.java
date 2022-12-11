import java.io.*;
import java.util.*;
import java.math.*;

class Monkey {
  List<Long> items = new ArrayList<>();

  String operator = "+";
  int modifier = 0;
  boolean modifyBySelf = false;

  int divisor = 0;
  int modulo = 0;

  // Monkey to throw too if divisible by divisor
  int targetIdx = 0;

  // Monkey  to throw too if NOT divisible by divisor
  int altTargetIdx = 0;

  // Track how many total items this monkey inspects
  int itemsInspected = 0;

  public void play(List<Monkey> monkeys) {
    // Loop over all current items the monkey has
    for (int i = 0; i < items.size(); i++) {
      long item = items.get(i);

      // Modify the stress level of the item
      item = this.modifyLevel(item) % this.modulo;

      // (Part One Only)
      // Divide by three (constant)
      // item /= 3;

      int target = this.getTarget(item);

      Monkey monkey = monkeys.get(target);

      // Give this item to target monkey
      monkey.items.add(item);
    }

    // All items should've been thrown so we can just clear the items list
    this.itemsInspected += items.size();
    this.items.clear();
  }

  /**
   * Check value to determine which monkey to throw to
   */
  private int getTarget(long val) {
    if (val % this.divisor == 0) {
      return this.targetIdx;
    }

    return this.altTargetIdx;
  }

  /**
   * Use monkey's modifier to manipulate stress level of item
   */
  private long modifyLevel(long val) {
    long modifier = this.modifier;
    if (this.modifyBySelf) {
      modifier = val;
    }

    switch (this.operator) {
      case "+":
        return val + modifier;

      case "*":
        return val * modifier;
    }

    return val;
  }
}

public class AdventOfCode {
  public static void main(String[] args) {
    // Create a File object
    File file = new File("input");

    List<Monkey> monkeys = new ArrayList<Monkey>();
    int modulo = 1;

    // Start by processing file and initializing all the monkeys
    try {
      BufferedReader br = new BufferedReader(new FileReader(file));

      while (true) {
        String line = br.readLine();
        if (line == null) { break; }

        String[] parts = line.split(":", 0);

        String key = parts[0].trim();
        if (key.length() == 0) { continue; }

        if (key.indexOf("Monkey") == 0) {
          monkeys.add(new Monkey());
          continue;
        }

        String value = parts[1].trim();
        Monkey monkey = monkeys.get(monkeys.size() - 1);

        if (key.indexOf("Starting items") == 0) {
          List<Long> items = new ArrayList<>();
          String[] values = value.split(", ", 0);

          for (int i = 0; i < values.length; i++) {
            items.add(Long.parseLong(values[i]));
          }

          monkey.items = items;
          continue;
        }

        if (key.indexOf("Operation") == 0) {
          // Skip all the way up to operator
          String[] formula = value.substring(10).split(" ");

          monkey.operator = formula[0];
          String modifier = formula[1];

          if (modifier.indexOf("old") == 0) {
            monkey.modifyBySelf = true;
          } else {
            monkey.modifier = Integer.parseInt(modifier);
          }

          continue;
        }

        if (key.indexOf("Test") == 0) {
          String divisor = value.substring(13);
          monkey.divisor = Integer.parseInt(divisor);
          modulo *= monkey.divisor;
          continue;
        }

        if (key.indexOf("If true") == 0) {
          String idx = value.substring(16);
          monkey.targetIdx = Integer.parseInt(idx);
          continue;
        }

        if (key.indexOf("If false") == 0) {
          String idx = value.substring(16);
          monkey.altTargetIdx = Integer.parseInt(idx);
          continue;
        }
      }

      br.close();
    } catch (IOException e) {
      /* noop */
    }

    /* (Part One) */
    /*
    // Monkeys are ready to play (for 20 turns)
    for (int i = 0; i < 20; i++) {
      for (int j = 0; j < monkeys.size(); j++) {
        monkeys.get(j).play(monkeys);
      }
    }

    for (int i = 0; i < monkeys.size(); i++) {
      System.out.printf("Monkey %d items: %d\n", i, monkeys.get(i).itemsInspected);
    }
    */

    /* (Part Two) */
    for (int j = 0; j < monkeys.size(); j++) {
      monkeys.get(j).modulo = modulo;
    }

    for (int i = 0; i < 10000; i++) {
      for (int j = 0; j < monkeys.size(); j++) {
        monkeys.get(j).play(monkeys);
      }

      System.out.printf("Monkey items after turn %d (modulo %d)\n=========================\n", i + 1, monkeys.get(0).modulo);

      for (int j = 0; j < monkeys.size(); j++) {
        System.out.printf("Monkey %d : Items Inspected: %d\n", j, monkeys.get(j).itemsInspected);
      }

      System.out.println();
    }
  }
}

