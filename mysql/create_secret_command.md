#Comando para criar o secret no kubernets

kubectl create secret generic desafio-kysql-k8s --from-literal=password='asd123'

desafio-kysql-k8s => nome da secret
password => chave
asd123 => valor