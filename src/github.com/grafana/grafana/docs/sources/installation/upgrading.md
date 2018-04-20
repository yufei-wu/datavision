+++
title = "Upgrading"
description = "Upgrading Grafana guide"
keywords = ["grafana", "configuration", "documentation", "upgrade"]
type = "docs"
[menu.docs]
name = "Upgrading"
identifier = "upgrading"
parent = "installation"
weight = 10
+++

# 升级 DataVision

如果您需要升级当前版本的DataVision到最新版本，请参考如下步骤

## 数据库备份

在升级之前您最好备份一下当前数据库，以保证您可以回退到前一个版本。 在新版本的启动过程中，如果新版本有数据库变更，系统会自动进行数据库合并. 

#### sqlite 备份

如果系统使用 sqlite 数据库，那么仅仅简单备份存储于 `/var/lib/grafana/grafana.db` 文件夹下的 `grafana.db` 文件即可。
如果您不确定系统使用的数据库类型，请检查您的系统配置文件。数据库一般会存放在 `<datavision_install_dir>/data` 文件夹下。

#### mysql 备份

```bash
备份:
> mysqldump -u root -p[root_password] [datavision] > datavision_backup.sql

恢复:
> mysql -u root -p datavision < datavision_backup.sql
```

#### postgres 备份

```bash
备份:
> pg_dump datavision > datavision_backup

恢复:
> psql datavision < datavision_backup
```

### Ubuntu / Debian

#### 从DEB包进行升级

下载最新版的deb安装包，执行 `dpkg -i` 命令即可.

#### 从TAR包进行升级

备份conf目录下的custom.ini文件，下载最新版的TAR安装包，解压安装包覆盖已有的安装文件，将原来 conf 目录下的 custom.ini 覆盖到安装目录中。

### Centos / RHEL

直接下载RPM包，执行 `rpm -i` 进行RPM包的安装。

### Windows

备份conf目录下的custom.ini文件，下载最新版的ZIP安装包，解压安装包覆盖已有的安装文件，将原来 conf 目录下的 custom.ini 覆盖到安装目录中。

