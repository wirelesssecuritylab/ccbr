package cis_1_2_3
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some i
    some container in input.spec.containers
    startswith(container.command[i], "--token-auth-file")
    msg = sprintf("%v Ensure that the --token-auth-file parameter is not set (Automated)",[container.name])

}





