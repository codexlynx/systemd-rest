## systemd-rest / A minimal HTTP REST interface for systemd.
[![](https://github.com/codexlynx/systemd-rest/workflows/CI/badge.svg)](https://github.com/codexlynx/systemd-rest/actions) [![AUR](https://img.shields.io/aur/license/yaourt.svg)](LICENSE)

### Architecture:
This service communicates with systemd via dbus (IPC). ```http <-> systemd-rest <-> dbus <-> systemd```

### Features:
* Manage units (start, stop, status)
* List units
* Read unit journal

### Configuration:
Currently configured using environment variables.

| Variable | Description | Default |
| -- | -- | -- |
| __PORT__ | HTTP service listen port | 6789 |
| __ADDRESS__ | HTTP service listen address | 127.0.0.1 |
| __MODE__ | Service mode, can be `release` or `debug` | release |

### Development:

* Launch tests:
```
$ make test
```

* Launch service for tests development:
```
$ make test ARGS="follow"
$ PORT=7777 pytest -sv
```

* Generate build:
```
$ make
```

### License:
> GPL (GNU General Public License) 3.0

More info: [here](LICENSE)

### About
This service was created by: __@codexlynx__.

* Twitter: [https://twitter.com/codexlynx](https://twitter.com/codexlynx)
* GitHub: [https://github.com/codexlynx](https://github.com/codexlynx)
