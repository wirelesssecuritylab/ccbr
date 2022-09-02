package cis_2_4
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "etcd"
    some container in input.spec.containers
    is_argument_not_set(container)
    msg = sprintf("%v Ensure that the --peer-cert-file and --peer-key-file arguments are set as appropriate (Automated)",[container.name])
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--peer-cert-file")
    }
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--peer-key-file")
    }
}
