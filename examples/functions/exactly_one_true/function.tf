variable "mutually_exclusive" {
  description = "Mutually exclusive object fields."
  type = object({
    first  = optional(string)
    second = optional(string)
    third  = optional(string)
  })
  default = {
    first = "I'm set"
  }

  validation {
    condition = provider::logic::exactly_one_true([
      var.mutually_exclusive.first, # null is falsy
      var.mutually_exclusive.second,
      var.mutually_exclusive.third
    ])
    error_message = "You must set one and only one of `first`, `second` or `third`."
  }
}
