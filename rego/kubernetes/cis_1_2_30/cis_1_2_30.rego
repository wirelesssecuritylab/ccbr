package cis_1_2_29
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    is_argument_not_set(container)
    msg = sprintf("%v Ensure that the --tls-cert-file and --tls-private-key-file arguments are set as appropriate (Automated)",[container.name])
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--tls-cert-file")
    }
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--tls-private-key-file")
    }
}



