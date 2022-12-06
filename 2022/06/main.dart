import 'dart:io';

String getInput() {
  return new File('./input').readAsStringSync();
}

int getStartOfPacket(String input) {
  for (int i = 3; i < input.length; i++) {
    // Grab previous 4 characters
    final code = input.substring(i - 3, i + 1);

    // Check for any duplicates
    final runes = code.runes.toList();

    var uniqueRunes = Set<int>();
    for (var i = 0; i < runes.length; i++) {
      uniqueRunes.add(runes[i]);
    }

    if (uniqueRunes.length == 4) {
      // Find 4 unique runes, so return index + 1 to signify the count
      return i + 1;
    }
  }

  return -1;
}

// Same as start of packet, but requires 14 unique characters instead of 4
int getStartOfMessage(String input) {
  for (int i = 13; i < input.length; i++) {
    // Grab previous 14 characters
    final code = input.substring(i - 13, i + 1);

    // Check for any duplicates
    final runes = code.runes.toList();

    var uniqueRunes = Set<int>();
    for (var i = 0; i < runes.length; i++) {
      uniqueRunes.add(runes[i]);
    }

    if (uniqueRunes.length == 14) {
      // Find 14 unique runes, so return index + 1 to signify the count
      return i + 1;
    }
  }

  return -1;
}

void main() {
  final input = getInput();
  final startOfPacket = getStartOfPacket(input);
  final startOfMessage = getStartOfMessage(input);

  print('Start of Packet: ${startOfPacket}');
  print('Start of Message: ${startOfMessage}');
}
