# keeptrack

Small desktop app that watches your Azure DevOps PRs and tickets in the background so you don't have to keep a browser tab open.

Built with [Wails](https://v3.wails.io/) (Go + Vue).

## Install

Head to the [releases page](https://github.com/SilahicAmil/keeptrack/releases) and grab the latest build for your OS:

- Windows: `keeptrack-windows-amd64.exe`
- macOS: `keeptrack-macos-arm64`
- Linux: `keeptrack-linux-amd64`

Run it. First launch walks you through connecting Azure DevOps (org, project, and a PAT with read access to Work Items and Code).

## Contributing

PRs welcome. Standard flow:

1. Fork the repo
2. `git checkout -b your-branch`
3. Commit your changes
4. Push and [open a PR](https://github.com/SilahicAmil/keeptrack/compare) against `main`

For anything bigger than a small fix, please [open an issue](https://github.com/SilahicAmil/keeptrack/issues) first.

## Running locally

You'll need:

- [Go](https://go.dev/dl/) 1.25+
- [Node.js](https://nodejs.org/) 20+
- [Task](https://taskfile.dev/)
- [Wails v3 CLI](https://v3.wails.io/getting-started/installation/): `go install -v github.com/wailsapp/wails/v3/cmd/wails3@latest`
- On Linux: `libgtk-3-dev` and `libwebkit2gtk-4.1-dev`

Then:

```bash
task dev      # hot-reloaded dev build
task build    # production build into bin/
```

## Releases

Tagged pushes (`v*`) trigger [`.github/workflows/release.yml`](.github/workflows/release.yml), which builds binaries for Windows, macOS, and Linux and attaches them to a draft GitHub release.

```bash
git tag v0.1.0
git push origin v0.1.0
```

## Project layout

```
azuredevops/   Azure DevOps API client
config/        Config file handling
internal/
  services/    Wails services exposed to the frontend
poller/        Background poller
store/         SQLite cache
frontend/      Vue 3 + Tailwind UI
main.go        Entry point
```

## Built with

- [Wails v3](https://v3.wails.io/)
- [Go](https://go.dev/)
- [Vue 3](https://vuejs.org/) / [Vue Router](https://router.vuejs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Vite](https://vitejs.dev/)
- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)
