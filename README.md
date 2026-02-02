# GOVA 🏃

A simple CLI tool to visualize your Strava running and trail running statistics directly from your terminal.

## Features

- 🔐 OAuth authentication with Strava
- 📊 Weekly and monthly activity summaries
- 🏃 Supports multiple sport types (Run, Trail Run, Hike, Ride, Weight Training)
- 📈 Displays distance, duration, and elevation gain
- ⏮️ View previous periods with `--last` flag

## Prerequisites

### 1. Go 1.25 or higher

[Install Go](https://go.dev/doc/install) if you haven't already.

### 2. Create a Strava API Application

1. Go to [https://www.strava.com/settings/api](https://www.strava.com/settings/api)
2. Click **"Create an App"** or **"My API Application"**
3. Fill in the application details:
   - **Application Name**: `gova` (or any name you prefer)
   - **Category**: Choose appropriate category
   - **Website**: Your website or `http://localhost`
   - **Authorization Callback Domain**: `localhost`
4. Click **"Create"**
5. You'll receive:
   - **Client ID** (visible)
   - **Client Secret** (keep this private!)

## Installation

### From source

```bash
git clone https://github.com/yourusername/gova.git
cd gova
go build -o gova
```

This creates the `gova` binary in the current directory. You can then:
- Run it directly with `./gova`
- Or install it globally: `sudo cp gova /usr/local/bin/` (then you can use `gova` from anywhere)

### Using go install (recommended)

```bash
# From the cloned repository
go install

# Or directly from GitHub (once published)
go install github.com/yourusername/gova@latest
```

This installs `gova` to your `$GOPATH/bin` (usually `~/go/bin`). Make sure this directory is in your PATH:

```bash
# Add to your ~/.bashrc or ~/.zshrc if not already done
export PATH="$PATH:$HOME/go/bin"
```

## Configuration

Create a `.env` file in the project directory (or set environment variables):

```bash
STRAVA_CLIENT_ID=your_client_id_here
STRAVA_CLIENT_SECRET=your_client_secret_here
AUTH_REDIRECT_URI=http://localhost:8085/exchange_token
```

> **Important**: Never commit your `.env` file with real credentials to version control!

### Example

```bash
cp .env.example .env
# Edit .env with your Strava API credentials
```

## Usage

### 1. Authenticate with Strava

```bash
gova login
```

> **Note**: If you built with `go build -o gova` and didn't install globally, use `./gova login` instead.

This will:
1. Open your browser automatically
2. Redirect to Strava authorization page
3. After approval, redirect back to the app
4. Store your credentials securely in `~/.config/gova/credentials.json`

### 2. View your stats

**Current week:**
```bash
gova week
```

**Last week:**
```bash
gova week --last
```

**Current month:**
```bash
gova month
```

**Last month:**
```bash
gova month --last
```

### Example Output

```
Du 06/01/2026 au 12/01/2026
Activité Run (3): 25.4 km, 2.1h, 120m de dénivelé positif
Activité TrailRun (1): 12.3 km, 1.5h, 450m de dénivelé positif
Activité WeightTraining (2): 0.0 km, 1.2h, 0m de dénivelé positif
```

## Commands

| Command | Description |
|---------|-------------|
| `gova login` | Authenticate with Strava (opens browser) |
| `gova week` | Display current week's stats (Monday to Sunday) |
| `gova week --last` | Display last week's stats |
| `gova month` | Display current month's stats |
| `gova month --last` | Display last month's stats |
| `gova --help` | Show help message |

## Project Structure

The project follows Clean Architecture principles:

```
gova/
├── cmd/                    # CLI commands (Cobra)
│   ├── root.go            # Root command & app initialization
│   ├── login.go           # OAuth login flow
│   ├── week.go            # Weekly stats
│   ├── month.go           # Monthly stats
│   └── me.go              # Profile info (WIP)
├── internal/
│   ├── config/            # Configuration management
│   ├── core/              # Core interfaces (ports)
│   │   ├── ports.go       # ApiClient, OauthClient, TokenProvider
│   │   ├── activity_dto.go
│   │   └── auth_dto.go
│   ├── domain/            # Domain models & business logic
│   │   ├── activity_summary.go
│   │   ├── period.go      # Week/Month calculation
│   │   └── sport_type.go
│   ├── service/           # Application services
│   │   ├── auth_service.go    # Authentication & token management
│   │   └── stat_service.go    # Statistics aggregation
│   └── strava/            # Strava API implementation
│       ├── api_client.go      # Strava API v3 client
│       └── oauth_client.go    # OAuth 2.0 flow
├── main.go
├── go.mod
└── .env.example
```

## Technologies

- **Go 1.25** 
- **[Cobra](https://github.com/spf13/cobra)**
- **Strava API v3**
- **OAuth 2.0**

## Security

- Credentials are stored in `~/.config/gova/credentials.json` with `0600` permissions
- The config directory is created with `0700` permissions
- Access and refresh tokens are managed automatically
- Never commits `.env` file (in `.gitignore`)

---

**Note**: This project is not affiliated with Strava, Inc.
