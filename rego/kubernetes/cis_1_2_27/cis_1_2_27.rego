package cis_1_2_27
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some i
    some container in input.spec.containers
    container.command[i] == "--service-account-lookup=false"
    msg = sprintf("%v Ensure that the --service-account-lookup argument is set to true (Automated)",[container.name])

}
