<h1>Mox <img src="https://dl.dropboxusercontent.com/u/759617/mox/black-atom-hi.png" width="27" height="27"></h1>

Mox is a very simple mock server as web api.

[![Travis](https://img.shields.io/travis/linyows/mox.svg?style=flat-square)][travis]
[![GitHub release](http://img.shields.io/github/release/linyows/mox.svg?style=flat-square)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[travis]: https://travis-ci.org/linyows/mox
[release]: https://github.com/linyows/mox/releases
[license]: https://github.com/linyows/mox/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/linyows/mox

Description
-----------

It is a mock server simply returns response files.

Usage
-----

```sh
$ mox --root /var/www/mox --protocol json-rpc --delay 2 --loglevel debug
```

use config file:

```sh
$ mox --config /etc/mox/mox.conf
```

Install
-------

To install, use `go get`:

```sh
$ go get -d github.com/linyows/mox
```

Contribution
------------

1. Fork ([https://github.com/linyows/mox/fork](https://github.com/linyows/mox/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

Author
------

[linyows](https://github.com/linyows)
