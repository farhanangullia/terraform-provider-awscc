---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_lightsail_load_balancer_tls_certificate Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate
---

# awscc_lightsail_load_balancer_tls_certificate (Resource)

Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `certificate_domain_name` (String) The domain name (e.g., example.com ) for your SSL/TLS certificate.
- `certificate_name` (String) The SSL/TLS certificate name.
- `load_balancer_name` (String) The name of your load balancer.

### Optional

- `certificate_alternative_names` (Set of String) An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.
- `is_attached` (Boolean) When true, the SSL/TLS certificate is attached to the Lightsail load balancer.

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `load_balancer_tls_certificate_arn` (String)
- `status` (String) The validation status of the SSL/TLS certificate.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_lightsail_load_balancer_tls_certificate.example <resource ID>
```