package cis_1_2_31
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    is_argument_not_set(container)
    msg = sprintf("%v Ensure that the --client-ca-file argument is set as appropriate (Automated)",[container.name])
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--client-ca-file")
    }
}




