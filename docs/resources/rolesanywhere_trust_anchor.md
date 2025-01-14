---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_rolesanywhere_trust_anchor Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::RolesAnywhere::TrustAnchor Resource Type.
---

# awscc_rolesanywhere_trust_anchor (Resource)

Definition of AWS::RolesAnywhere::TrustAnchor Resource Type.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `source` (Attributes) (see [below for nested schema](#nestedatt--source))

### Optional

- `enabled` (Boolean)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `trust_anchor_arn` (String)
- `trust_anchor_id` (String)

<a id="nestedatt--source"></a>
### Nested Schema for `source`

Optional:

- `source_data` (Attributes) (see [below for nested schema](#nestedatt--source--source_data))
- `source_type` (String)

<a id="nestedatt--source--source_data"></a>
### Nested Schema for `source.source_data`

Optional:

- `acm_pca_arn` (String)
- `x509_certificate_data` (String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_rolesanywhere_trust_anchor.example <resource ID>
```
