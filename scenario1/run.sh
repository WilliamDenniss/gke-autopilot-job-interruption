#!/bin/bash

kubectl apply -f shortjob.yaml
kubectl apply -f shortjob2.yaml
kubectl apply -f shortjob3.yaml
kubectl apply -f shortjob4.yaml

kubectl apply -f longjob.yaml