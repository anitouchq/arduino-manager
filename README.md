# arduino-manager

[![CI](https://github.com/anitouchq/arduino-manager/actions/workflows/ci.yml/badge.svg)](https://github.com/anitouchq/arduino-manager/actions)

Учебное GUI-приложение на Go и Fyne для просмотра информации из **Arduino CLI**. Приложение запускает команды CLI и показывает подключённые платы, установленные платформы (cores), библиотеки и версию инструмента.

## Возможности

- `arduino-cli board list` — поиск подключённых плат и вывод портов/FQBN.
- `arduino-cli core list` — список установленных board cores.
- `arduino-cli lib list` — список установленных библиотек.
- `arduino-cli version` — версия Arduino CLI.
- Вывод ошибок CLI и подсказка, если инструмент не установлен.

Arduino CLI устанавливается отдельно и должен быть доступен через системный PATH. Приложение не загружает прошивки и не прошивает платы автоматически.

## CI checks

- `gofmt` — проверка форматирования;
- `go vet ./...` — статический анализ;
- `go test ./... -v` — тесты логики разбора вывода;
- `go build` — проверочная сборка GUI.

Fyne использует CGO и системные графические библиотеки. Release-сборка выполняется на нативных runner'ах Linux, macOS и Windows.

## Release artifacts

| Платформа | Артефакт |
|---|---|
| Linux x64 | `arduino-manager-linux-x64` |
| macOS | `arduino-manager-macos` |
| Windows x64 | `arduino-manager-windows-x64.exe` |

Отправьте тег версии, чтобы запустить сборку и публикацию:

```bash
git tag v1.0.0
git push origin v1.0.0
```

Готовые файлы будут доступны в [GitHub Releases](https://github.com/anitouchq/arduino-manager/releases).

## Локальный запуск

Требуются Go 1.27+, GCC, графические библиотеки Fyne и отдельно установленный Arduino CLI.

Для Ubuntu/Debian:

```bash
sudo apt-get install gcc libgl1-mesa-dev xorg-dev libwayland-dev libxkbcommon-dev
go test ./... -v
go run .
```

## Структура проекта

| Файл или каталог | Назначение |
|---|---|
| `.github/workflows/ci.yml` | Проверки, сборка под три ОС и публикация Release |
| `main.go` | Создание окна Fyne и запуск команд Arduino CLI |
| `board.go` | Логика обработки информации о платах |
| `board_test.go` | Тесты обработки данных о платах |
| `go.mod` | Имя Go-модуля и прямые зависимости |
| `go.sum` | Контрольные суммы модулей |
| `.gitignore` | Локальные файлы, исключаемые из Git |

## Результаты

- [GitHub Actions](https://github.com/anitouchq/arduino-manager/actions)
- [GitHub Releases](https://github.com/anitouchq/arduino-manager/releases)

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
