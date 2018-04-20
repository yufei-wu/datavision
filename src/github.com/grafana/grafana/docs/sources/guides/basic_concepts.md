+++
title = "Basic Concepts"
description = "Grafana intro and concept guide"
keywords = ["grafana", "intro", "guide", "concepts"]
type = "docs"
[menu.docs]
name = "Basic Concepts"
identifier = "basic_concepts"
parent = "guides"
weight = 2
+++

# 基本概念

本文档是DataVision基本概念的“自下而上”介绍，可以作为熟悉核心功能的起点。

### 数据源

DataVision为您的时间序列数据支持许多不同的存储后端 (数据源)。 每个数据源都有一个特定的查询编辑器，可根据特定数据源公开的功能进行自定义。

以下数据源是系统原生支持的: 
[Graphite]({{< relref "features/datasources/graphite.md" >}}), 
[InfluxDB]({{< relref "features/datasources/influxdb.md" >}}), 
[OpenTSDB]({{< relref "features/datasources/opentsdb.md" >}}), 
[Prometheus]({{< relref "features/datasources/prometheus.md" >}}), 
[Elasticsearch]({{< relref "features/datasources/elasticsearch.md" >}}), 
[CloudWatch]({{< relref "features/datasources/cloudwatch.md" >}}).

每个数据源的查询语言和功能显然非常不同。 您可以将来自多个数据源的数据组合到单个仪表板上，但每个面板都绑定到属于特定组织的特定数据源。

### 组织

DataVision支持多个组织，以支持各种各样的部署模式，可以使用单个DataVision实例为多个不同的组织提供服务。

在很多情况下，DataVision只在系统中部署一个组织。

每个组织有一个或者多个数据源。

所有仪表板都由特定组织拥有。

 > 注意: 请务必记住，大多数度量标准数据库不提供任何类型的每个用户的系列认证。 因此，在DataVision中，数据源和仪表板可供特定组织中的所有用户使用。

有关DataVision用户模型的更多详细信息，请参阅[Admin](/reference/admin/)

### 用户

用户是DataVision中的一个帐户。 用户可以属于一个或多个组织，并且可以通过角色分配不同级别的权限。

DataVision支持多种内部和外部方式供用户进行身份验证。 这些包括来自平台自身的集成数据库、来自外部的SQL服务器或来自外部LDAP服务器。

想要了解更多详情请参阅 [User Auth](/reference/http_api/#users)

### 行

行是仪表板内的逻辑分隔符，用于将面板组合在一起。

行总是分隔为12个“单位”宽。 这些单位会根据浏览器的水平分辨率自动缩放。 您可以通过设置自己的宽度来控制一行中面板的相对宽度。

我们利用这种单位的抽象，能让DataVision在所有小型和巨大的屏幕上都看起来很棒。

 > 注意: 借助MaxDataPoints功能，DataVision可以向您显示完美的数据点数量，无论您的分辨率或时间范围如何。

利用 [Repeating Row functionality](/reference/templating/#utilizing-template-variables-with-repeating-panels-and-repeating-rows) 可以根据所选的模板变量动态创建或删除整个行。

通过单击行标题可以折叠行。 如果保存一个带有折叠式行的仪表板，它将保存在该状态下，并且不会预加载行中的图形，直到该行被展开。

### 面板

面板是DataVision中的基本可视化构建模块。 每个面板都提供了一个查询编辑器（取决于在面板中选择的面板类型），它允许您利用查询编辑器来提取数据并将数据完美可视化显示在面板上。

每个面板都有各种各样的样式和格式选项来让您创建完美的图像。

面板可以在仪表板上拖放和重新排列。 您也可以调整他们的大小。

目前有四种面板类型: 
[Graph](/reference/graph/), 
[Singlestat](/reference/singlestat/), 
[Dashlist](/reference/dashlist/), 
[Table](/reference/table_panel/), 以及 
[Text](/reference/text/).

像这种 [Graph](/reference/graph/) 面板允许您根据需要绘制尽可能多的指标和系列图表化。 其他像 [Singlestat](/reference/singlestat/) 这种面板要求将单个查询缩减为单个数字。 [Dashlist](/reference/dashlist/) 以及 [Text](/reference/text/) 是不连接到任何数据源的特殊面板。

利用面板配置中的 [Dashboard Templating](/reference/templating/) 变量字符串（包括通过查询编辑器配置的数据源查询）可以使面板变得更加动态。

利用 [Repeating Panel](/reference/templating/#utilizing-template-variables-with-repeating-panels-and-repeating-rows) 功能基于已选择的 [Templating Variables](/reference/templating/#utilizing-template-variables-with-repeating-panels-and-repeating-rows) 来动态创建或删除面板。

面板上的时间范围通常是通过 [Dashboard time picker](/reference/timerange/) 设置的，但这可以利用 [Panel specific time overrides](/reference/timerange/#panel-time-overrides-timeshift) 来覆盖。

面板 (或者整个仪表盘) 能够很简单的 [Shared](/reference/sharing/) 。 你可以发送一个链接给登录到系统上的人。您可以使用 [Snapshot](/reference/sharing/#snapshots) 功能将当前正在查看的所有数据编码为静态和交互式JSON文档; 它比发送截图更好！


### 查询编辑器

查询编辑器可以让你利用数据源提供的功能，查询它包含的指标。

使用查询编辑器在时间序列数据库中构建一个或多个查询（针对一个或多个系列），面板将立即更新，让您可以实时有效地探索数据，并为该特定面板构建完美的查询。

您可以利用 [Template variables](/reference/templating/) 在查询编辑器中。 这提供了一种基于Dashboard上选定的模板变量动态探索数据的强大方法。

DataVision允许您通过查询编辑器所在的行来引用查询。 如果您将第二个查询添加到图形中，则只需输入#A即可引用第一个查询。 这为构建复合查询提供了一种简单方便的方法。

### 仪表盘

仪表板是面板聚集在一起的地方。 仪表板可以被认为是组织和排列成一个或多个行的一组或多组面板的集合。

仪表板的时间段可以利用在仪表板的右上方的 [Dashboard time picker](/reference/timerange/) 来控制。

仪表盘能够利用 [Templating](/reference/templating/) 功能使仪表盘更具动态性和互动性。

仪表盘能够利用 [Annotations](/reference/annotations/) 来跨面板显示事件数据。这有助于将面板中的时间序列数据与其他事件相关联。

仪表盘 (或一个特定的面板) 能够很简单的 [Shared](/reference/sharing/) 。 你可以发送一个链接给登录到系统上的人。您可以使用 [Snapshot](/reference/sharing/#snapshots) 功能将当前正在查看的所有数据编码为静态和交互式JSON文档; 它比发送截图更好！

仪表板可以被标记，仪表板选择器可以快速，可搜索地访问特定组织中的所有仪表板。
