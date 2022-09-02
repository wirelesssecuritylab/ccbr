package cis_1_3_2
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-controller-manager"
    some i
    some container in input.spec.containers
    container.command[i] != "--profiling=false"
    msg = sprintf("%v Ensure that the --profiling argument is set to false (Automated)",[container.name])

}

