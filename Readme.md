# Dockit 🐳 - GitHub Repository Docker Image Builder

## Overview

Dockit is a powerful and simple CLI tool that allows you to build Docker images directly from GitHub repositories with ease. Whether you're a developer, DevOps engineer, or just looking to quickly containerize a project, Dockit simplifies the process of building Docker images from remote repositories.

## Features ✨

- 🚀 Quick Docker image building from GitHub repositories
- 🔍 Automatic Dockerfile detection in repository subfolders
- 🧹 Automatic cleanup of temporary cloned repositories
- 📝 Detailed logging and image information display
- 🌐 Support for specifying branches
- 🏷️ Custom image tagging

## Prerequisites

Before using Dockit, ensure you have the following installed:

- Go (1.21 or later)
- Docker
- Git

## Installation

### Option 1: Build from Source

1. Clone the repository:
```bash
git clone https://github.com/sundaram2021/dockit.git
cd dockit
```

2. Build the application:
```bash
go mod tidy
go build -o dockit
```

### Option 2: Go Install

```bash
go install github.com/sundaram2021/dockit@latest
```

## Usage

### Basic Usage

```bash
# Build Docker image from a GitHub repository
dockit https://github.com/example/repo

# Build image from a specific branch
dockit --url https://github.com/example/repo --branch develop

# Custom image tag
dockit --url https://github.com/example/repo --tag my-custom-image

# Verbose logging
dockit --url https://github.com/example/repo --verbose
```

### Command Line Flags

| Flag        | Description                         | Default Value |
|-------------|-------------------------------------|---------------|
| `--url`     | GitHub repository URL (required)    | None          |
| `--branch`  | Branch to clone                     | `main`        |
| `--tag`     | Custom Docker image tag             | Repo name     |
| `--output`  | Custom output directory             | Temp directory|
| `--verbose` | Enable verbose logging              | `false`       |

## Example

### Building a Node simple Application

```bash
dockit https://github.com/trulymittal/docker --branch master --tag latest
```

## How It Works

1. Clone the specified GitHub repository
2. Automatically detect Dockerfile in repository
3. Build Docker image using detected Dockerfile
4. Log detailed image information
5. Clean up temporary files

## Troubleshooting

- Ensure Docker is installed and running
- Check repository URL and branch name
- Use `--verbose` flag for detailed logs
- Verify Dockerfile exists in the repository

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Your Name - [@jhsundaram](https://twitter.com/your_twitter)

Project Link: [https://github.com/sundaram2021/dockit](https://github.com/sundaram2021/dockit)

---

**Happy Dockerizing! 🐳🚀**