variable "mutually_exclusive" {
  description = "Mutually exclusive variable"
  type = object({
    first  = optional(bool)
    second = optional(bool)
  })
  default = {
    first = true
  }

  validation {
    condition     = provider::logic::xor(var.mutually_exclusive.first, var.mutually_exclusive.second)
    error_message = "You must set one and only one of `first` or `second`."
  }
}
