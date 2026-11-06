import 'dart:io';

import 'package:flutter_test/flutter_test.dart';

void main() {
  test('Android release tags require a persistent signing key', () {
    final gradle = File('android/app/build.gradle.kts').readAsStringSync();
    final workflow = File(
      '.github/workflows/android-arm64-pq.yaml',
    ).readAsStringSync();

    expect(gradle, contains('FLCLASH_ANDROID_KEYSTORE_PATH'));
    expect(gradle, contains('FLCLASH_REQUIRE_RELEASE_SIGNING'));
    expect(gradle, contains('throw GradleException'));

    expect(workflow, contains("- 'pq-v*'"));
    expect(workflow, contains('secrets.ANDROID_KEYSTORE_BASE64'));
    expect(workflow, contains('secrets.ANDROID_STORE_PASSWORD'));
    expect(workflow, contains('secrets.ANDROID_KEY_ALIAS'));
    expect(workflow, contains('secrets.ANDROID_KEY_PASSWORD'));
    expect(workflow, contains(r'${{ runner.temp }}'));
    expect(workflow, contains('FLCLASH_REQUIRE_RELEASE_SIGNING'));
    expect(workflow, contains('contents: write'));
    expect(workflow, contains('softprops/action-gh-release'));
    expect(workflow, isNot(contains('> android/app/keystore.jks')));
  });
}
