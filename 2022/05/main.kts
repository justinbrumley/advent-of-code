package advent.of.code

import java.io.*

fun getInput(): List<String> = File("./input").useLines { it.toList() }

// Parses stacks from input, and returns a list of lists
fun parseStacks(lines: List<String>): MutableList<MutableList<Char>> {
  val stacks: MutableList<MutableList<Char>> = mutableListOf()

  var first = true

  for (line in lines) {
    if (line == "") {
      break
    }

    // Stacks are three characters wide, with one space between
    // But stack could be empty so loop instead of split string
    var index = 0

    while (index < line.length) {
      if (first) {
        stacks.add(mutableListOf())
      }

      val item = line.get(index + 1)

      if (item != ' ') {
        stacks.get(index / 4).add(line.get(index + 1))
      }

      index += 4
    }

    first = false
  }

  return stacks
}

// Move items from one stack to another based on commands
fun processStacks(lines: List<String>, stacks: MutableList<MutableList<Char>>) {
  var processing = false

  for (line in lines) {
    if (line == "") {
      processing = true
      continue
    }

    if (!processing) {
      continue
    }

    val commands = line.split(' ')

    // Hardcoding indices. I'm sure part two will have different commands though.
    val amount = commands.get(1).toInt()
    val from = commands.get(3).toInt()
    val to = commands.get(5).toInt()

    val fromStack = stacks.get(from - 1)
    val toStack = stacks.get(to - 1)

    // Store elements the crane is current holding in order
    var holding: MutableList<Char> = mutableListOf()

    for (i in 1..amount) {
      // Pop one from "from" and push to "to"
      val value = fromStack.get(0)
      holding.add(0, value)
      fromStack.removeAt(0)
    }

    // Pop the holding list on top of the target stack
    // Since the order is flipped,
    // looping over and adding one at a time will unflip it
    for (i in 1..holding.count()) {
      val value = holding.get(0)
      toStack.add(0, value)
      holding.removeAt(0)
    }
  }
}

// Pretty printer for the stacks
fun printStacks(stacks: MutableList<MutableList<Char>>) {
  for (stack in stacks) {
    println("$stack")
  }
}

fun main() {
  val input = getInput()

  val stacks = parseStacks(input)

  processStacks(input, stacks)

  printStacks(stacks)
}

main()
