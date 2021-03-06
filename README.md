# Inner Dev Loop with Skaffold

## Prequisites

1. You have [Docker](https://docs.docker.com/get-docker/), [Helm](https://helm.sh/docs/intro/install/) and [Skaffold](https://skaffold.dev/docs/install/) installed.
2. You have access to a kubernetes cluster and kubectl in the correct context.
3. Your cluster has a secret called `regcred` of type `kubernetes.io/dockerconfigjson` to authenticate your cluster against your docker registry.
4. You are locally authenticated against your docker register via `docker login`.

## Running the example

1. Set an environment variable corresponding to your repository for example `REPO=gcr.io/your-google-project-id`
2. From within the project directory, run `skaffold dev --default-repo $REPO`
3. Observe the logs from Skaffold are successful in deploying the code.
4. Make a change to the log message in `main.go`.
5. Observe the output in the Skaffold logs or use kubectl to inspect the pod itself.
6. You've hot-reloaded some code in a cluster!
