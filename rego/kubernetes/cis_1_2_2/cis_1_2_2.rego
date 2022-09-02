package cis_1_2_2
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some i
    some container in input.spec.containers
    startswith(container.command[i], "--basic-auth-file")
    msg = sprintf("%v Ensure that the --basic-auth-file parameter is not set (Automated)",[container.name])

}
