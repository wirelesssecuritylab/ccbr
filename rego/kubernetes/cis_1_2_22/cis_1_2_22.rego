package cis_1_2_22
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    every comArgs in commandArgs {
      not startswith(comArgs,"--audit-log-path")
    }
    msg = sprintf("%v Ensure that the --audit-log-path argument is set (Automated)",[container.name])

}




