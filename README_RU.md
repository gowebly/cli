<img width="256px" alt="gowebly logo" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-logo.svg">

# gowebly – CLI-инструмент нового поколения для создания удивительных веб-приложений на языке Go с использованием htmx и hyperscript

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

[English][repo_url] | **Русский** | [简体中文][repo_readme_cn_url] | [Español][repo_readme_es_url]

С помощью этого CLI-инструмента можно легко создавать удивительные веб-приложения с **Go** на стороне бэкенда, с использованием [**htmx**][htmx_url], [**hyperscript**][hyperscript_url] и наиболее популярных атомарных/утилитарных **CSS-фреймворков** на стороне фронтенда.

Особенности:

- 100% **свободный** и **открытый** исходный код под лицензией [Apache 2.0][repo_license_url];
- Для **любого** уровня знаний и технической экспертизы разработчика;
- **Хорошо документирован**, содержит множество советов и подсказок от авторов;
- Кроссплатформенность и мультиархитектурность позволяют **успешно работать** под GNU/Linux, MS Windows (включая WSL) и Apple macOS;
- Умный CLI, который **делает большую часть** рутинной настройки и подготовки к продакшену;
- Помогает быстрее войти в стек технологий **Go** + **htmx** + **hyperscript**;
- Возможность простого добавления в свой проект готового к использованию и полностью настроенного атомарного/утилитарного **CSS-фреймворка**;
- Готов к установке в качестве **PWA** (Progressive Web App) в браузере или на мобильном устройстве;
- Поддерживает **режим live-reloading** для ваших CSS стилей;
- Имеет библиотеку **удобных** функций для вашего Go-кода;
- Содержит исчерпывающий **пример** использования из коробки.

> 💬 От авторов: Чтобы дать вам полное представление о проекте, мы записали короткое [📺 видео][gowebly_youtube_video_url] и подготовили вводную [📝 статью][gowebly_devto_article_url], демонстрирующую основные возможности `gowebly` CLI.

## ⚡️ Быстрый старт

Сначала [скачайте][go_download_url] и установите **Go**. Требуется версия `1.21` (или выше).

Теперь вы можете использовать `gowebly` без установки. Просто запустите [`go run`][go_run_url], который создаст новый проект с [конфигурацией][repo_default_config] по умолчанию:

```console
go run github.com/gowebly/gowebly@latest create
```

Вот и все! 🔥 Замечательное веб-приложение, использующее встроенный пакет **net/http** (в качестве Go-бэкенда), **htmx** и **hyperscript** — доступны в ваших HTML-шаблонах Go.

### 🔹 Полный путь к быстрому старту на Go

Если вы все же хотите установить `gowebly` CLI в свою систему средствами Golang, воспользуйтесь командой [`go install`][go_install_url]:

```console
go install github.com/gowebly/gowebly@latest
```

### 🍺 Быстрый старт с Homebrew

Пользователям GNU/Linux и Apple macOS доступен способ установки `gowebly` CLI через [Homebrew][brew_url].

Тапните новую формулу:

```console
brew tap gowebly/tap
```

Установите `gowebly`:

```console
brew install gowebly/tap/gowebly
```

### 🐳 Быстрый старт с Docker

