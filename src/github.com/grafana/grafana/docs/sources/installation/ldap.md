+++
title = "LDAP Authentication"
description = "Grafana LDAP Authentication Guide "
keywords = ["grafana", "configuration", "documentation", "ldap"]
type = "docs"
[menu.docs]
name = "LDAP Authentication"
identifier = "ldap"
parent = "admin"
weight = 2
+++

# LDAP 认证

DataVision支持LDAP集成功能。可以支持LDAP的认证和授权。

## 配置
您可以通过修改配置文件中的LDAP路径 [main config file]({{< relref "configuration.md#auth-ldap" >}}) 来起用LDAP功能。(默认为: `/etc/datavision/ldap.toml`).

### 配置示例

```toml
# Set to true to log user information returned from LDAP
verbose_logging = false

[[servers]]
# Ldap server host (specify multiple hosts space separated)
host = "127.0.0.1"
# Default port is 389 or 636 if use_ssl = true
port = 389
# Set to true if ldap server supports TLS
use_ssl = false
# Set to true if connect ldap server with STARTTLS pattern (create connection in insecure, then upgrade to secure connection with TLS)
start_tls = false
# set to true if you want to skip ssl cert validation
ssl_skip_verify = false
# set to the path to your root CA certificate or leave unset to use system defaults
# root_ca_cert = "/path/to/certificate.crt"

# Search user bind dn
bind_dn = "cn=admin,dc=grafana,dc=org"
# Search user bind password
# If the password contains # or ; you have to wrap it with triple quotes. Ex """#password;"""
bind_password = 'grafana'

# User search filter, for example "(cn=%s)" or "(sAMAccountName=%s)" or "(uid=%s)"
search_filter = "(cn=%s)"

# An array of base dns to search through
search_base_dns = ["dc=datavision,dc=org"]

# In POSIX LDAP schemas, without memberOf attribute a secondary query must be made for groups.
# This is done by enabling group_search_filter below. You must also set member_of= "cn"
# in [servers.attributes] below.

## Group search filter, to retrieve the groups of which the user is a member (only set if memberOf attribute is not available)
# group_search_filter = "(&(objectClass=posixGroup)(memberUid=%s))"
## An array of the base DNs to search through for groups. Typically uses ou=groups
# group_search_base_dns = ["ou=groups,dc=grafana,dc=org"]

# Specify names of the ldap attributes your ldap uses
[servers.attributes]
name = "givenName"
surname = "sn"
username = "cn"
member_of = "memberOf"
email =  "email"

# Map ldap groups to datavision org roles
[[servers.group_mappings]]
group_dn = "cn=admins,dc=datavision,dc=org"
org_role = "Admin"
# The datavision organization database id, optional, if left out the default org (id 1) will be used.  Setting this allows for multiple group_dn's to be assigned to the same org_role provided the org_id differs
# org_id = 1

[[servers.group_mappings]]
group_dn = "cn=users,dc=datavision,dc=org"
org_role = "Editor"

[[servers.group_mappings]]
# If you want to match all (or no ldap groups) then you can use wildcard
group_dn = "*"
org_role = "Viewer"

```

## LDAP 和 DataVision 组织角色同步

### 映射
通过 `[[servers.group_mappings]]` 您可以将LDAP组合DataVision组织进行对应.  用户每次登陆时会进行数据同步. 
