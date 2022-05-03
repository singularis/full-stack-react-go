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
    https://www.linode.com/docs/guides/enabling-https-using-certbot-with-nginx-on-ubuntu/
2) Setup Node
    sudo apt install nodejs npm
    cd parcel_blueprint
    npm init -y
    npm install --save parcel-bundler
    npm install --save @babel/core
    npm install --save @babel/plugin-proposal-class-properties
    npm install --save @babel/plugin-transform-runtime
    npm install --save react@next react-dom@next
    npm install --save styled-components
3) Install mongoDB
    Setup user and admin users
    Add auth
4) Install Golang