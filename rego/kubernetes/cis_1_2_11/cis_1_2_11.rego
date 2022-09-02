package cis_1_2_11

import future.keywords.in
import future.keywords.every

violation[msg]{
    input.metadata.labels.component == "kube-apiserver"
    some container in input.spec.containers
    commandArgs := container.command
    argus_judement(container.command) = "allow"
    msg = sprintf("%v Ensure that the admission control plugin AlwaysAdmit is not set (Automated)",[container.name])

}
argus_judement(commandArgs) = "allow"{
   every comArgs in commandArgs {
           not startswith(comArgs,"--enable-access-plugins")
   }
}else = "allow"{

     args_contains_string(commandArgs,"--enable-access-plugins","AlwaysAdmit")
}else = "notallow"

args_contains_string(commandArgs,argskey,argsvalue){
 some commandstring in commandArgs
 #--enable-access-plugins=AlwaysPulImages,hello
 command_Arr :=split(commandstring,"=")
 command_Arr[0] == argskey
 some name in split(command_Arr[1],",")
 trim(name, " ") == argsvalue
}
