package cis_1_2_26
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    every comArgs in commandArgs {
      not startswith(comArgs,"--request-timeout")
    }
    msg = sprintf("%v Ensure that the --request-timeout argument is set as appropriate (Manual)",[container.name])

}
