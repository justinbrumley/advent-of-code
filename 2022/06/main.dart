import 'dart:io';

String getInput() {
  return new File('./input').readAsStringSync();
}

int getStartOfSequence(String input, int seqLen) {
  for (int i = seqLen - 1; i < input.length; i++) {
    // Grab previous X characters
    final code = input.substring(i - seqLen + 1, i + 1);

    // Check for any duplicates
    final runes = code.runes.toSet();

    if (runes.length == seqLen) {
      // Find X unique runes, so return index + 1 to signify the count
      return i + 1;
    }
  }

  return -1;
}

void main() {
  final input = getInput();
  final startOfPacket = getStartOfSequence(input, 4);
  final startOfMessage = getStartOfSequence(input, 14);

  print('Start of Packet: ${startOfPacket}');
  print('Start of Message: ${startOfMessage}');
}
