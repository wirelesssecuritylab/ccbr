package cis_2_7
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "etcd"
    some container in input.spec.containers
    is_argument_not_set(container)
    msg = sprintf("%v Ensure that the --trusted-ca-file arguments are set as appropriate (Manual)",[container.name])
}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--trusted-ca-file")
    }
}




#"=/etc/kubernetes/pki/etcd/ca.crt"
