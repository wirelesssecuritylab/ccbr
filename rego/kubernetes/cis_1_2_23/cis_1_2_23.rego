package cis_1_2_23
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    every comArgs in commandArgs {
      not startswith(comArgs,"--audit-log-max")
    }
    msg = sprintf("%v Ensure that the --audit-log-max age argument is set to 30 or as appropriate (Automated)",[container.name])

}


