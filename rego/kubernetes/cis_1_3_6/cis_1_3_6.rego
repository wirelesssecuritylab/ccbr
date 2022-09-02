package cis_1_3_6
import future.keywords.in
import future.keywords.every
violation[msg]{
    input.metadata.labels.component == "kube-controller-manager"
    some container in input.spec.containers
    commandArgs := container.command
    argus_judement(commandArgs)= true
    #some comArgs in commandArgs
    #contains(comArgs,"RotateKubeletServerCertificate=true")

    msg = sprintf("%v Ensure that the RotateKubeletServerCertificate argument is set to true (Automated)",[container.name])

}


argus_judement(commandArgs)= true {
   every comArgs in commandArgs {
           not startswith(comArgs,"--feature-gates")
   }
}else = true{
    some comArgs in commandArgs
    startswith(comArgs,"--feature-gates")

   every comArgs in commandArgs {
    not contains(comArgs,"RotateKubeletServerCertificate=true")
   }
}else = false



