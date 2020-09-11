## How to run

From your local computer

```bash
cd cmd/
go build
docker build . -f Dockerfile -t <your dockerhub account>/gocql-test:<image tag of your choice>
docker push <your dockerhub account>/gocql-test:<image tag of your choice>
kubectl run gocql -it --rm --image=<your dockerhub account>/gocql-test:<image tag of your choice> -- bash
```

Now you are inside the container on k8s

```
/app/cmd --cosmos-db-url <your cosmos db account name with cassandra api>.cassandra.cosmos.azure.com --username <your cosmos db account name> --password <your cosmos db accoutn key>
```
