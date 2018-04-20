+++
title = "Installing Plugins"
type = "docs"
[menu.docs]
parent = "plugins"
weight = 1
+++

# DataVision 插件

DataVision不仅支持数据源插件，还支持面板插件和应用程序插件。 将面板作为插件可以轻松创建和添加任何类型的面板，以显示数据或改进您最喜爱的仪表板。 应用程序插件在DataVision中是一项新功能，它可以将数据源，面板，仪表板和DataVision页面捆绑到一个解决方案中。

搜索插件，请查看 [Plugin Repository](https://www.dataconnect.com/plugins).

# 安装插件

安装插件的最简单方法是使用与DataVision捆绑在一起的CLI工具datavision-cli。 在修改插件之后发生任何修改之前，DataVision服务器需要重新启动。

### DataVision 插件目录

在Linux系统上，datavision-cli会将datavision插件目录设置为 `/var/lib/datavision/plugins`. 通过指定--pluginsDir参数可以覆盖datavision-cli默认的插件目录。 在Windows系统中，必须为每个调用指定此参数。

### datavision-cli 命令

列出所有插件
```bash
datavision-cli plugins list-remote
```

安装插件的最新版本
```bash
datavision-cli plugins install <plugin-id>
```

安装一个特定版本的插件
```bash
datavision-cli plugins install <plugin-id> <version>
```

列出已经安装的插件
```bash
datavision-cli plugins ls
```

更新所有已经安装的插件
```bash
datavision-cli plugins update-all
```

更新一个插件
```bash
datavision-cli plugins update <plugin-id>
```

移除一个插件
```bash
datavision-cli plugins remove <plugin-id>
```

### 手动安装插件

如果您的DataVision服务器无法访问互联网，则该插件必须下载并手动复制到您的DataVision服务器中。

插件的下载链接是下面这种类型:

`https://www.dataconnect.com/api/plugins/<plugin id>/versions/<version number>/download`

您可以通过 `--pluginUrl` 选项来设定一个本地目录进行插件的安装.
```bash
datavision-cli --pluginUrl https://nexus.company.com/grafana/plugins/<plugin-id>-<plugin-version>.zip plugins install <plugin-id>
```

To manually install a Plugin via the Grafana.com API:

1. 找到你想下载的插件，插件ID可以在dataconnect.com插件页面的安装选项卡上找到。 在这个例子中插件ID为 `jdbranham-diagram-panel`:

    {{< imgbox img="/img/docs/installation-tab.png" caption="Installation Tab" >}}

2. 使用DataVision API通过 url `https://grafana.com/api/plugins/<plugin id from step 1>` 来查找插件. 比如: https://grafana.com/api/plugins/jdbranham-diagram-panel 应该返回如下:
    ```bash
    {
      "id": 145,
      "typeId": 3,
      "typeName": "Panel",
      "typeCode": "panel",
      "slug": "jdbranham-diagram-panel",
      "name": "Diagram",
      "description": "Diagram panel for datavision",
    ...
    ```

3. 找到下载链接:
    ```bash
    {
       "rel": "download",
       "href": "/plugins/jdbranham-diagram-panel/versions/1.4.0/download"
    }
    ```

4. 下载插件通过URL `https://grafana.com/api/plugins/<plugin id from step 1>/versions/<current version>/download` (比如: https://grafana.com/api/plugins/jdbranham-diagram-panel/versions/1.4.0/download). 解压插件包到 `plugins` 目录.

5. 重启您的DataVision服务器.
