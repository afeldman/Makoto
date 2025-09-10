# Makoto

Makoto Nishimura was a Japanese biologist.  
In 1928 he designed *Gakutensoku*, a pneumatic writing robot.  
Inspired by this pioneering work, the Makoto tool aims to bring structured,
reproducible package management to **Fanuc Karel projects**.

---

## What is Makoto?

Makoto is a **Karel Package Configurator (KPC)** and lightweight package manager.  
It uses a **local database** (`kpc.db`) to keep track of all your Karel packages,
their metadata, dependencies and conflicts.

Conceptually it is similar to tools like **Cargo (Rust)** or **Go Modules**, but
adapted to the needs of **Fanuc Karel development**.

---

## Features

- Manage packages via a declarative `kpc.toml` file.
- Register and store packages in a local BoltDB/Storm database.
- Inspect package metadata from the CLI:
  - `makoto package author` – show authors
  - `makoto package version` – show version(s)
  - `makoto package source` – show repository URL/commit
  - `makoto package requirements` – list dependencies
  - `makoto package conflicts` – list known conflicts
  - `makoto package prefix` – show installation prefix
- Remove packages again (`makoto package reject`).
- Cross-platform builds via [GoReleaser](https://goreleaser.com).

---

## Installation

Clone and build manually:

```bash
git clone https://github.com/afeldman/Makoto.git
cd Makoto
go build -o makoto
./makoto version
```

Or use GoReleaser to create binaries for Linux/Windows/macOS.

## Usage

Examples:

```bash
# Register a package into the local DB
makoto register ./examples/hello_world/kpc.toml

# Show all packages currently known
makoto all

# Inspect details
makoto package version hello_world
makoto package author hello_world
makoto package source hello_world
makoto package requirements hello_world
makoto package conflicts hello_world
makoto package prefix hello_world

# Remove again
makoto package reject hello_world 0.1.0
```

## Example kpc.toml
```toml
kpc_version = "0.2.0"

name = "hello_world"
description = "Minimal example Karel project"
version = "0.1.0"
url = "https://example.com/hello_world"

[[authors]]
name  = "Max Mustermann"
email = "max@example.com"

[source]
url = "https://github.com/yourname/hello_world.git"
tag = "v0.1.0"

[[requirements]]
name    = "motion_lib"
version = ">=1.2.0"

[[conflicts]]
name     = "old_motion_lib"
versions = ["<1.0.0"]

prefix   = "/fr/robonet"
main     = "hello.kl"
srcs     = ["hello.kl", "util.kl"]
forms    = ["hello.frm"]
dicts    = ["hello.dtx"]
```

## Configuration

Makoto stores its config and database under:

- ~/.config/makoto/makoto.toml
- ~/.config/makoto/kpc.db

You can override the config path with:
```bash
makoto --config /path/to/makoto.toml
```

## Roadmap

- Git integration: automatically clone sources (git clone, git fetch, git checkout tag/rev).
- Build helpers for Fanuc Karel projects.
- Export dependency graph.

## License

MIT – © Anton Feldmann