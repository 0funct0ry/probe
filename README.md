# Probe

**Probe** is a command-line utility written in Go for testing container behavior. 

## Usage

```bash
./probe [command]
```

### Available Commands

- **`clock`**: Logs the current time at regular intervals.
- **`completion`**: Generates the autocompletion script for the specified shell.
- **`help`**: Displays help information about any command.
- **`serve`**: Starts a server.


## Serve Command

The `serve` command starts a server with additional subcommands to perform specific actions.

### Usage

```bash
probe serve [flags]
probe serve [command]
```

### Available Subcommands

- **`env`**: Runs an HTTP server that provides OS environment variables in JSON or text format.
- **`proxy`**: Proxies requests to a remote URL.

### Serve Flags

- `-h, --help`: Displays help information for the `serve` command.

For more details about a specific command, use:

```bash
probe [command] --help
```

## Installation

Clone the repository and build the tool using Go:

```bash
git clone https://github.com/yourusername/probe-cli.git
cd probe-cli
go build -o probe
```

You can then run the tool using:

```bash
./probe
```
