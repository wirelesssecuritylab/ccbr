package cis_1_2_1
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    every name in container.command {
        name != "--anonymous-auth=false"
    }
    msg = sprintf("%v Ensure that the --anonymous-auth argument is set to false (Manual)",[container.name])

}
