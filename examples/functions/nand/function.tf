locals {
  a = true
  b = true
  c = false

  # nand(a, b) => false
  nand_ab = provider::logic::nand(local.a, local.b)

  # nand(a, c) => true
  nand_ac = provider::logic::nand(local.a, local.c)
}
