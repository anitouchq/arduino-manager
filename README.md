# 🔧 Arduino Manager — Go + Fyne → GitHub Releases

Учебное GUI-приложение на Go и Fyne для просмотра информации из **Arduino CLI**. Приложение запускает команды CLI и показывает подключённые платы, установленные платформы (cores), библиотеки и версию инструмента.

## Возможности

- `arduino-cli board list` — поиск подключённых плат и вывод портов/FQBN.
- `arduino-cli core list` — список установленных board cores.
- `arduino-cli lib list` — список установленных библиотек.
- `arduino-cli version` — версия Arduino CLI.
- Отображение ошибок CLI и подсказки, если инструмент не установлен.

Arduino CLI устанавливается отдельно и должен быть доступен через системный PATH. Приложение не устанавливает прошивки автоматически.

## CI checks

- `gofmt` — форматирование;
- `go vet ./...` — статический анализ;
- `go test ./... -v` — тесты парсера списка плат;
- `go build` — smoke-build GUI.

Fyne использует CGO и системные графические библиотеки. Release-сборка выполняется на нативных Linux, macOS и Windows runner'ах.

## Release artifacts

| Платформа | Артефакт |
|---|---|
| Linux x64 | `arduino-manager-linux-x64` |
| macOS | `arduino-manager-macos` |
| Windows x64 | `arduino-manager-windows-x64.exe` |

Отправьте тег для создания релиза:

```bash
git tag v1.0.0
git push origin v1.0.0
```

## Локальный запуск

Требуется Go 1.27+, GCC, графические библиотеки Fyne и отдельно установленный Arduino CLI.

Ubuntu/Debian:

```bash
sudo apt-get install gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev
```

Запуск:

```bash
go mod tidy
go test ./... -v
go run .
```

## Permissions

Только Release job получает `contents: write` для публикации артефактов через встроенный `GITHUB_TOKEN`.

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
