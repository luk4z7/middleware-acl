# middleware-acl

### Getting started

Clone the repository in folder do you prefer
```bash
cd /var/www
git clone https://github.com/luk4z7/middleware-acl
```

**Execute the file `init.sh` for up the docker containers**

```bash

https://github.com/luk4z7/middleware-acl for the canonical source repository
Lucas Alves 2017 (c) Middleware Access Control Library - Authorization API


           _     _     _ _                                              _
 _ __ ___ (_) __| | __| | | _____      ____ _ _ __ ___        __ _  ___| |
| '_ ` _ \| |/ _` |/ _` | |/ _ \ \ /\ / / _` | '__/ _ \_____ / _` |/ __| |
| | | | | | | (_| | (_| | |  __/\ V  V / (_| | | |  __/_____| (_| | (__| |
|_| |_| |_|_|\__,_|\__,_|_|\___| \_/\_/ \__,_|_|  \___|      \__,_|\___|_|

middleware

DOCKER
Generate new containers ? [ 1 ]
Delete all containers ?   [ 2 ]
Start new build ?         [ 3 ]
Preview the logs ?        [ 4 ]
Install dependencies ?    [ 5 ]
Update dependencies ?     [ 6 ]

```

First step
```bash
Start new build          [ 3 ]
```

Second step
```bash
Generate new containers  [ 1 ]
```

Preview the all logs of containers
```bash
Preview the logs         [ 4 ]
```

Or access the single container
```bash
docker logs api -f
```
```bash
docker logs mongo -f
```

### API REST

**Routers**

In this examples I using jq for pretty the result, for more information view in : [jq](https://stedolan.github.io/jq/)

**List Role**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/roles | jq
```

**Create Role and User**
```bash
curl -H "Content-Type: application/json" -X POST -d '{"user":"alice", "role":"visitante"}' http://127.0.0.1:6060/v1/roles | jq
```

**Get role and its permissions**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/roles/alice | jq
```

**Update role (add permissions)**
```bash
curl -H "Content-Type: application/json" -X PUT -d '{"user":"alice", "permission": "read"}' http://127.0.0.1:6060/v1/roles | jq
```

**Delete role**
```bash
curl -H "Content-Type: application/json" -X DELETE http://127.0.0.1:6060/v1/roles/administrador | jq
```

**List user roles**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/users/alice/roles | jq
```

**Remove role from user**
```bash
curl -H "Content-Type: application/json" -X DELETE http://127.0.0.1:6060/v1/users/alice/roles/visitante3 | jq
```

**Check user permission**
```bash
curl -H "Content-Type: application/json" -X GET http://127.0.0.1:6060/v1/users/alice/resource/data1/permission/read | jq
```

### Tests

Access the container
```bash
docker exec -it api bash
```

```bash
cd /go/src/middleware
```

and execute this:
```bash
./coverage.sh --html
```
and so, the file `html` is generated, only copy the file for the current directory
```bash
cp /tmp/cover314639520/coverage.html .
```

Access the html file inside the project folder outside the container