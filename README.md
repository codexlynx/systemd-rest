## systemd-rest / A minimal HTTP REST interface for systemd.
![](https://github.com/codexlynx/systemd-rest/workflows/CI/badge.svg) [![AUR](https://img.shields.io/aur/license/yaourt.svg)](blob/master/LICENSE)

### Architecture:
This service communicates with systemd via dbus (IPC). ```http <-> systemd-rest <-> dbus <-> systemd```

### Features:
* Manage units (start. stop, status)
* List units

### Configuration:
Currently configured using environment variables.

| Variable | Description | Default |
| -- | -- | -- |
| __PORT__ | HTTP server listen port | 6789 |
| __ADDRESS__ | HTTP server listen address | 127.0.0.1 |
| __MODE__ | Service mode, can be `release` or `debug` | release |

### Development:

* Launch tests:
> $ make test

* Launch service for tests development:
> $ make test ARGS="follow"
> $ PORT=7777 CONTAINER_ID=[Testing container id] pytest -sv

* Generate build:
> $ make

### License:
> GPL (GNU General Public License) 3.0

More info: [here](LICENSE)

### About
This service was created by: __@codexlynx__.

* Twitter: [https://twitter.com/codexlynx](https://twitter.com/codexlynx)
* GitHub: [https://github.com/codexlynx](https://github.com/codexlynx)
