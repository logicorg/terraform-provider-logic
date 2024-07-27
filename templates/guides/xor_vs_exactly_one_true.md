---
page_title: "XOR Versus Exactly One True"
---

# XOR Versus Exactly One True
This page will guide you toward choosing the correct function between `provider::logic::xor` and `provider::logic::exactly_one_true`.

Although the two functions may look similar, and they behave exactly the same if you pass a list of length two (`2`) to `exactly_one_true`, the two functions
have a subtle difference in their behavior. Indeed, the `exactly_one_true` which may naively be thinked of as a multi-argument `xor` function, have a different
behaviour than a combination of `xor` functions, but let's go through an example to understand the difference.

```hcl
locals {
    # this will be false
    xor_result = provider::logic::exactly_one_true([true, true, true])

    # this will be true
    xor_result = provider::logic::xor(true, provider::logic::xor(true, true))
}

```

As shown in the example above, the `exactly_one_true` function will return `false`, while the `xor` function will return `true`.

In fact, a multi-argument `xor` (such as the one built in the example by composing multiple `xor` functions) will return `true` for any odd number of
`true` values, and false otherwise. The `exactly_one_true` function on the other hand will return `true` if one and only one `true` value is found in the
supplied list. This is the reason why the `exactly_one_true` function is called like that, to distinghuish it from the `xor` function.

When writing your Terraform code, you should keep in mind this difference and choose the correct function for your use case:
- If you want to guarantee mutual exclusivity between two values, use either `xor` or `exactly_one_true`
- If you want to ensure mutual exclusivity between more than two values, always use `exactly_one_true`
- If you want to build a N-way `xor` function, use a composition of `xor` functions.
