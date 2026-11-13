# Лабораторная работа №7 — вариант 5 (Caching)

## Реализовано
- GitHub Actions CI при push и pull request.
- Матрица Go 1.21 и Go 1.22.
- Кэш `~/go/pkg/mod` и `~/.cache/go-build` через `actions/cache@v4`; ключ включает ОС, версию Go и хэш `go.sum`, задан restore-key.
- Запуск тестов с race detector и формированием `coverage.out`.
- Сохранение `coverage.out` как артефакта для каждой версии Go.
- Статический анализ через golangci-lint.
- Dockerfile и workflow публикации Docker Hub.

## Проверка кэширования и покрытия
После двух последовательных запусков workflow сравните длительность шага установки зависимостей/тестов и сообщения cache hit/cache miss на вкладке Actions. Кэш должен восстановиться по ключу ОС + Go + go.sum; при изменении зависимостей будет создан новый ключ.

Для получения отчёта о покрытии:
1. Откройте успешный запуск CI → Artifacts.
2. Скачайте `coverage-go-1.22` и распакуйте архив.
3. Выполните `go tool cover -html=coverage.out`.

Фактическое ускорение не указано: оно измеряется по двум запускам в GitHub Actions и зависит от runner и состояния кэша.

## Настройка секретов Docker Hub
В Settings → Secrets and variables → Actions добавьте:
- `DOCKERHUB_USERNAME`
- `DOCKERHUB_TOKEN`

Публикация запускается при push в main или при push тега вида `v*`. До добавления секретов workflow публикации не сможет выполнить вход в Docker Hub.

## Branch protection
В Settings → Branches создайте правило для `main` и включите требование успешных CI checks перед merge. Эта настройка выполняется в интерфейсе GitHub владельцем репозитория; автоматически в кодовую базу она не включена.
