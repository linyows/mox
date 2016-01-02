Pox
===

Pox is a very simple mock server as web api.

Description
-----------

It is a mock server simply returns response files.

Usage
-----

```sh
$ pox --root /var/www/pox --protocol json-rpc --delay 2 --loglevel debug
```

use config file:

```sh
$ pox --config /etc/pox/pox.conf
```

Install
-------

To install, use `go get`:

```sh
$ go get -d github.com/linyows/pox
```

Contribution
------------

1. Fork ([https://github.com/linyows/pox/fork](https://github.com/linyows/pox/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

Author
------

[linyows](https://github.com/linyows)
