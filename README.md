# Cronus
Cronus is a lightweight, fast command-line interface (CLI) time-tracking tool built in Go. It uses a SQLite database for additional analysis and summary.

Named after the Titan [Cronus](https://en.wikipedia.org/wiki/Cronus) - the maintainer of the cycles of seasons and periods of time. 


***The stream of time devours us all.***

## Features

* **Start Sessions:** Quickly start a timer for any project.
* **Check Status:** View the elapsed time of an actively running session.
* **Stop Sessions:** Stop the timer and automatically calculate the total elapsed time.
* **Local Storage:** All data is safely stored locally on your machine in a SQLite database.

## Installation

Ensure you have [Go installed](https://go.dev/doc/install) on your machine.

Clone the repository and build the binary:

```bash
git clone https://github.com/MasEo9/cronus.git 
cd cronus
go build -o cronus main.go
```

You can then move the `cronus` binary to a directory in your system's `PATH` (e.g., `/usr/local/bin` on Linux/macOS) to use it from anywhere.

## Usage

Cronus uses the `-p` or `--project` flag to identify which project you are tracking.

**Add a project:**
```bash
cronus project add -p "Project"
```

**List all projects:**
```bash
cronus project
```

**Start tracking a project:**
```bash
cronus start -p "Project"
```

**Check the status of an active project:**
```bash
cronus status -p "Project"
```

**Check the status of all active projects:**
```bash
cronus status
```

**Stop tracking and save the session:**
```bash
cronus stop -p "Project"
```

## Database & Manual Querying

Cronus stores all of your time-tracking data in a SQLite database named `cronus.db`. 

By default, this file is created in your operating system's default user configuration directory 
(e.g., `~/.config/cronus.db` on Linux, `~/Library/Application Support/cronus.db` on macOS, or `%AppData%\\cronus.db` on Windows).

### Querying with SQLite3

If you want to view your raw data, export it, or run custom analytics, you can easily query the database directly using the `sqlite3` CLI tool.

1. Open your terminal and connect to the database (replace the path with the actual path to your `cronus.db` based on your OS):
   ```bash
   sqlite3 ~/.config/cronus.db
   ```

2. Format the output to make it easier to read by turning on headers and setting the display mode to column:
   ```sqlite
   sqlite> .headers on
   sqlite> .mode column
   ```

3. Run a SQL query to see all of your tracked sessions:
   ```sqlite
   sqlite> SELECT s.*, p.project_name FROM sessions s LEFT JOIN projects p on p.id = s.project_id;
   ```

**Example Output:**
```text
id  project_name        date        time_start                  time_end                    elapsed_time  hours
--  ------------------  ----------  --------------------------  --------------------------  ------------  ---------
0   Project             2026.04.15  2026-04-15 10:00:00.000000  2026-04-15 11:30:00.000000  5400000000000 1.5
```

**Other Useful Commands:**
* `.tables` - List all tables in the database.
* `.schema sessions` - View the structure of the `sessions` table.
* `.quit` - Exit the sqlite3 prompt.
