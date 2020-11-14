## Manual Installation
* build: `GOOS=linux go build -o gome-core .`
* service file goes into `/usr/lib/systemd/system/`
* env file goes into `/etc/gome/core/`
* database is located `/etc/gome/core/db`
* reload systemd: `systemctl daemon-reload`
* open FW: `firewall-cmd --zone=public --permanent --add-port 6660/tcp`
* save FW: `firewall-cmd --reload`