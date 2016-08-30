

docker run -d -p 8080:8080 --name go_tutorial gcr.io/$PROJECT_ID/godocker:v8






docker run --rm -v "$(pwd):/src" -v /var/run/docker.sock:/var/run/docker.sock centurylink/golang-builder gcr.io/jntlserv0/godocker:v4


gcloud docker push gcr.io/$PROJECT_ID/godocker:v10

 gcloud container clusters create hello-world

kubectl run godocker --image=gcr.io/$PROJECT_ID/godocker:v4 --port=8080


kubectl expose deployment godocker --type="LoadBalancer"

kubectl scale deployment godocker --replicas=4


gcloud compute disks create --project "jntlserv0" --zone "europe-west1-c" --size 200GB mongo-disk



docker run --rm -v "$(pwd):/src" -v /var/run/docker.sock:/var/run/docker.sock centurylink/golang-builder gcr.io/jntlserv0/godocker:v10
gcloud docker push gcr.io/$PROJECT_ID/godocker:v10


kubectl create -f godeploiment.yaml

kubectl set image deployment/godocker godocker=gcr.io/$PROJECT_ID/godocker:v10


docker run -d --link mongo:mongo my_container

#kubectl label pods <pod-name> new-label=awesome

kubectl expose pod godocker --port=8080 --type="LoadBalancer" --name=frontend

kubectl expose deployment godocker --type="LoadBalancer" --target-port=8080  --port=80


kubectl replace --force -f db-service.yml 



kubectl delete service,deployment godocker

gcloud container clusters delete hello-world


gsutil ls

gsutil rm -r gs://artifacts.<$PROJECT_ID>.appspot.com/

gcloud compute firewall-rules list
