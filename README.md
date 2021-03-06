Mox
===

<strong>Mox</strong> is a very simple mock server as web api.

[![Travis](https://img.shields.io/travis/linyows/mox.svg?style=for-the-badge)][travis]
[![codecov](https://img.shields.io/codecov/c/github/linyows/mox.svg?style=for-the-badge)][codecov]
[![GitHub release](http://img.shields.io/github/release/linyows/mox.svg?style=for-the-badge)][release]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=for-the-badge)][godocs]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=for-the-badge)][license]

[travis]: https://travis-ci.org/linyows/mox
[codecov]: https://codecov.io/gh/linyows/mox
[release]: https://github.com/linyows/mox/releases
[godocs]: http://godoc.org/github.com/linyows/mox
[license]: https://github.com/linyows/mox/blob/master/LICENSE

Description
-----------

It is a mock server simply returns response files.

Usage
-----

```sh
$ mox --root /var/www/mox --protocol json-rpc --delay 1 --log-level debug
```

use config file:

```sh
$ mox --config /etc/mox/mox.conf
```

### Dockefile:

https://github.com/linyows/mox/blob/master/misc/Dockerfile

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
