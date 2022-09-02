package cis_1_2_10

import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    argus_judement(container.command) = "allow"
    msg = sprintf("%v Ensure that the admission control plugin EventRateLimit is set (Manual)",[container.name])

}
argus_judement(commandArgs) = "allow"{
   every comArgs in commandArgs {
           not startswith(comArgs,"--enable-access-plugins")
   }
}else = "allow"{

   not args_contains_string(commandArgs,"--enable-access-plugins","EventRateLimit")
}else = "notallow"

args_contains_string(commandArgs,argskey,argsvalue){
 some commandstring in commandArgs
 #--enable-access-plugins=AlwaysPulImages,hello
 command_Arr :=split(commandstring,"=")
 command_Arr[0] == argskey
 some name in split(command_Arr[1],",")
 trim(name, " ") == argsvalue
}
