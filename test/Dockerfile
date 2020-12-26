FROM debian:buster

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update -y \
    && apt-get install systemd systemd-sysv -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
       /lib/systemd/system/multi-user.target.wants/* \
       /etc/systemd/system/*.wants/* \
       /lib/systemd/system/local-fs.target.wants/* \
       /lib/systemd/system/sockets.target.wants/*udev* \
       /lib/systemd/system/sockets.target.wants/*initctl* \
       /lib/systemd/system/sysinit.target.wants/systemd-tmpfiles-setup* \
       /lib/systemd/system/systemd-update-utmp*

COPY testing-unit.service /etc/systemd/system/
RUN chmod 664 /etc/systemd/system/testing-unit.service

ENTRYPOINT ["/lib/systemd/systemd"]
