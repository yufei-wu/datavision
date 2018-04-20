+++
title = "Installing on Debian / Ubuntu"
description = "Install guide for Grafana"
keywords = ["grafana", "installation", "documentation"]
type = "docs"
aliases = ["/installation/installation/debian"]
[menu.docs]
name = "Installing on Ubuntu / Debian"
identifier = "debian"
parent = "installation"
weight = 1
+++

# 在 Debian / Ubuntu 系统中安装 DataVision

描述 | 下载
------------ | -------------
Debian-based Linux稳定版 | [datavision_5.0.4_amd64.deb](https://www.dataconnect.com/release/datavision_5.0.4_amd64.deb)

## 安装稳定版

```bash
sudo apt-get install -y adduser libfontconfig
sudo dpkg -i datavision_5.0.4_amd64.deb
```

## 安装详情

- 安装二进制文件到 `/usr/sbin/datavision-server`
- 安装 Init.d 脚本到 `/etc/init.d/datavision-server`
- 创建默认文件到 `/etc/default/datavision-server`
- 安装配置文件到 `/etc/datavision/datavision.ini`
- 安装 systemd 服务的名称为 `datavision-server.service`
- 系统默认的日志文件为 `/var/log/datavision/datavision.log`
- 系统默认 sqlite3 数据库存放在 `/var/lib/datavision/datavision.db`
- 安装 HTML/JS/CSS 以及其他文件到 `/usr/share/datavision`

## 启动服务 (init.d service)

通过运行以下命令:

```bash
sudo service datavision-server start
```

使用 `datavision` 用户来启动 `datavision-server` 进程, 该用户在系统安装期间自动进行创建。系统默认的HTTP端口为 `3000`, 默认的用户以及用户组为 `admin`.

配置DataVision系统开机启动:

```bash
sudo update-rc.d datavision-server defaults
```

## 启动服务 (通过 systemd)

通过 systemd 来启动服务:

```bash
systemctl daemon-reload
systemctl start datavision-server
systemctl status datavision-server
```

配置DataVision系统开机启动:

```bash
sudo systemctl enable datavision-server.service
```

## 环境文件

systemd 服务和 init.d 脚本都使用位于
`/etc/sysconfig/datavision-server` 来启动后台服务。

### 日志

缺省情况下，系统将日志写入到 `/var/log/datavision`

### 数据库

缺省情况下 sqlite3 数据库位于
`/var/lib/datavision/datavision.db` 目录下. 在升级前请备份该数据库. 您也可以使用MySQL 或者 Postgres 作为平台的数据库, 详情请参考 [the configuration page]({{< relref "configuration.md#database" >}}).

## 配置

系统配置文件存放于 `/etc/datavision/datavision.ini`.  配置详情请参考
[Configuration]({{< relref "configuration.md" >}})。

### 添加数据源

- [Graphite]({{< relref "features/datasources/graphite.md" >}})
- [InfluxDB]({{< relref "features/datasources/influxdb.md" >}})
- [OpenTSDB]({{< relref "features/datasources/opentsdb.md" >}})
- [Prometheus]({{< relref "features/datasources/prometheus.md" >}})

### 服务端图片渲染

服务端PNG图片渲染是一个可选的并且非常有用的一个功能。

```bash
yum install fontconfig
yum install freetype*
yum install urw-fonts
```

## 从二进制Tar包进行安装

下载Tar安装包，将安装包解压到安装目录。

通过添加 `custom.ini` 文件到
`conf` 目录下来创建自定义的属性配置.

通过执行 `./bin/datavision-server web` 来启动系统. 
