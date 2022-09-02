package cis_1_4_2
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-scheduler"
    some container in input.spec.containers
    every name in container.command {
        name != "--bind-address=127.0.0.1"
    }
    msg = sprintf("%v Ensure that the --bind-address argument is set to 127.0.0.1 (Automated)",[container.name])

}

