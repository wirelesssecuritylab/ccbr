package cis_1_2_33
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    is_argument_not_set(container)
    msg = sprintf("%v Ensure that the --encryption-provider-config argument is set as appropriate (Manual)",[container.name])
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--encryption-provider-config")
    }
}

