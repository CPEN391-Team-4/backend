###### tags: `System`

# Server Configuration

IPv4: `192.53.126.159`
IPv6: `2600:3c01::f03c:92ff:fe2a:56491`

## Setup
Update

```shell=
apt update && apt upgrade
```

Users:

```shell=
apt install sudo zsh tmux
```

```shell=
useradd --create-home \
        --groups sudo \
        --shell /usr/bin/zsh john
```

```shell=
for u in brendon mudit aniket sam; do
    useradd --create-home --groups sudo "${u}"
done
```

Hostname:

Edit ` /etc/hosts`, `/etc/hostname`

Reboot:

```shell=
reboot now
```

## Database

Install:

```shell=
apt install mysql-server mysql-client
systemctl enable --now mysql
```

Secure:

```shell=
mysql_secure_installation
```

## Backend Server

### Go

More recent go.

```shell=
apt install software-properties-common
add-apt-repository ppa:longsleep/golang-backports
apt update
apt-key adv --keyserver keyserver.ubuntu.com --recv-keys F6BC817356A3D45E
apt install golang-go make
```

```shell=
cd /usr/local/share
git clone https://github.com/CPEN391-Team-4/backend.git cpen391-backend
```

Install:

```shell=
GOBIN=/usr/local/bin/ make install
```

### DB

Setup DB:

```shell=
mysql -e"
CREATE USER 'cpen391'@'localhost' IDENTIFIED BY '***********';
CREATE DATABASE cpen391_backend;
GRANT ALL PRIVILEGES ON cpen391_backend.* TO 'cpen391'@'localhost';
"
```

Setup schema:

```shell=
mysql cpen391_backend < db/schema.sql
```

### Service

User:

```shell=
useradd --home /var/lib/cpen391 \
        --create-home \
        --shell /usr/sbin/nologin cpen391
```

```shell=
mkdir -p /var/lib/cpen391/{imagestore,videostore}
chown cpen391:cpen391 /var/lib/cpen391/{imagestore,videostore}
```

Add `/etc/cpen391/key.json`

```shell=
mkdir /etc/cpen391
touch /etc/cpen391/key.json
chmod 600 /etc/cpen391/key.json
chown -R cpen391:cpen391 /etc/cpen391
```

```shell=
mkdir /etc/systemd/system/cpen391-backend.service.d
cp service/cpen391-backend.service /etc/systemd/system/
cp service/cpen391-backend.service.d/env.conf /etc/systemd/system/cpens91-backend.service.d
chmod 660 /etc/systemd/system/cpen391-backend.service.d/env.conf
```

Edit `/etc/systemd/system/cpen391-backend.service.d/env.conf`

Enable:

```shell=
systemctl enable --now cpen391-backend
```

View logs:

```shell=
journalctl -u cpen391-backend
```

Restart service:

```shell=
systemctl restart cpen391-backend
```