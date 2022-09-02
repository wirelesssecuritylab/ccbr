package cis_1_3_4
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-controller-manager"
    some container in input.spec.containers
    commandArgs := container.command
    every comArgs in commandArgs {
          not startswith(comArgs,"--service-account-private-key-file")
    }
    msg = sprintf("%v Ensure that the --service-account-private-key-file argument is set as appropriate (Automated)",[container.name])

}









