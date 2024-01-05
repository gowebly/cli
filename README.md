<div align="center">

<a href="https://gowebly.org" target="_blank" title="Go to the Gowebly CLI website"><img width="196px" alt="gowebly logo" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-logo.svg"></a>

<a name="readme-top"></a>

# The Gowebly CLI

A next-generation CLI tool that makes it easy to create amazing web applications<br/>with **Go** on the backend, using **htmx** & **hyperscript** and the most popular<br/>**CSS** frameworks on the frontend.

**English** · [Русский][docs_ru_url] · [简体中文][docs_zh_hk_url] · [Español][docs_es_url]

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]<br/>
[![Code coverage][go_code_coverage_img]][go_code_coverage_url]
[![License][repo_license_img]][repo_license_url]

**&searr;&nbsp;&nbsp;Share the project's link to your friends&nbsp;&nbsp;&swarr;**

[![Share on X][x_share_img]][x_share_url]
[![Share on Telegram][telegram_share_img]][telegram_share_url]
[![Share on WhatsApp][whatsapp_share_img]][whatsapp_share_url]
[![Share on Reddit][reddit_share_img]][reddit_share_url]

<a href="https://gowebly.org" target="_blank"><img width="99%" alt="gowebly create command" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly_create.gif"></a>

</div>

## ✨ Features

- 100% **free** and **open source** under the [Apache 2.0][repo_license_url] license;
- For **any** developer's level of knowledge and technical expertise, as the intelligent CLI does most of the routine project setup for you, creates an understandable structure, and prepares code for use and deployment in production;
- Cross-platform and multi-architecture allows **successful running** on any GNU/Linux distros, Microsoft Windows (including WSL) and Apple macOS;
- [**Well-documented**][docs_url], includes translations in **many other languages** ([Русский][docs_ru_url], [简体中文][docs_zh_hk_url], [Español][docs_es_url]);
- Enables you to start a new project faster with **Go**, **htmx** & **hyperscript** libraries;
- Supports the built-in **net/http** package and the most popular Go web frameworks and routers out of the box, such as **Fiber**, **Gin**, **Echo**, **Chi** and **HttpRouter**;
- Supports the most popular CSS frameworks out of the box, such as **Tailwind CSS**, **daisyUI**, **Flowbite**, **UnoCSS**, **Bootstrap** and **Bulma**;
- Supports a way to build HTML for the htmx projects with Go using the **Templ** package;
- Supports a **live-reload** mode for your Go code and frontend files using **Air** tool;
- Ready to install as **PWA** (Progressive Web App) in your browser or mobile device;
- Has a library of **user-friendly** helpers for your Go code.

## ⚡️ Quick start

> [!NOTE]
> If you are looking for a version of the **Gowebly** CLI below `v2.0.0`, it's in [this][repo_branch_v1_url] branch.

First, [download][go_download_url] and install **Go**. Version `1.21` (or higher) is required.

Now, you can use the **Gowebly** CLI without installation. Just run it with [`go run`][go_run_url] to create a new project with a [default][repo_default_config_url] configuration:

```console
go run github.com/gowebly/gowebly@latest create
```

That's it! 🔥 A wonderful web application has been created in the current folder.

<div align="right">

