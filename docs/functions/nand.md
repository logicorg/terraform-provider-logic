---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nand function - logic"
subcategory: ""
description: |-
  Return the NAND between two values.
---

# function: nand

Return the NAND between two values.

## Example Usage

```terraform
locals {
  a = true
  b = true
  c = false

  # nand(a, b) => false
  nand_ab = provider::logic::nand(local.a, local.b)

  # nand(a, c) => true
  nand_ac = provider::logic::nand(local.a, local.c)
}
```

## Signature

<!-- signature generated by tfplugindocs -->
```text
nand(first bool, second bool) bool
```

## Arguments

<!-- arguments generated by tfplugindocs -->
1. `first` (Boolean, Nullable) The first value.
1. `second` (Boolean, Nullable) The second value.
