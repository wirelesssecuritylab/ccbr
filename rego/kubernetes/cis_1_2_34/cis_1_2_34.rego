package cis_1_2_34
import future.keywords.in
import future.keywords.every
violation[msg]{
  input.metadata.labels.component == "kube-apiserver"
  some container in input.spec.containers
  is_argument_not_set(container)
  msg = sprintf("%v Ensure that encryption providers are appropriately configured (Manual)",[container.name])

}

is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--encryption-provider-config")
    }
}
