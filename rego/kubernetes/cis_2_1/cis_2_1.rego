package cis_2_1
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "etcd"
    some container in input.spec.containers
    is_argument_not_set(container)
    msg = sprintf("%v Ensure that the --cert-file and --key-file arguments are set as appropriate (Automated)",[container.name])
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--cert-file")
    }
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--key-file")
    }
}