Вы можете использовать `gowebly` CLI из нашего [официального Docker-образа][docker_image_url] и запускать его в изолированном контейнере:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest create
```

### 📦 Другой путь для быстрого старта

Скачать готовые `exe` файлы для Windows, пакеты `deb`, `rpm`, `apk` или Arch Linux — можно на странице [Releases][repo_releases_url].

## 📖 Полное руководство пользователя

Чтобы получить полное руководство по использованию и понять основные принципы работы `gowebly` CLI, мы подготовили исчерпывающее объяснение каждой команды в этом файле README.

> 💬 От авторов: Мы всегда ценим ваше время и хотим, чтобы вы как можно скорее начали создавать действительно замечательные веб-продукты на этом потрясающем технологическом стеке!

Мы надеемся, что вы найдете ответы на все вопросы! 👌 Но если вы не нашли нужной информации, не стесняйтесь создать [issue][repo_issues_url] или отправить [PR][repo_pull_request_url] в этот репозиторий.

### `init`

Команда создания **дефолтного** файла конфигурации ([`.gowebly.yml`][repo_default_config]) в текущей папке.

```console
gowebly init
```

> 💡 Примечание: Конечно, вы можете пропустить этот шаг, если вас устраивает следующая конфигурация по умолчанию для вашего нового проекта:
>
> - Имена модулей Go (`go.mod`) и `package.json` установлены, как **project**;
> - Без Go-фреймворка для backend-части (только встроенный пакет **net/http**);
> - Без CSS-фреймворка для frontend-части (только стили по умолчанию для примера кода);
> - Среда выполнения JavaScript для frontend-части будет использовать **Node.js**;
> - Порт сервера – `5000`, таймаут (в секундах): `5` для чтения, `10` для записи;
> - Последние версии **htmx** и **hyperscript**.

<img width="720" alt="gowebly init" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_init.png">

Как правило, созданный файл конфигурации содержит следующие настройки:

```yaml
backend:
   module_name: project # (string) option can be any name of your Go module (for example, 'github.com/user/project')
   go_framework: default # (string) option can be one of the values: 'fiber', 'echo', 'chi', or 'default'
   port: 5000 # (int) option can be any port that is not taken up on your system
   timeout:
      read: 5 # (int) option can be any number of seconds, 5 is recommended
      write: 10 # (int) option can be any number of seconds, 10 is recommended

frontend:
   package_name: project # (string) option can be any name of your package.json (for example, 'project')
   css_framework: default # (string) option can be one of the values: 'tailwindcss', 'unocss', or 'default'
   runtime_environment: default # (string) option can be one of the values: 'bun', or 'default'
   htmx: latest # (string) option can be any existing version
   hyperscript: latest # (string) option can be any existing version
