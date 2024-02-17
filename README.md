# :trumpet: bugle

Send messages from command line. Written in Go. Run from the command line. Designed as a tool for distributed systems.

## Features

- Terse: 1-liner from command line
- Portable: Work with executable, go package, Docker image, Apptainer compatible
- Lightweight: (~10MB Docker Image; ~6MB Compressed Image; ~6MB Binary)

## Installation

### Golang
```bash
go install github.com/debeshmandal/bugle
bugle --body="Hello!" --dry-run
```

### Manual Build (Golang)
```bash
git clone https://github.com/debeshmandal/bugle.git
cd bugle
go install .
bugle --body="Hello!" --dry-run
```

### Docker (DockerHub)
```bash
docker run debeshmandal/bugle --body="Hello!" --dry-run
```

### Docker (GitHub Container Registry)
```bash
docker run ghcr.io/debeshmandal/bugle --body="Hello!" --dry-run
```

### Manual Build (Docker)
```bash
git clone https://github.com/debeshmandal/bugle.git
cd bugle
docker build -t bugle .
docker run bugle --body="Hello!" --dry-run
```

### Apptainer (DockerHub)
```bash
apptainer run docker://debeshmandal/bugle --body="Hello!" --dry-run
```

## Usage

### 1. Set the SMTP environment variables
```bash
export BUGLE_SMTP_SERVER=<smtp.example.com>
export BUGLE_USERNAME=<username>
export BUGLE_PASSWORD=<password>
```
### 2. Set Sender, Recipient, Subject and Body as CLI arguments
```bash
bugle --sender="name@email.com" --recipient="name@email.com" --subject="Subject" --body="Message"
```
## Incomplete Features
- Add `--html` for HTML body
- Add `--attachments` for file attachments
- Multiple recipients
- Authentication
- Add `--slack` to post to slack
- Other builds that aren't linux/amd64
- Download executable via `curl` or `wget` as GH release artifact

## License
Bugle is licensed under the GPL-3.0 license. See the LICENSE file for details.