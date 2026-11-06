<div align="center">

<p><a href="README.md">Русский</a> | <strong>English</strong></p>

<img src="assets/images/icon.png" alt="FLClash" width="88">

# FLClash PQ · Firefox 148

**An unofficial FLClash Android fork featuring the Firefox 148 TLS fingerprint and REALITY + X25519MLKEM768 compatibility.**

[Download releases](https://github.com/CaliostRa/FLClashPQ/releases) · [Fork source code](https://github.com/CaliostRa/FLClashPQ) · [Original FLClash](https://github.com/chen08209/FlClash)

</div>

> **Disclaimer:** This is an independent, unofficial fork, not an official FLClash or Mihomo release. Credit for the original application and UI belongs to their respective maintainers. Licensed under [GPL-3.0](LICENSE).

## Why this fork exists

Newer Xray-core versions (26.9.8 and later) introduced REALITY handshake requirements that can require the hybrid **X25519MLKEM768** key exchange. In the original FLClash 0.9.3, Mihomo's `firefox` fingerprint emulates Firefox 120, which lacks this group. As a result, a connection to an Xray-core 26.10.31 REALITY server may fail. The `chrome` fingerprint supports it, but some users reported connection instability on mobile networks.

The goal of this fork is to **keep FLClash's existing interface, settings, routing, DNS and subscriptions** while restoring the ability to use a modern `firefox` TLS fingerprint with current Xray/REALITY servers without downgrading the server.

## Changes compared with original FLClash 0.9.3

| Component | Original | FLClash PQ |
| --- | --- | --- |
| `firefox` TLS fingerprint | Firefox 120 in standard uTLS | **Firefox 148** via `legiz-ru/prizrak-utls` |
| REALITY key exchange | Older Firefox without X25519MLKEM768 | **X25519MLKEM768 + X25519** in ClientHello |
| REALITY client version | Original Mihomo behavior | Adjusted client-version bytes for newer Xray compatibility |
| `reality-opts.support-x25519mlkem768` | Original Mihomo behavior | Accepted as a deprecated compatibility option |
| Android application ID | `com.follow.clash` | **`com.follow.clash.pq`** for side-by-side installation |
| Android builds | Upstream APKs | **arm64-v8a** and **armeabi-v7a** APKs using this fork's GitHub Actions |

The implementation changes are in `core/go.mod`, a patch applied to `Clash.Meta`, Android build configuration, tests and GitHub Actions. The FLClash UI and existing VLESS, xHTTP, gRPC, Hysteria2 and Trojan implementations were intentionally retained.

See the [technical ARM64 verification report](reports/firefox148-android-arm64.md) and the [Prizrak uTLS project](https://github.com/legiz-ru/prizrak-utls).

## Download and install

**[GitHub Releases — FLClash PQ APKs](https://github.com/CaliostRa/FLClashPQ/releases)**. Development/test APKs may also be available in successful [GitHub Actions](https://github.com/CaliostRa/FLClashPQ/actions) workflow artifacts.

- **arm64-v8a:** modern Android ARM64 devices, including OnePlus 13.
- **armeabi-v7a:** older 32-bit ARMv7 Android devices.
- **x86_64:** no dedicated build currently published by this fork.

The separate package ID lets FLClash PQ be installed alongside the original FLClash. **Signing matters:** test builds and release builds can have different signing certificates. An APK can update an installed app only when its package ID and signing certificate are compatible. Export your settings before switching build types. This fork does not receive upstream FLClash app updates automatically.

### REALITY client settings

Keep your existing server address, UUID, SNI, REALITY public key, short ID and transport. For REALITY, you can use:

```yaml
client-fingerprint: firefox
reality-opts:
  public-key: "your existing public key"
  short-id: "your existing short ID"
  support-x25519mlkem768: true
```

The last option is retained for compatibility with existing Mihomo YAML profiles; in this fork, the modern Firefox fingerprint already provides the required hybrid key share. If your subscription generator forces `chrome`, change the server-side subscription mapping or a local override specifically for this client.

This patch addresses **REALITY compatibility**. It does not guarantee unrestricted access on every network, or reliable operation for every possible xHTTP configuration.

## Build from source

Baseline: FLClash 0.9.3, Flutter 3.47.6, Go 1.27.1, Android SDK/NDK r28c and Rust.

```bash
git clone --recursive https://github.com/CaliostRa/FLClashPQ.git
cd FLClashPQ
flutter pub get

# Apply the REALITY compatibility patch to the Mihomo submodule
./tool/apply_core_patches.sh

# Android ARM64
dart setup.dart android --arch arm64 --targets apk

# Android ARMv7
dart setup.dart android --arch arm --targets apk
```

Build outputs are located in `dist/`. GitHub Actions applies the same patch in its build jobs. See the [ARMv7 workflow](.github/workflows/build-armv7.yml) and the [release workflow](.github/workflows/android-arm64-pq.yaml).

## Credits and license

- **[chen08209/FlClash](https://github.com/chen08209/FlClash)** — original application and user interface.
- **[Mihomo / Clash.Meta](https://github.com/MetaCubeX/mihomo)** — proxy core.
- **[legiz-ru/prizrak-utls](https://github.com/legiz-ru/prizrak-utls)** — newer TLS fingerprint library.
- Compatibility fixes and Android release infrastructure for this unofficial fork are maintained here. This fork is not affiliated with or endorsed by the upstream maintainers.

Distributed under [GPL-3.0](LICENSE); third-party dependencies retain their respective licenses.
