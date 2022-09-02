package cis_1_2_18
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some i
    some container in input.spec.containers
    startswith(container.command[i], "--insecure-bind-address")
    msg = sprintf("%v Ensure that the --insecure-bind-address argument is not set (Automated)",[container.name])

}
