package cis_1_3_1
import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-controller-manager"
    some container in input.spec.containers
    is_argument_not_set(container)
    msg = sprintf("%v Ensure that the --terminated-pod-gc-threshold argument is set as appropriate (Manual)",[container.name])

}
is_argument_not_set(container) {
    every name in container.command{
          not startswith(name,"--terminated-pod-gc-threshold")
    }
}

