# terraform-provider-hash-sum


Data provider to hashsum strings like:

main.tf:

```tf

terraform {
  required_providers {
    hashsum = {
      version = "0.2"
      source  = "github.com/snakeice/hash-sum"
    }
  }
}

provider "hashsum" {

}

data "hashsum" "foo1" {
  data = "fooa-1"
}
data "hashsum" "foo2" {
  data = "fooz-2"
}
data "hashsum" "foo3" {
  data = "foox-3"
}

data "hashsum" "foo4" {
  data = "fooy-4"
}

output "foo1" {
  value = {
    foo1    = data.hashsum.foo1.sum
    foo1Mod = data.hashsum.foo1.sum % 4
    foo2    = data.hashsum.foo2.sum
    foo2Mod = data.hashsum.foo2.sum % 4
    foo3    = data.hashsum.foo3.sum
    foo3Mod = data.hashsum.foo3.sum % 4
    foo4    = data.hashsum.foo4.sum
    foo4Mod = data.hashsum.foo4.sum % 4
  }
}
```
output:

```tfs
Changes to Outputs:
  + foo1 = {
      + foo1    = 1686100410
      + foo1Mod = 2
      + foo2    = 1482346890
      + foo2Mod = 2
      + foo3    = 2008161051
      + foo3Mod = 3
      + foo4    = 1686100410
      + foo4Mod = 2
    }
```
