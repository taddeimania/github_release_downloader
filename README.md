# Github Release Downloader

### Purpose

This project's intention is to build a single binary that can deliver a private github release asset.

### Configuration

Rename the `config.template.json` file to `config.json` and replace the configuration values as follows:

```json
{
  "access_token": "YOUR GITHUB ACCESS TOKEN",
  "repo_owner": "the repository owner",
  "repo_name": "the repository name"
}
```

### Releasing

To release your downloader, first make sure you've configured your `config.json` file with working values and have tested the downloader works with:

```bash
go run main.go
```

If you see the asset you expect, then you are good to build.

Run the `./build.sh` command and it will prompt you what you want to name your release downloader. This will prepend the downloader files with whatever name you give it. I wouldn't suggest giving it a name with spaces or special characters.