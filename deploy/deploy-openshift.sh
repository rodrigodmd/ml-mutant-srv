#!/bin/bash
oc apply -f ocp/consul-imagestream.yaml
oc apply -f ocp/consul-deploymentconfig.yaml
oc apply -f ocp/consul-service.yaml

oc apply -f ocp/elasticsearch-imagestream.yaml
oc apply -f ocp/elasticsearch-deploymentconfig.yaml
oc apply -f ocp/elasticsearch-service.yaml

oc apply -f ocp/api-imagestream.yaml
oc apply -f ocp/api-deploymentconfig.yaml
oc apply -f ocp/api-service.yaml

oc apply -f ocp/gateway-imagestream.yaml
oc apply -f ocp/gateway-deploymentconfig.yaml

oc apply -f ocp/mutant-imagestream.yaml
oc apply -f ocp/mutant-deploymentconfig.yaml

oc apply -f ocp/stats-imagestream.yaml
oc apply -f ocp/stats-deploymentconfig.yaml
