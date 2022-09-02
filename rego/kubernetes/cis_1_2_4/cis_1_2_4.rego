package cis_1_2_4
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some i
    some container in input.spec.containers
    container.command[i] == "--kubelet-https=false"
    msg = sprintf("%v Ensure that the --kubelet-https argument is set to true (Automated)",[container.name])

}





