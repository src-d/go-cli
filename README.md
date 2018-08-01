# go-cli

**This is currently experimental. API, package name and import path might change any time.**

A thin wrapper around common libraries used in our CLI apps (`jessevdk/go-flags`, `src-d/go-log`, `pprof`) to reduce boilerplate code and help in being more homogeneous with respect how our CLI work and look like.

It provides:
- Struct tags to specify command names and descriptions (see below).
- Default version subcommand.
- Flags and environment variables to setup logging with src-d/go-log.
- Flags and environment variables to setup a http/pprof endpoint.
- Signal handling.

For further details, look at `doc.go`.

## License

Apache License Version 2.0, see LICENSE(LICENSE).
