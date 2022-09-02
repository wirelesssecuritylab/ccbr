package cis_2_2
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "etcd"
    some i
    some container in input.spec.containers
    container.command[i] == "--client-cert-auth=false"
    msg = sprintf("%v Ensure that the --client-cert-auth argument is set to true (Automated)",[container.name])

}

