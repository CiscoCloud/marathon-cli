# marathon-cli
A golang cli to manage apps and groups in Marathon

##Usage

```
COMMANDS:
   info		Gets information about the marathon cluster
   leader	Gets the current leader
   rmleader	Forces current leader to abdicate
   rmapp	Deletes an app
   lsapp	List all apps or lists provided apps in the arguments
   ping		Test Marathon connection
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host "localhost:8080"	Marathon host (default localhost:8080) [$MARATHON_HOST]
   --format "json"		Output format (json)
   --help, -h			show help
   --version, -v		print the version
```

##Examples
###Server Information
####Ping the Marathon endpoint

```
./marathon-cli --host host-01:8080 ping

INFO[0000] Ping successful, recieved 'pong' from host-02:8080 

```

####Get information about a Marathon cluster
```
./marathon --host host-01:8080 info

{
   "Code": 200,
   "leader": "host-01:8080",
   "frameworkId": "20150212-005751-3034589101-5050-27381-0000",
   "http_config": {
      "http_port": 8080,
      "https_port": 8443
   },
   "marathon_config": {
      "checkpoint": true,
      "executor": "//cmd",
      "failover_timeout": 604800,
      "ha": true,
      "hostname": "host-01",
      "local_port_min": 10000,
      "local_port_max": 20000,
      "master": "zk://zookeeper.service.consul:2181/mesos",
      "mesos_role": "",
      "mesos-user": "",
      "reconciliation_initial_delay": 15000,
      "reconciliation_interval": 300000,
      "task_launch_timeout": 300000
   },
   "name": "marathon",
   "version": "0.8.0",
   "zookeeper_config": {
      "zk": "zk://zookeeper.service.consul:2181/marathon",
      "zk_future_timeout": {
         "duration": 10
      },
      "zk_hosts": "zookeeper.service.consul:2181",
      "zk_path": "/marathon",
      "zk_state": "",
      "zk_timeout": 10
   }
}

```


####Get the leader of a marathon cluster

```
 ./marathon-cli --host host-01:8080 leader

{
   "Code": 200,
   "leader": "host-02:8080"
}
```

####Force the leader step down
```
./marathon-cli --host host-01:8080 rmleader

{
   "Code": 200,
   "message": "Leadership abdicted"
}
```

Confirm new leader:

```
./marathon-cli --host host-01:8080 leader
{
   "Code": 200,
   "leader": "host-01:8080"
}
```

###Managing Applications
Applications are identified by their "id" field.

####List all apps running on the cluster

To shorten the command line, we can set the `MARATHON_HOST` environment variable:

```
export MARATHON_HOST=host:8080
```

```
./marathon-cli lsapp
{
   "Code": 200,
   "apps": [
      {
         "id": "/hello-rails",
         "cmd": "cd hello \u0026\u0026 bundle install \u0026\u0026 bundle exec unicorn -p $PORT",
         "cpus": 1,
         "env": {
            "RAILS_ENV": "production"
         },
         "instances": 1,
         "mem": 100,
         "ports": [
            10001
         ],

     .. <output omitted>
```

####List a single app running on the cluster

```
./marathon-cli lsapp hello-rails
{
   "Code": 200,
   "app": {
      "id": "/hello-rails",
      "cmd": "cd hello \u0026\u0026 bundle install \u0026\u0026 bundle exec unicorn -p $PORT",
      "cpus": 1,
      "env": {
         "RAILS_ENV": "production"
      },
      "instances": 1,
      "mem": 100,
      "ports": [

      ... <output omitted>

}

```

####Delete an application
This will terminate and delete an application

```
./marathon-cli rmapp hello-rails
INFO[0000] Application deleted: hello-rails
```


####Launch an application form a json description file.
Use a standard Marathon application configuration file to launch an application. The `examples/` directory contains sample files.


```
./marathon-cli app --file examples/rails.json
INFO[0000] Application deployed app: hello-rails
{
   "Code": 201,
   "version": "2015-02-16T14:27:55.572Z"
}
```

Refer to the [Marathon Documentation](https://mesosphere.github.io/marathon/docs/rest-api.html#post-/v2/apps) for more configuration examples.


 
