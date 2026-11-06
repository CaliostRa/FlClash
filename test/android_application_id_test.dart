import 'dart:io';

import 'package:flutter_test/flutter_test.dart';

void main() {
  test('Android package installs beside upstream FLClash', () {
    final gradle = File('android/app/build.gradle.kts').readAsStringSync();

    expect(gradle, contains('namespace = "com.follow.clash"'));
    expect(gradle, contains('applicationId = "com.follow.clash.pq"'));
  });
}
