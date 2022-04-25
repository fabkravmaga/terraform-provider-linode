---
layout: "linode"
page_title: "Linode: linode_database_mysql"
sidebar_current: "docs-linode-datasource-database-mysql"
description: |-
Provides information about a Linode MySQL Database.
---

# Data Source: linode\_database\_mysql

**NOTICE:** Managed Databases are currently in beta. Ensure `api_version` is set to `v4beta` in order to use this data source.

Provides information about a Linode MySQL Database.

## Example Usage

Get information about a MySQL database:

```hcl
data "linode_database_mysql" "my-db" {
  database_id = 12345
}
```

## Argument Reference

The following arguments are supported:

* `database_id` - The ID of the MySQL database.

## Attributes

The `linode_database_mysql` data source exports the following attributes:

* `allow_list` - A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.

* `ca_cert` - The base64-encoded SSL CA certificate for the Managed Database instance.

* `cluster_size` - The number of Linode Instance nodes deployed to the Managed Database.

* `created` - When this Managed Database was created.

* `encrypted` - Whether the Managed Databases is encrypted.

* `engine` - The Managed Database engine. (e.g. `mysql`)

* `engine_id` - The Managed Database engine in engine/version format. (e.g. `mysql/8.0.26`)

* `host_primary` - The primary host for the Managed Database.

* `host_secondary` - The secondary/private network host for the Managed Database.

* `label` - A unique, user-defined string referring to the Managed Database.

* `region` - The region that hosts this Linode Managed Database.

* `root_password` - The randomly-generated root password for the Managed Database instance.

* `root_username` - The root username for the Managed Database instance.

* `replication_type` - The replication method used for the Managed Database. (`none`, `asynch`, `semi_synch`)

* `ssl_connection` - Whether to require SSL credentials to establish a connection to the Managed Database.

* `status` - The operating status of the Managed Database.

* `type` - The Linode Instance type used for the nodes of the  Managed Database instance.

* `updated` - When this Managed Database was last updated.

* `version` - The Managed Database engine version. (e.g. `v8.0.26`)