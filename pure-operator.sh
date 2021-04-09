#!/usr/bin/env bash
kubectl delete crd mongodbcommunity.mongodbcommunity.mongodb.com
kubectl delete deployments.apps mongodb-kubernetes-operator