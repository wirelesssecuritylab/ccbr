package cis_1_2_28
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    every comArgs in commandArgs {
       not startswith(comArgs,"--service-account-key-file")
    }
    msg = sprintf("%v Ensure that the --service-account-key-file argument is set as appropriate (Automated)",[container.name])

}



