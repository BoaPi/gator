# Gator

Gator is a command line tool to aggregate rss feeds on the terminal.

## Prerequisites

- go version >= 1.23.4
- postgresql >= 15

## Installation

- first clone the repository
- use `go install` from within the repository
  ```cmd
  go install
  ```

## Configuration

In your home directory create a `.gatorconfig.json` file with the postgresql url like so

```json
{
  "db_url": "url-to-your-postgresql?sslmode=disable"
}
```

Be aware to add `?sslmode=disable`

## Commands

**register** - adds a new user

```cmd
gator register <user-name>
```

**login** - logging in as a registered user

```cmd
gator login <user-name>
```

**reset** - reset the complete database

```cmd
gator reset
```
