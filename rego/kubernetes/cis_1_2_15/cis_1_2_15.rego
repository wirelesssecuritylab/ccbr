package cis_1_2_15

import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    args_contains_string(commandArgs,"--disable-admission-plugins","NamespaceLifecycle")
    msg = sprintf("%v Ensure that the admission control plugin NamespaceLifecycle is set (Automated)",[container.name])

}

args_contains_string(commandArgs,argskey,argsvalue){
 some commandstring in commandArgs
 #--enable-access-plugins=AlwaysPulImages,hello
 command_Arr :=split(commandstring,"=")
 command_Arr[0] == argskey
 some name in split(command_Arr[1],",")
 trim(name, " ") == argsvalue
}

