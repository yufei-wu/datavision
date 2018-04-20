+++
title = "DataVision CLI"
description = "Guide to using datavision-cli"
keywords = ["grafana", "cli", "datavision-cli", "command line interface"]
type = "docs"
[menu.docs]
parent = "admin"
weight = 8
+++

# DataVision CLI

DataVision cli is a small executable that is bundled with datavision-server and is supposed to be executed on the same machine datavision-server is running on.

## Plugins

The CLI allows you to install, upgrade and manage your plugins on the machine it is running on.
You can find more information about how to install and manage your plugins in the
[plugins page]({{< relref "plugins/installation.md" >}}).

## Admin

> This feature is only available in Grafana 4.1 and above.

To show all admin commands:
`datavision-cli admin`

### Reset admin password

You can reset the password for the admin user using the CLI. The use case for this command is when you have lost the admin password.

`datavision-cli admin reset-admin-password ...`

If running the command returns this error:

> Could not find config defaults, make sure homepath command line parameter is set or working directory is homepath

then there are two flags that can be used to set homepath and the config file path.

`datavision-cli admin reset-admin-password --homepath "/usr/share/grafana" newpass`

If you have not lost the admin password then it is better to set in the Grafana UI. If you need to set the password in a script then the [Grafana API](http://docs.grafana.org/http_api/user/#change-password) can be used. Here is an example using curl with basic auth:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
  "oldPassword": "admin",
  "newPassword": "newpass",
  "confirmNew": "newpass"
}' http://admin:admin@<your_grafana_host>:3000/api/user/password
```
