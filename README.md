# keeptrack

Small desktop app that watches your Azure DevOps tickets and PRs (coming soon) in the background so you don't have to rely on Azure Devops to send notifications.

Built with [Wails](https://v3.wails.io/) (Go + Vue).

## Install

Head to the [releases page](https://github.com/SilahicAmil/keeptrack/releases) and grab the latest build for your OS:

- Windows: `keeptrack-windows-amd64.exe`
- macOS: `keeptrack-macos-arm64.dmg`
- Linux: `keeptrack-linux-amd64.AppImage` (run `chmod +x` on it once, then double-click)

Run it. First launch walks you through connecting Azure DevOps. You'll need:

- Your **org**
- Your **project**
- A **Personal Access Token (PAT)**

keeptrack only ever _reads_ from Azure DevOps. It never writes, comments, or changes anything. A god-level PAT will work fine, but for your own safety we recommend a read-only token.

## Contributing

PRs are welcome. Please [open an issue](https://github.com/SilahicAmil/keeptrack/issues) first before doing any changes.

### Clone

```bash
git clone https://github.com/SilahicAmil/keeptrack.git
cd keeptrack
```

### Running locally

You'll need:

- [Go](https://go.dev/dl/) 1.25+
- [Node.js](https://nodejs.org/) 20+
- [Wails v3 CLI](https://v3.wails.io/getting-started/installation/): `go install -v github.com/wailsapp/wails/v3/cmd/wails3@latest`
- On Linux: `libgtk-3-dev` and `libwebkit2gtk-4.1-dev`

Then:

```bash
wails3 dev      # hot-reloaded dev build
wails3 build    # production build into bin/
```

See the [Wails v3 docs](https://v3.wails.io/) for more.

### Project layout

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

## Technologies used

- [Wails v3](https://v3.wails.io/)
- [Go](https://go.dev/)
- [Vue 3](https://vuejs.org/) / [Vue Router](https://router.vuejs.org/)
- [Tailwind CSS](https://tailwindcss.com/)
- [Vite](https://vitejs.dev/)
- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite)
