+++
title = "Installing on Windows"
description = "Installing Grafana on Windows"
keywords = ["grafana", "configuration", "documentation", "windows"]
type = "docs"
[menu.docs]
parent = "installation"
weight = 3
+++

# 在 Windows 中安装 DataVision

描述 | 下载
------------ | -------------
windows的最新版本 | [datavision-5.0.4.windows-x64.zip](https://www.dataconnect.com/release/datavision-5.0.4.windows-x64.zip)

从已有的安装中进行升级请参阅 [Upgrading Grafana]({{< relref "installation/upgrading.md" >}}) .

## 配置

解压安装包到 DataVision 的安装目录.  进入
`conf` 目录下，拷贝 `sample.ini` 文件为 `custom.ini`. 你可以修改
`custom.ini` 文件来添加自定义配置, 请不要修改 `defaults.ini` 文件.

系统默认的端口为 `3000`, 在windows上该端口需要额外的权限，您可以通过编辑 `custom.ini` 中的 `http_port`
配置项来将端口改变为类似于 `8080` 的其他端口.

通过CMD来执行bin目录下的 `datavision-server.exe` 来启动整个服务. 如果您想将DataVision作为Windows的一个服务, 请下载
[NSSM](https://nssm.cc/). 使用该工具可以很容易的将DataVision作为一个Windows服务.

更多配置请查看 [configuration options]({{< relref "configuration.md" >}}).
