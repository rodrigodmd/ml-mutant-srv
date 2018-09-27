# Deployment scripts

The deploymeny uses the pushed docker images:
https://hub.docker.com/r/rodrigodmd/ml-mutant-srv/tags/

To deploy in kubernates run:

    ./deploy-kubernetes.sh

(NOT TESTED) To deploy in openshift run:

    ./deploy-openshift.sh

![Kubernates 1](../doc/k8_1.png)
![Kubernates 2](../doc/k8_2.png)