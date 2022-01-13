
# Kuknos Cli

Kuknos cli is a command line tool for managing tool and specialy git management for Kuknos Development team.


## Features

- Create and manage git branches
- Create merge request with command line


## Documentation

Full document comming soon...


## Installation

Download the deb package from [Releases](https://github.com/vahidid/kuknos-cli/releases) and install the package:

```bash
  dpkg -i kuknos.deb
```
    
## Usage

After you installed the package, you can access it with run:

```bash

kuknos --help

```

In the first step you can initialize with:

```bash

kuknos git init

```

Note: the current version(v0.0.1) just support git module!

for create new branch use ``start`` operation with the namespace:

```bash

kuknos git feature start <feature-name>

```

#### Namespaces

- feature
- bugfix
- hotfix
- improvement
- release

#### Operations

- start
- finish
- push
- pull

#### Defination of each namespace

- ``feature``: create new branch from ``develop`` branch and push it to remote server with create a merge request
- ``bugfix`` : create new branch from ``develop`` branch and push it to remote server with create a merge request
- ``hotfix`` : create new branch from ``master`` branch and push it to remote server with create a merge request
- ``improvement`` : create new branch from ``develop`` branch and push it to remote server with create a merge request
- ``release`` : create new branch from ``develop`` branch and create a tag with push and create merge request.

Note: namespace will not merged with target branch. They just create new merge request when they finished.

Note: namespace branches will have a prefix that you define it in ``init`` command.

Note: The remote name is hardcoded to ``origin`` for now :)
## Feedback

If you have any feedback, please reach out to us at v.hasani@kuknos.ir or create an issue!

