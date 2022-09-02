package cis_1_2_24
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    every comArgs in commandArgs {
      not startswith(comArgs,"--audit-log-maxbackup")
    }
    msg = sprintf("%v Ensure that the --audit-log-maxbackup argument is set to 10 or as appropriate (Automated)",[container.name])

}

