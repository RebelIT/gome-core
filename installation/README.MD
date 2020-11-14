## Manual Installation
* build: 
* service file goes into `/usr/lib/systemd/system/`
* env file goes into `/etc/gome/core/`
* database is located `/etc/gome/core/db`
* open FW: `firewall-cmd --zone=public --permanent --add-port 6660/tcp`
* save FW: `firewall-cmd --reload`