import 'dart:math';

List<int> GenerateRandomArray(int size) {
  final random = Random();
  return List<int>.generate(size, (_) => random.nextInt(100));
}