```

Но вы можете выбрать любой другой **Go-фреймворк** для бэкенда вашего проекта:

| Go-фреймворк | Описание                                                                         |
| ------------ | -------------------------------------------------------------------------------- |
| `default`    | Без какого-либо Go-фреймворка (только встроенный пакет [net/http][net_http_url]) |
| `fiber`      | Использовать [Fiber][fiber_url] веб-фреймворк в качестве бэкенда                 |
| `echo`       | Использовать [Echo][echo_url] веб-фреймворк в качестве бэкенда                   |
| `chi`        | Использовать композитный роутер [chi][chi_url] в качестве бэкенда                |

Кроме того, вы можете выбрать для своего проекта версии **htmx**, **hyperscript**, а также один из самых популярных атомарных/утилитарных **CSS-фреймворков**:

| CSS-фреймворк | Описание                                                                    |
| ------------- | --------------------------------------------------------------------------- |
| `default`     | Без какого-либо CSS-фреймворка (только стили по умолчанию для примера кода) |
| `tailwindcss` | Использовать [Tailwind CSS][tailwindcss_url] в качестве CSS-фреймворка      |
| `unocss`      | Использовать [UnoCSS][unocss_url] в качестве CSS-фреймворка                 |

А также, вы можете выбрать другую среду выполнения JavaScript для
frontend-части:

| Среда выполнения JavaScript | Описание                                                                  |
| --------------------------- | ------------------------------------------------------------------------- |
| `default`                   | Использовать [Node.js][nodejs_url] в качестве среды выполнения JavaScript |
| `bun`                       | Использовать [Bun][bun_url] в качестве среды выполнения JavaScript        |

### `create`

Команда для создания нового проекта с бэкендом **Go**, **htmx** и **hyperscript**, а также (_опционально_) атомарным/утилитарным **CSS-фреймворком**.

```console
gowebly create
```

> 💡 Примечание: Если не выполнить команду `init` для создания файла конфигурации (`.gowebly.yml`), то `gowebly` CLI создает новый проект с конфигурацией [по умолчанию][repo_default_config].

<img width="720" alt="gowebly create" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_create.png">

Каждый раз, когда вы выполняете команду `create` для своего проекта:

1. CLI проверяет конфигурацию и применяет все настройки к текущему проекту;
2. CLI подготавливает backend-часть проекта (генерирует структуру проекта и необходимые файлы, запускает `go mod tidy`);
3. CLI подготавливает frontend-часть проекта (генерирует необходимые файлы, запускает `npm|bun install` и `npm|bun run build:dev` в первый раз);
4. CLI скачивает минимизированные версии **htmx** и **hyperscript** (с официального и надежного CDN [unpkg.com][unpkg_url]) в папку `./static` и размещает их в виде разделенных тегов `<script>` (внизу тега `<body>`) в HTML-шаблоне Go [`templates/main.html`][repo_main_layout].

Как правило, созданный проект содержит следующие файлы и папки:

```console
.
├── assets
│   └── styles.css
├── static
│   ├── favicons
│   │   ├── apple-touch-icon.png
│   │   ├── favicon.ico
│   │   ├── favicon.png
│   │   ├── favicon.svg
│   │   ├── manifest-screenshot-desktop.jpeg
│   │   ├── manifest-mobile-desktop.jpeg
│   │   └── manifest-touch-icon.svg
│   ├── images
│   │   └── logo.svg
│   ├── htmx.min.js
│   ├── hyperscript.min.js
│   ├── styles.css
│   └── manifest.json
├── templates
│   ├── pages
│   │   └── index.html
│   └── main.html
├── .gitignore
├── go.mod
├── go.sum
├── handlers.go
├── main.go
├── package.json
└── server.go
```

### `run`

Команда для запуска проекта в режиме **разработки** (непроизводственном).

```console
gowebly run
```

> 💡 Примечание: Если не выполнить команду `init` для создания файла конфигурации (`.gowebly.yml`), то `gowebly` CLI запустит ваш проект с конфигурацией [по умолчанию][repo_default_config].

<img width="720" alt="gowebly run" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_run.png">

Каждый раз, когда вы выполняете команду `run` для своего проекта:

1. CLI проверяет конфигурацию и применяет все настройки к текущему проекту;
2. CLI подготавливает frontend-часть проекта (выполняется команда `npm|bun run watch`);
3. CLI подготавливает версию для разработки (непроизводственную) выбранного **CSS-фреймворка** в папке `./static` и размещает ее в виде тега `<link>` (внизу тега `<head>`) в HTML-шаблоне Go [`templates/main.html`][repo_main_layout];
4. CLI запускает бэкенд проекта с настройками из конфигурации по умолчанию (или из конфигурационного файла `.gowebly.yml`) простой командой `go run`.

### `build`

Команда для сборки проекта для **продакшена** и подготовки Docker-файлов для развертывания.

```console
gowebly build [OPTION]
```

> 💡 Примечание: Если не выполнить команду `init` для создания конфигурационного файла (`.gowebly.yml`), то `gowebly` CLI собирает проект с конфигурацией [по умолчанию][repo_default_config].

<img width="720" alt="gowebly build" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_build.png">

Вы можете добавить следующие опции:

| Название опции  | Описание                                                                               | Обязательно? |
| --------------- | -------------------------------------------------------------------------------------- | ------------ |
| `--skip-docker` | Пропустить процесс генерации Docker-файлов (полезно, если у вас есть свои собственные) | Нет          |

Каждый раз, когда вы выполняете команду `build` для своего проекта:

1. CLI проверяет конфигурацию и применяет все настройки к текущему проекту;
2. CLI подготавливает версию для разработки (непроизводственную) выбранного **CSS-фреймворка** в папке `./static` и размещает ее в виде тега `<link>` (внизу тега `<head>`) в HTML-шаблоне Go [`templates/main.html`][repo_main_layout];
3. CLI подготавливает продакшен-версию выбранного **CSS-фреймворка** и размещает ее в виде тега `<link>` (в нижней части тега `<head>`) в шаблоне Go HTML [`templates/main.html`][repo_main_layout];
4. Если опция `--skip-docker` не была задана, CLI генерирует в корне папки проекта понятные и хорошо документированные Docker-файлы (`.dockerignore`, `Dockerfile`, `docker-compose.yml`) для развертывания проекта в изолированных контейнерах через [Portainer][portainer_url] (_рекомендуется_), или вручную, на удаленном сервере.

### `doctor`

Команда для отображения полезной **информации** о вашей системе.

```console
gowebly doctor
```

> 💡 Примечание: Это очень полезно для процесса отладки, или создания нового [issue][repo_issues_url] с сообщением об ошибке в этом GitHub-репозитории.

<img width="720" alt="gowebly doctor" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_doctor.png">

Каждый раз, когда вы выполняете команду `doctor` для вашей системы:

1. CLI проверяет версии всех необходимых инструментов для успешной реализации проекта (таких, как `gowebly` CLI, Go, Node.js, Docker, Docker Compose и т.д.);
2. CLI формирует отчет с указанием установленной версии для каждого инструмента.

## 🙋 Удобные функции

CLI `gowebly` имеет библиотеку удобных [функций][gowebly_helpers_url] для вашего кода. Это поможет вам еще быстрее начать создавать потрясающие веб-приложения на Go.

```console
go get -u github.com/gowebly/helpers
```

> 💡 Примечание: Библиотека `gowebly helpers` **уже** включена в Go-бэкенд, создаваемый командой `create`, но вы можете использовать эти функции и в любых других проектах.

## 🎯 Мотивация к созданию

Скажите, как часто вам приходилось начинать новый проект с нуля и выполнять мучительные ручные настройки? 🤔 Особенно, когда вы только начинаете своё знакомство с новой технологией или стеком, где все для вас в новинку.

Для многих разработчиков, _в том числе и для нас_, этот процесс максимально утомителен и даже уныл, и не несет никакой полезной нагрузки. Это **очень** удручающий процесс, который может сильно оттолкнуть любого разработчика от этих технологий.

Почему бы просто не отдать всю эту ужасную ручную работу машинам? Пусть они сделают всю тяжелую работу за нас, а мы будем просто создавать потрясающие веб-продукты и не думать о сборке и развертывании.

Именно поэтому мы создали `gowebly` CLI и библиотеку удобных функций, которые помогут вам создавать потрясающие веб-приложения на **Go** с использованием **htmx**, **hyperscript** и популярных атомарных/утилитарных **CSS-фреймворков**.

Мы здесь, чтобы избавить вас (_и себя_) от этой рутинной боли! ✨

> 💬 От авторов: Ранее, мы уже спасали мир — это был [Create Go App][cgapp_url] (да, это тоже наш проект). Статистика [GitHub stars][cgapp_stars_url] этого проекта не может врать: более **2,2 тыс.** разработчиков любого уровня и из разных стран начинают новый проект с помощью этого CLI-инструмента.

## 🏆 Взаимовыгодное сотрудничество

Если вам понравился проект `gowebly` и вы нашли его полезным для своих задач то, пожалуйста, нажмите на кнопку 👁️ **Watch**, чтобы не пропустить уведомления о выходе новых версий, и поставьте ему 🌟 **GitHub Star**!

Это очень **мотивирует** нас делать этот продукт **еще** лучше.

<img width="100%" alt="gowebly star and watch" src="https://github.com/gowebly/gowebly/assets/11155743/6f92ec26-1fe3-44c6-9a13-3abd3ffa58eb">

А теперь я приглашаю вас принять участие в этом проекте! Давайте работать **вместе** над созданием и популяризацией **самого полезного** инструмента для разработчиков в Интернете на сегодня.

- [Issues][repo_issues_url]: задавайте вопросы и предлагайте свои улучшения.
- [Pull requests][repo_pull_request_url]: присылайте свои улучшения.
- Расскажите несколько слов о проекте в своих социальных сетях и блогах (Dev.to, Medium, Хабр и других).

Ваши PR, вопросы и любые слова — приветствуются! Спасибо 😘

### 🌟 Звездные люди

[![gowebly stargazers][repo_badge_reporoster_url]][repo_stargazers_url]

## ⚠️ Лицензия

[`gowebly`][repo_url] — это свободное программное обеспечение с открытым исходным кодом, лицензируемое по [Apache 2.0 лицензии][repo_license_url], созданное и поддерживаемое [Vic Shóstak][author_url] с 🩵 к людям и роботам. Официальный логотип распространяется под [лицензией Creative Commons][repo_cc_license_url] (CC BY-SA 4.0 International).

<!-- Go links -->

[go_download_url]: https://golang.org/dl/
[go_run_url]: https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program
[go_install_url]: https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies
[go_report_url]: https://goreportcard.com/report/github.com/gowebly/gowebly
[go_dev_url]: https://pkg.go.dev/github.com/gowebly/gowebly
[go_version_img]: https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
[go_code_coverage_url]: https://codecov.io/gh/gowebly/gowebly
[go_code_coverage_img]: https://img.shields.io/codecov/c/gh/gowebly/gowebly.svg?logo=codecov&style=for-the-badge
[go_report_img]: https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none

<!-- Repository links -->

[repo_url]: https://github.com/gowebly/gowebly
[repo_issues_url]: https://github.com/gowebly/gowebly/issues
[repo_pull_request_url]: https://github.com/gowebly/gowebly/pulls
[repo_releases_url]: https://github.com/gowebly/gowebly/releases
[repo_license_url]: https://github.com/gowebly/gowebly/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_license_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_readme_ru_url]: https://github.com/gowebly/gowebly/blob/main/README_RU.md
[repo_readme_cn_url]: https://github.com/gowebly/gowebly/blob/main/README_CN.md
[repo_readme_es_url]: https://github.com/gowebly/gowebly/blob/main/README_ES.md
[repo_default_config]: https://github.com/gowebly/gowebly/blob/main/internal/attachments/configs/default.yml
[repo_main_layout]: https://github.com/gowebly/gowebly/blob/main/internal/attachments/templates/frontend/main.html
[repo_stargazers_url]: https://github.com/gowebly/gowebly/stargazers
[repo_badge_reporoster_url]: https://reporoster.com/stars/notext/gowebly/gowebly

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->

[gowebly_youtube_video_url]: https://www.youtube.com/watch?v=qazYscnLku4
[gowebly_devto_article_url]: https://dev.to/koddr/a-next-generation-cli-tool-for-building-amazing-web-apps-in-go-using-htmx-hyperscript-336d
[gowebly_helpers_url]: https://github.com/gowebly/helpers
[cgapp_url]: https://github.com/create-go-app/cli
[cgapp_stars_url]: https://github.com/create-go-app/cli/stargazers
[nodejs_url]: https://nodejs.org
[bun_url]: https://bun.sh
[docker_image_url]: https://hub.docker.com/repository/docker/gowebly/gowebly
[portainer_url]: https://docs.portainer.io
[brew_url]: https://brew.sh
[htmx_url]: https://htmx.org
[hyperscript_url]: https://hyperscript.org
[tailwindcss_url]: https://tailwindcss.com
[unocss_url]: https://unocss.dev
[unpkg_url]: https://unpkg.com
[net_http_url]: https://pkg.go.dev/net/http
[fiber_url]: https://github.com/gofiber/fiber
[echo_url]: https://github.com/labstack/echo
[chi_url]: https://github.com/go-chi/chi
