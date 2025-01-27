# Table: aws_docdb_orderable_db_instance_options

This table shows data for Amazon DocumentDB Orderable DB Instance Options.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_OrderableDBInstanceOption.html

The composite primary key for this table is (**account_id**, **region**, **db_instance_class**, **engine**, **engine_version**).

## Relations

This table depends on [aws_docdb_engine_versions](aws_docdb_engine_versions.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|availability_zones|`json`|
|db_instance_class (PK)|`utf8`|
|engine (PK)|`utf8`|
|engine_version (PK)|`utf8`|
|license_model|`utf8`|
|storage_type|`utf8`|
|vpc|`bool`|