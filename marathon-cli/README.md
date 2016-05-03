# shipped-utils
MIsc utilities and scripts related shipped.

## Utility Marathon API
This folder includes main.go which work based out of two senario

### 1) Search Details based out of Host Name & Port Number.
Need to Build and Run the project and run the following command.

```
./shipped-utils -u={usernmae} -p={password} -url={marathon api endpoint} -host={hostname} -port={port} 
```

### Output Reposponse.

```

App Id : /ihbe--newproj--gotest5--79d066
project_name :  newproj
project_id :  7fe682da-0200-11e6-abf7-0242ac11367e
env_name :  ihbe
env_id :  c42bfb78-0201-11e6-8bc7-0242ac110004
service_name :  gotest5
service_id :  290e8bc3-05b6-11e6-89da-0242ac110003
host : hostname
port : portnum
```

### 2) Search Details based out of App Id.
No Need to Build and Run again just run the following command.

```
./shipped-utils -u={usernmae} -p={password} -url={marathon api endpoint} -id={application id deployed on the given endpoint} 
```

### Output Reposponse.

```

App Id : /ihbe--newproj--gotest5--79d066
project_name :  newproj
project_id :  7fe682da-0200-11e6-abf7-0242ac11367e
env_name :  ihbe
env_id :  c42bfb78-0201-11e6-8bc7-0242ac110004
service_name :  gotest5
service_id :  290e8bc3-05b6-11e6-89da-0242ac110003
[ 
Host_0 : shipped-tx3-worker-103.node.consul
Port_0 : 11288
]
```
