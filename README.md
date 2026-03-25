# ssh-manager

A terminal-based SSH connection manager written in Go. Store your SSH profiles in a local JSON file and connect to any of them from a numbered menu.

## Requirements

- Go 1.21+
- SSH targets using password authentication

## Setup

1. Clone the repo and navigate into the `manager/` directory:
   ```bash
   cd manager
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create your `profile.json` from the sample:
   ```bash
   cp sample_profile.json profile.json
   ```
   Then edit `profile.json` with your actual hosts and credentials.

## Usage

```bash
go run main.go
```

You'll see a numbered list of your profiles. Enter the number to connect:

```
SSH Manager v0.0.1
Date & Time: 2026-03-25 14:30:00

1. Production
2. Dev Server

Select a profile: 1
Connecting to 192.168.1.100:22 ...
```

Once connected, you get a full interactive terminal session. Type `exit` to disconnect and return to your local shell.

## Profile Configuration

Profiles are stored in `manager/profile.json` (gitignored). Use `sample_profile.json` as a reference:

```json
[
  {
    "name": "Production",
    "host": "192.168.1.100",
    "port": 22,
    "username": "root",
    "password": "your-password",
    "keypath": "",
    "isActive": true
  }
]
```

| Field | Description |
|---|---|
| `name` | Display name shown in the menu |
| `host` | Hostname or IP address |
| `port` | SSH port (usually `22`) |
| `username` | SSH login username |
| `password` | SSH password |
| `keypath` | Path to private key (not yet implemented) |
| `isActive` | Reserved for future filtering |

You can add as many profiles as you need — just add more objects to the array.

## Build

To compile a standalone executable:

```bash
go build -o ssh-manager.exe
```

## Current Limitations

- Password authentication only (`keypath` field is defined but not yet wired up)
- Host key verification is disabled (`InsecureIgnoreHostKey`)
- No encryption on stored credentials — keep `profile.json` secure

## License

Apache License 2.0
