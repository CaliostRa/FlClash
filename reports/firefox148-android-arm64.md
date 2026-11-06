# FLClash Firefox 148 Android ARM64 verification

Date: 2026-11-06

## Source baseline

- FLClash: `be8473660552` (`0.9.3`, upstream `main`)
- Mihomo submodule: `024245dc72bc` (`1.19.33`)
- Reference Prizrak-Core: `v1.19.33-r11` (`90df9cb1e527`)
- uTLS replacement: `github.com/legiz-ru/prizrak-utls` at
  `2807fc753c53856049883d2964d17a6fa852d935`
- Target server compatibility: Xray-core `26.10.31`

## Changes

- Replaced Mihomo's uTLS implementation through a pinned Go module replacement.
- Preserved Firefox 148's `X25519MLKEM768, X25519` key-share order in the
  REALITY handshake instead of stripping the hybrid share by default.
- Aligned REALITY client-version bytes with Prizrak-Core (`26.7.11`) for
  Xray's `minClientVer` compatibility gate.
- Kept `support-x25519mlkem768` accepted as a deprecated no-op, so existing
  YAML profiles remain parseable.
- Kept the FLClash UI and existing protocol implementations unchanged.
- Added compatibility tests for VLESS REALITY/Vision, xHTTP, gRPC, Hysteria2,
  Trojan, and existing YAML parsing.
- Changed only the Android application ID to `com.follow.clash.pq`; the
  namespace and Kotlin package remain unchanged.
- Added a pinned Android ARM64 GitHub Actions workflow with tag-triggered,
  persistent release signing and GitHub Release publishing.
- Added Gradle distribution checksum verification.
- No Remnawave server or configuration was accessed or changed.

## Verification

- Go Firefox 148/uTLS test: passed.
- Mihomo REALITY hybrid key-share test: passed.
- REALITY client-version bytes test (`26.7.11`): passed.
- YAML protocol compatibility test: passed.
- Full core test package: passed when run as the unprivileged `nobody` user;
  its inode ownership test intentionally rejects execution as root.
- Android ARM64 `go vet`: passed.
- `flutter analyze --no-fatal-infos`: passed with no issues.
- Full Flutter suite: `4,867` passed, `6` skipped.
- Flutter coverage: `87.65%` (`34,723/39,614`), all group floors passed.
- Flutter Android application-ID test: passed.
- Android JVM unit tests: passed (`309` actionable tasks).
- Core IPC smoke test: the host core connected to a Unix socket and remained
  running for two seconds.
- APK manifest application ID: `com.follow.clash.pq`.
- APK ABI set: only `arm64-v8a`.
- Bundled `libclash.so`: ELF64 AArch64 shared object.
- APK signature: valid Android v2 release signature.
- Release certificate SHA-256:
  `61b7c8a2e24314fd5337fb4703e2ba1df70318ad2b4fa4247dcc9e8b6b94de63`.

The release workflow fails closed when its persistent signing key or any
credential is missing. The key and credentials are stored outside Git and in
GitHub Actions Secrets; only the public certificate fingerprint is committed.
An earlier debug-signed test build must be uninstalled once before installing
the first release-signed APK; later release APKs retain update compatibility.

## Artifact

- File: `dist/FlClash-0.9.3-android-arm64-v8a.apk`
- Size: `60,861,097` bytes
- SHA-256: `fa45bb6aa0d83fb5fdefea8f70570130b67e464783dd7f0729bf92befadd42ed`

## Unavailable verification

No OnePlus 13 device, Xray server address, REALITY credentials, or test profile
was available in the build environment. Therefore installation/UI launch on the
physical phone and a live connection to Xray-core 26.10.31 were not tested.
The automated checks verify the generated ClientHello, core startup, APK
identity/ABI/signature, and configuration parsing, but do not replace a live
end-to-end connection test.
