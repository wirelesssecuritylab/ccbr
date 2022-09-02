package cis_1_2_25
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    every comArgs in commandArgs {
      not startswith(comArgs,"--audit-log-maxsize")
    }
    msg = sprintf("%v Ensure that the --audit-log-maxsize argument is set to 100 or as appropriate (Automated)",[container.name])

}
