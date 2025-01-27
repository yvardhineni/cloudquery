# Table: aws_elasticbeanstalk_configuration_settings

This table shows data for AWS Elastic Beanstalk Configuration Settings.

https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ConfigurationSettingsDescription.html

The composite primary key for this table is (**environment_arn**, **solution_stack_name**, **application_arn**).

## Relations

This table depends on [aws_elasticbeanstalk_environments](aws_elasticbeanstalk_environments.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|environment_arn (PK)|`utf8`|
|application_name|`utf8`|
|date_created|`timestamp[us, tz=UTC]`|
|date_updated|`timestamp[us, tz=UTC]`|
|deployment_status|`utf8`|
|description|`utf8`|
|environment_name|`utf8`|
|option_settings|`json`|
|platform_arn|`utf8`|
|solution_stack_name (PK)|`utf8`|
|template_name|`utf8`|
|application_arn (PK)|`utf8`|