
kubectl replace -f godocker_deployment.yaml

kubectl delete deployment githubnginx
kubectl create -f github_nginx.yaml 
kubectl expose deployment githubnginx --type="LoadBalancer"

kubectl delete svc githubnginx



kubectl expose rs githubnginx-3350899980  --port=8080 --target-port=8080

kubectl exec githubnginx-4228723764-i16bf -c nginx -i -t -- bash -il
