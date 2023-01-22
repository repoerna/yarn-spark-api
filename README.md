# YARN SPARK MONITORING CLI
## Overview 
This is an cli tools to call Spark REST API that using YARN as resource manager.

## How to use
### Get All YARN Application
```sh
./build/yarnspark-cli yarn-cluster-apps spark-application-job <resouce-manager-url> -u ampid -s RUNNING --applicationTypes SPARK
```
### Get All Spark Job in All YARN Application
```sh
./build/yarnspark-cli yarn-cluster-apps spark-application-job <resouce-manager-url> -u ampid -s RUNNING --applicationTypes SPARK --status RUNNING
```
## Help
```sh
./build/yarnspark-cli yarn-cluster-apps -h
./build/yarnspark-cli yarn-cluster-apps spark-application-job -h
```