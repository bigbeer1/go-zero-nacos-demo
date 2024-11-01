### 1、预装Gcc docker环境  还有docker-compose

```

#### 步骤1: 部署redis集群
$ docker-compose -f docker-compose-redis.yml up -d

#### 步骤2：启动redis集群
$ docker exec -it redis-1 redis-cli --cluster create 172.20.99.11:6381 172.20.99.12:6382 172.20.99.13:6383 172.20.99.14:6384 172.20.99.15:6385 172.20.99.16:6386 --cluster-replicas 1

#### 步骤3: 部署程序
$ docker-compose -f docker-compose-env.yml up -d

#### 步骤4: 部署程序
$ docker-compose up -d

```


```

$ docker exec -it mysql mysql -uroot -p
##输入密码：PXDNA999999
$ use mysql;
$ update user set host='%' where user='root';
$ FLUSH PRIVILEGES;

```