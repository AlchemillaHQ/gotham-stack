# Gotham Stack Template

This is a simple template which contains a todo app for bootstrapping a new Go web application project. It ues:

* [Go](https://golang.org) as the programming language.
* [HTML Templates](https://golang.org/pkg/html/template/) as the template engine.
* [HTMX](https://htmx.org) as the front-end framework.
* [SQLite](https://www.sqlite.org) or [PostgreSQL](https://www.postgresql.org) as the database.

## Getting Started

1. Clone this repository.

```bash
git clone https://github.com/AlchemillaHQ/gotham-stack.git
```

2. Install the dependencies.

```bash
go mod download
```

3. Create a configuration file.

```bash
cp config.example.json config.json
```

4. Run the application.

```bash
go run main.go
```

## Testing & Building

For both testing and building we use a Makefile. To run the tests:

```bash
make test
```

To build the application:

```bash
make build
```

It should create binary files for some common platforms in the dist/ directory.