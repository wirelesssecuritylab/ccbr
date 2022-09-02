package cis_1_2_8

import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    argus_judement(container.command) = "allow"
    msg = sprintf("%v Ensure that the --authorization-mode argument includes Node (Automated)",[container.name])

}


argus_judement(commandArgs) = "allow"{
   every comArgs in commandArgs {
           not startswith(comArgs,"--authorization-mode")
   }
}else = "allow"{

   not args_contains_string(commandArgs,"--authorization-mode","Node")
}else = "notallow"

args_contains_string(commandArgs,argskey,argsvalue){
 some commandstring in commandArgs
 #--enable-access-plugins=AlwaysPulImages,hello
 command_Arr :=split(commandstring,"=")
 command_Arr[0] == argskey
 some name in split(command_Arr[1],",")
 trim(name, " ") == argsvalue
}

