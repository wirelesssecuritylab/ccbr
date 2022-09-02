package cis_2_6
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "etcd"
    some i
    some container in input.spec.containers
    container.command[i] == "--peer-client-cert-auth=false"
    msg = sprintf("%v Ensure that the --peer-auto-tls argument is set to true (Automated)",[container.name])

}

