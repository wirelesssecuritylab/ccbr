package cis_2_3
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "etcd"
    some i
    some container in input.spec.containers
    container.command[i] == "--auto-tls=true"
    msg = sprintf("%v Ensure that the --auto-tls argument is set to false (Automated)",[container.name])

}