[&nwarr; Back to top](#readme-top)

</div>

### 🐳 Docker-way to quick start

Feel free to using the **Gowebly** CLI from our [official Docker image][docker_image_url] and run it in the isolated container:

```console
docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} gowebly/gowebly:latest create
```

> [!IMPORTANT]
> This Docker image works **only** on the GNU/Linux systems (`amd64` or `arm64`, including WSL).

<div align="right">

[&nwarr; Back to top](#readme-top)

</div>

### 🍺 Homebrew-way to quick start

GNU/Linux and Apple macOS users available way to install **Gowebly** CLI via [Homebrew][brew_url].

Tap a new formula:

```console
brew tap gowebly/tap
```

Install:

```console
brew install gowebly/tap/gowebly
```

<div align="right">

[&nwarr; Back to top](#readme-top)

</div>

### 📦 Other way to quick start

Download ready-made `exe` files for Windows, `deb` (for Debian, Ubuntu), `rpm` (for CentOS, Fedora), `apk` (for Alpine), or Arch Linux packages from the [Releases][repo_releases_url] page.

<div align="right">

[&nwarr; Back to top](#readme-top)

</div>

## 📖 Complete user guide

I always treasure your time and want you to start building really great web products on this awesome technology stack as soon as possible! Therefore, to get a complete guide to use and understand the basic principles of the **Gowebly** CLI, we have prepared a comprehensive explanation of the project in this 📖 [**Complete user guide**][docs_url].

<a href="https://gowebly.org" target="_blank" title="Go to the Gowebly's Complete user guide"><img width="480px" alt="gowebly docs banner" src="https://raw.githubusercontent.com/gowebly/.github/main/images/gowebly-docs-banner.svg"></a>

I have taken care to make it **as comfortable as possible** for you to learn this wonderful tool, so each CLI command has a sufficient textual description, as well as a visual diagram of how it works.

> [!IMPORTANT]
> Don't forget to switch the documentation to your language to make it even more comfortable to learn new knowledge! Supported languages: [English][docs_url], [Русский][docs_ru_url], [简体中文][docs_zh_hk_url], [Español][docs_es_url].

<div align="right">

[&nwarr; Back to top](#readme-top)

</div>

### The learning path

It's highly recommended to start exploring with short articles "[**What is Gowebly CLI?**][docs_what_is_gowebly_cli_url]" and "[**How does it work?**][docs_how_it_works_url]" to understand the basic principle and the main components built into the **Gowebly** CLI.

Next steps are:

1. [Install the CLI to your system](https://gowebly.org/complete-user-guide/installation)
2. [Configure your project](https://gowebly.org/complete-user-guide/configuration)
3. [Start creating a new project](https://gowebly.org/complete-user-guide/create-new-project)
4. [Running your project locally](https://gowebly.org/complete-user-guide/run-your-project)
5. [Build your project for the production](https://gowebly.org/complete-user-guide/build-project)

Hope you find answers to all of your questions! 😉

<div align="right">

[&nwarr; Back to top](#readme-top)

</div>

## 🎯 Motivation to create

Tell me, how often have you had to start a new project from scratch and had to make painful manual configurations? 🤔 Especially, when you are just getting acquainted with a new technology or stack, where everything is new to you.

For many developers, _including me_, this process is as tedious and even depressing as possible, and doesn't carry any useful workload. It is a **very** frustrating process that can push any developer away from technology a lot.

Why not just give all that awful manual work to machines? Let them do all the hard work for us, and we will just create awesome web products and not have to think about build and deploy.

That's why the **Gowebly** CLI was born. It allows you to start a new project faster with **Go** and **htmx** with **hyperscript**, or modern **TypeScript** and **CSS** frameworks.

I am here to save you from this routine pain! ✨

<div align="right">

[&nwarr; Back to top](#readme-top)

</div>

## 🏆 A win-win cooperation

If you liked the **Gowebly** CLI and found it useful for your tasks, please click a 👁️ **Watch** button to avoid missing notifications about new versions, and give it a 🌟 **GitHub Star**!

It really **motivates** me to make this product **even** better.

<img width="100%" alt="gowebly star and watch" src="https://github.com/gowebly/gowebly/assets/11155743/6f92ec26-1fe3-44c6-9a13-3abd3ffa58eb">

And now, I invite you to participate in this project! Let's work **together** to create and popularize the **most useful** tool for developers on the web today.

- [Issues][repo_issues_url]: ask questions and submit your features.
- [Pull requests][repo_pull_request_url]: send your improvements to the current codebase.
- [Discussions][repo_discussions_url]: discuss and share your ideas.
- Share the project's link to your friends on [X (Twitter)][x_share_url], [Telegram][telegram_share_url], [WhatsApp][whatsapp_share_url], [Reddit][reddit_share_url].
- Say a few words about the project on your social networks and blogs ([Dev.to][dev_to_url], [Medium][medium_url], [Хабр][habr_url], and so on).

Your PRs, issues & any words are welcome! Thank you 😘

<div align="right">

[&nwarr; Back to top](#readme-top)

</div>

### 🌟 Stargazers

<img src="https://starchart.cc/gowebly/gowebly.svg" alt="Stargazers over time" style="max-width: 100%"/>

## ⚠️ License

[`The Gowebly CLI`][repo_url] is free and open-source software licensed under the [Apache 2.0 License][repo_license_url], created and supported by [Vic Shóstak][author_url] with 🩵 for people and robots. Official logo distributed under the [Creative Commons License][repo_cc_license_url] (CC BY-SA 4.0 International).

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
[repo_branch_v1_url]: https://github.com/gowebly/gowebly/tree/v1
[repo_issues_url]: https://github.com/gowebly/gowebly/issues
[repo_pull_request_url]: https://github.com/gowebly/gowebly/pulls
[repo_discussions_url]: https://github.com/gowebly/gowebly/discussions
[repo_releases_url]: https://github.com/gowebly/gowebly/releases
[repo_license_url]: https://github.com/gowebly/gowebly/blob/main/LICENSE
[repo_license_img]: https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none
[repo_cc_license_url]: https://creativecommons.org/licenses/by-sa/4.0/
[repo_default_config_url]: https://github.com/gowebly/gowebly/blob/main/internal/attachments/configs/default.yml

<!-- Docs links -->

[docs_url]: https://gowebly.org
[docs_ru_url]: https://gowebly.org/ru/
[docs_zh_hk_url]: https://gowebly.org/zh_HK/
[docs_es_url]: https://gowebly.org/es/
[docs_what_is_gowebly_cli_url]: https://gowebly.org/getting-started
[docs_how_it_works_url]: https://gowebly.org/getting-started/how-does-it-work

<!-- Author links -->

[author_url]: https://github.com/koddr

<!-- Readme links -->

[cgapp_url]: https://github.com/create-go-app/cli
[cgapp_stars_url]: https://github.com/create-go-app/cli/stargazers
[htmx_url]: https://htmx.org
[hyperscript_url]: https://hyperscript.org
[brew_url]: https://brew.sh
[docker_image_url]: https://hub.docker.com/repository/docker/gowebly/gowebly

<!-- Social links -->

[dev_to_url]: https://dev.to
[medium_url]: https://medium.com
[habr_url]: https://habr.com
[x_share_url]: https://x.com/intent/tweet?hashtags=gowebly%2Cgo%2Chtmx&text=A%20next-generation%20CLI%20tool%20to%20easily%20build%20amazing%20web%20applications%20with%20Go%20on%20the%20backend%2C%20using%20htmx%20%26%20hyperscript%20and%20the%20most%20popular%20CSS%20frameworks%20on%20the%20frontend.&url=https%3A%2F%2Fgowebly.org
[telegram_share_url]: https://t.me/share/url?text=A%20next-generation%20CLI%20tool%20to%20easily%20build%20amazing%20web%20applications%20with%20Go%20on%20the%20backend%2C%20using%20htmx%20%26%20hyperscript%20and%20the%20most%20popular%20CSS%20frameworks%20on%20the%20frontend.%20%23gowebly%20%23go%20%23htmx&url=https%3A%2F%2Fgowebly.org
[whatsapp_share_url]: https://api.whatsapp.com/send?text=A%20next-generation%20CLI%20tool%20to%20easily%20build%20amazing%20web%20applications%20with%20Go%20on%20the%20backend%2C%20using%20htmx%20%26%20hyperscript%20and%20the%20most%20popular%20CSS%20frameworks%20on%20the%20frontend.%20https%3A%2F%2Fgowebly.org
[reddit_share_url]: https://www.reddit.com/submit?title=A%20next-generation%20CLI%20tool%20to%20easily%20build%20amazing%20web%20applications%20with%20Go%20on%20the%20backend%2C%20using%20htmx%20%26%20hyperscript%20and%20the%20most%20popular%20CSS%20frameworks%20on%20the%20frontend.%20%23gowebly%20%23go%20%23htmx&url=https%3A%2F%2Fgowebly.org
[x_share_img]: https://img.shields.io/badge/x_(twitter)-black?style=for-the-badge&logo=x
[telegram_share_img]: https://img.shields.io/badge/telegram-black?style=for-the-badge&logo=telegram
[whatsapp_share_img]: https://img.shields.io/badge/whatsapp-black?style=for-the-badge&logo=whatsapp
[reddit_share_img]: https://img.shields.io/badge/reddit-black?style=for-the-badge&logo=reddit
