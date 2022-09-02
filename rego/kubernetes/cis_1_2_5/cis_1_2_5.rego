package cis_1_2_5
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    is_argument_not_set(container)
    msg = sprintf("%v Ensure that the --kubelet-client-certificate and --kubelet-client-key arguments are set as appropriate (Automated)",[container.name])
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--kubelet-client-certificate")
    }
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--kubelet-client-key")
    }
}
