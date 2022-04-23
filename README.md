There is amateur full-stack project 

Front - React

Back - Go

Infrastructure k8s

Bellow all steps that need to be done in order to bring it to life. 

Bash
1) Init scrip: init-setup.sh

Kubernetes
1) Setup NGINX:
    k apply -f nginx-service+deployment.yaml 
    #Install the NGINX Ingress Controller
    helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
    helm repo update
    helm install ingress-nginx ingress-nginx/ingress-nginx
    Setup Domain name and LB
    Setup HTTPS certificates
2) Setup Node

