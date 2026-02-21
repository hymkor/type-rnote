type-rnote
==========

`type-rnote` is a CLI tool that fetches release notes from a GitHub repository and outputs them to STDOUT in a Markdown-compatible ChangeLog format.

Features
--------

- **Flexible Input:** Supports repository URLs, `owner/repo` strings, or separate arguments.
- **Auto-Detection:** If no repository is specified, it automatically detects the repository information from your local `.git/config`.
- **Revision Support:** Use the `-r` flag to fetch a specific version/tag.

Usage
-----

```bash
# Provide username and repository separately
type-rnote [USERNAME] [REPOSITORY]

# Provide as a single string or URL
type-rnote USERNAME/REPOSITORY
type-rnote https://github.com/USERNAME/REPOSITORY

# Auto-detect from local .git/config (if inside a git repo)
type-rnote

# Fetch a specific revision
type-rnote -r v0.3.0 [USERNAME/REPOSITORY]
```

> [!NOTE]
> Square brackets `[ ]` indicate optional arguments.

Sample Output
-------------

Running the command will produce output formatted like this:

```markdown
Changelog
=========

v0.3.0
------
Jan 7, 2025

- Add the option `-r REVISION`

v0.2.0
------
...
```

Installation
------------

### From Releases

Download the pre-built binary for your platform from the [Releases](https://github.com/hymkor/type-rnote/releases) page and add it to your PATH.

### Via Go

```bash
go install github.com/hymkor/type-rnote@latest
```

### Via Scoop (Windows)

```bash
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install type-rnote
```

Or

```bash
scoop install https://raw.githubusercontent.com/hymkor/type-rnote/master/type-rnote.json
```
