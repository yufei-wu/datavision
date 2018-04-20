+++
title = "Installing on RPM-based Linux"
description = "Grafana Installation guide for Centos, Fedora, OpenSuse, Redhat."
keywords = ["grafana", "installation", "documentation", "centos", "fedora", "opensuse", "redhat"]
aliases = ["installation/installation/rpm"]
type = "docs"
[menu.docs]
name = "Installing on Centos / Redhat"
identifier = "rpm"
parent = "installation"
weight = 2
+++

# 在 RPM-based Linux (CentOS, Fedora, OpenSuse, RedHat) 的系统中安装 DataVision

Description | Download
------------ | -------------
适合于 CentOS / Fedora / OpenSuse / Redhat Linux 系统的安装包 | [5.0.4 (x86-64 rpm)](https://www.dataconnect.com/release/datavision-5.0.4-1.x86_64.rpm)

## 安装 DataVision

#### CentOS / Fedora / Redhat 系统安装:

```bash
$ sudo yum install initscripts fontconfig
$ sudo rpm -Uvh datavision-5.0.4-1.x86_64.rpm
```

#### OpenSuse 系统安装:

```bash
$ sudo rpm -i --nodeps datavision-5.0.4-1.x86_64.rpm
```

## 安装包详情

- 安装二进制文件到 `/usr/sbin/datavision-server`
- 拷贝 init.d 脚本到 `/etc/init.d/datavision-server`
- 安装默认文件到 `/etc/sysconfig/datavision-server`
- 拷贝配置文件到 `/etc/datavision/datavision.ini`
- 安装 systemd 服务名称为 `datavision-server.service`
- 缺省的日志文件配置为 `/var/log/datavision/datavision.log`
- 缺省的 sqlite3 数据库配置在 `/var/lib/datavision/datavision.db`

## 启动服务 (init.d service)

通过以下脚本启动 DataVision:

```bash
$ sudo service datavision-server start
```

使用 `datavision` 用户来启动 `datavision-server` 进程, 该用户在系统安装期间自动进行创建。系统默认的HTTP端口为 `3000`, 默认的用户以及用户组为 `admin`.

配置DataVision系统开机启动:

```bash
$ sudo /sbin/chkconfig --add  datavision-server
```

## 启动服务 (通过 systemd)

```bash
$ systemctl daemon-reload
$ systemctl start datavision-server
$ systemctl status datavision-server
```

### 开机时启动 systemd 服务

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
