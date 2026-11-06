<div align="center">

<p><strong>Русский</strong> | <a href="README_EN.md">English</a></p>

<img src="assets/images/icon.png" alt="FLClash" width="88">

# FLClash PQ · Firefox 148

**Неофициальная Android-модификация FLClash с современным TLS-отпечатком Firefox 148 и поддержкой REALITY + X25519MLKEM768.**

[Скачать релизы](https://github.com/CaliostRa/FLClashPQ/releases) · [Исходники модификации](https://github.com/CaliostRa/FLClashPQ) · [Оригинальный FLClash](https://github.com/chen08209/FlClash)

</div>

> **Важно:** это независимый пользовательский форк, а не официальный релиз FLClash или Mihomo. Исходный проект и интерфейс принадлежат своим авторам. Лицензия — [GPL-3.0](LICENSE).

## Зачем создан форк

В Xray-core 26.9.8+ изменились требования к рукопожатию REALITY: для некоторых новых серверных конфигураций требуется гибридная группа обмена ключами **X25519MLKEM768**. В FLClash 0.9.3 обычный отпечаток Mihomo `firefox` имитировал Firefox 120 без такой группы, поэтому соединение с Xray-core 26.10.31 могло не устанавливаться. Отпечаток `chrome` поддерживал её, но в некоторых мобильных сетях наблюдались нестабильные соединения.

Цель форка — **сохранить настройки, маршрутизацию, DNS, подписки и интерфейс FLClash**, но вернуть возможность использовать современный отпечаток `firefox` с актуальным Xray/REALITY без отката сервера.

## Что изменено относительно оригинального FLClash 0.9.3

| Компонент | Оригинал | FLClash PQ |
|---|---|---|
| TLS-отпечаток `firefox` | Firefox 120 в штатном uTLS | **Firefox 148** через `legiz-ru/prizrak-utls` |
| REALITY и гибридный обмен | Старый Firefox без X25519MLKEM768 | **X25519MLKEM768 + X25519** в ClientHello |
| Проверка версии REALITY | Исходное поведение Mihomo | Совместимые с новым Xray байты версии клиента |
| Старое поле `reality-opts.support-x25519mlkem768` | Используется штатным Mihomo | Принимается как совместимый устаревший параметр |
| Идентификатор Android-приложения | `com.follow.clash` | **`com.follow.clash.pq`** — установка рядом с оригиналом |
| Сборки этого форка | Официальные APK проекта | APK для **arm64-v8a** и **armeabi-v7a**, собственные GitHub Actions |

Технически изменения находятся в `core/go.mod`, патче для `Clash.Meta`, Android-конфигурации сборки, тестах и GitHub Actions. Реализации VLESS, xHTTP, gRPC, Hysteria2, Trojan и пользовательский интерфейс намеренно не заменялись целиком.

Подробный технический отчёт: [reports/firefox148-android-arm64.md](reports/firefox148-android-arm64.md). Исходная библиотека: [Prizrak uTLS](https://github.com/legiz-ru/prizrak-utls).

## Скачать и установить

**[GitHub Releases — APK этого форка](https://github.com/CaliostRa/FLClashPQ/releases)**. Тестовые APK также можно найти в артефактах успешных запусков [GitHub Actions](https://github.com/CaliostRa/FLClashPQ/actions).

- **arm64-v8a** — большинство современных Android-устройств, включая OnePlus 13.
- **armeabi-v7a** — 32-битные устройства ARM (ARMv7).
- **x86_64** — APK этого форка пока не выпускается.

Установку рядом с оригинальным FLClash обеспечивает отдельный package ID. **Важно о подписях:** тестовая и релизная сборки могут иметь разные сертификаты. Обновление поверх ранее установленной сборки возможно только при совместимых подписях и идентификаторе приложения. Экспортируйте настройки перед заменой. Приложение из этого репозитория не обновляется автоматически из официального репозитория FLClash.

### Настройки REALITY в клиенте

Используйте существующий сервер, UUID, SNI, public-key, short-id и транспорт. Для REALITY можно указать:

```yaml
client-fingerprint: firefox
reality-opts:
  public-key: "ваш public-key"
  short-id: "ваш short-id"
  support-x25519mlkem768: true
```

Последняя строка сохранена для совместимости с существующими Mihomo-профилями; в этой модификации современный Firefox уже содержит нужный гибридный key share. Если подписка автоматически выдаёт `chrome`, измените правило генерации подписки на стороне сервера либо локальное переопределение только для этого клиента.

Это исправление **совместимости REALITY**, а не гарантия прохождения любых сетевых фильтров, стабильности любой мобильной сети или работы всех конфигураций xHTTP.

## Сборка из исходников

Основа — FLClash 0.9.3, Flutter 3.47.6, Go 1.27.1, Android SDK/NDK r28c и Rust.

```bash
git clone --recursive https://github.com/CaliostRa/FLClashPQ.git
cd FLClashPQ
flutter pub get
./tool/apply_core_patches.sh

# ARM64
dart setup.dart android --arch arm64 --targets apk

# ARMv7
dart setup.dart android --arch arm --targets apk
```

Исходный патч ядра применяется при сборке workflow командой `./tool/apply_core_patches.sh`. Для ручной сборки примените его перед `dart setup.dart`. Итоговый APK находится в `dist/`. Автоматическая сборка ARMv7: [build-armv7.yml](.github/workflows/build-armv7.yml). Релизный workflow: [android-arm64-pq.yaml](.github/workflows/android-arm64-pq.yaml).

## Авторы и лицензия

- Оригинальный проект: **[chen08209/FlClash](https://github.com/chen08209/FlClash)** — интерфейс и приложение.
- Ядро: **[Mihomo / Clash.Meta](https://github.com/MetaCubeX/mihomo)**.
- Обновлённая TLS-библиотека: **[legiz-ru/prizrak-utls](https://github.com/legiz-ru/prizrak-utls)**.
- Доработки совместимости и Android-сборок находятся в этом форке. Это неофициальная сборка, без поддержки со стороны авторов оригинального проекта.

Код распространяется в соответствии с [GPL-3.0](LICENSE). Используемые компоненты сохраняют собственные лицензии.
