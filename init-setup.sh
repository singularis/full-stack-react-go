#!/bin/bash
#Install soft
config_file=full-stack-kubeconfig.yaml 
sudo apt install git wget vim bash-completion curl elinks
bash -c "$(curl -fsSL https://raw.githubusercontent.com/ohmybash/oh-my-bash/master/tools/install.sh)"
source /usr/share/bash-completion/bash_completion
#k8s completion
echo 'source <(kubectl completion bash)' >>~/.bashrc
echo 'alias k=kubectl' >>~/.bashrc
echo 'complete -F __start_kubectl k' >>~/.bashrc
#Import Linode ssh
cd ../linode || exit
chmod go-r "$config_file"
export KUBECONFIG=$config_file
#Install HELM
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
git config --global user.email "singularis314@gmail.com"
git config --global user.name "Yevhenii Parasochka"