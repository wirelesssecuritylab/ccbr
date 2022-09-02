package cis_1_3_3
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-controller-manager"
    some container in input.spec.containers
    every name in container.command {
         name != "--use-service-account-credentials=true"
    }
    msg = sprintf("%v Ensure that the --use-service-account-credentials argument is set to true (Automated)",[container.name])

}

##"--use-service-account-credentials=true"




