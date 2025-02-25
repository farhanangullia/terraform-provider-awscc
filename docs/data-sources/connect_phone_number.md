---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_connect_phone_number Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Connect::PhoneNumber
---

# awscc_connect_phone_number (Data Source)

Data Source schema for AWS::Connect::PhoneNumber



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `address` (String) The phone number e164 address.
- `country_code` (String) The phone number country code.
- `description` (String) The description of the phone number.
- `phone_number_arn` (String) The phone number ARN
- `prefix` (String) The phone number prefix.
- `tags` (Attributes Set) One or more tags. (see [below for nested schema](#nestedatt--tags))
- `target_arn` (String) The ARN of the target the phone number is claimed to.
- `type` (String) The phone number type, either TOLL_FREE or DID.

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
- `value` (String) The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
