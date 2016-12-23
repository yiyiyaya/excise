

###创建服务 etcd
$ETCDCTL mkdir /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662 # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/name "ETCD"
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/description "A Sample ETCD (v2.3.0) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/tags 'etcd,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/metadata '{"displayName":"ETCD","imageUrl":"https://coreos.com/assets/images/media/etcd2-0.png","longDescription":"Managed, highly available etcd clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"https://coreos.com/etcd/docs/latest","supportUrl":"https://coreos.com"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/plan/204F8288-F8D9-4806-8661-EB48D94504B3
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/plan/204F8288-F8D9-4806-8661-EB48D94504B3/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/plan/204F8288-F8D9-4806-8661-EB48D94504B3/description "HA etcd on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/plan/204F8288-F8D9-4806-8661-EB48D94504B3/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/5E397661-1385-464A-8DB7-9C4DF8CC0662/plan/204F8288-F8D9-4806-8661-EB48D94504B3/free false





###创建服务 spark
$ETCDCTL mkdir /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6 # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/name "Spark"
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/description "A Sample Spark (v1.5.2) cluster on Openshift with one worker node"
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/tags 'spark,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/metadata '{"displayName":"Spark","imageUrl":"http://spark.apache.org/images/spark-logo-trademark.png","longDescription":"Managed, highly available Spark clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"http://spark.apache.org/docs/latest/","supportUrl":"http://spark.apache.org"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/CB00754C-11FF-444F-8419-9AA9B1E04970
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/CB00754C-11FF-444F-8419-9AA9B1E04970/name "One_Worker"
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/CB00754C-11FF-444F-8419-9AA9B1E04970/description "Spark on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/CB00754C-11FF-444F-8419-9AA9B1E04970/metadata '{"bullets":["20 GB of Disk","one worker node"],"displayName":"Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/CB00754C-11FF-444F-8419-9AA9B1E04970/free false
###创建套餐2
$ETCDCTL mkdir /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/65242C9B-C266-4F1D-A28B-98B1A2043A84
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/65242C9B-C266-4F1D-A28B-98B1A2043A84/name "Three_Workers"
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/65242C9B-C266-4F1D-A28B-98B1A2043A84/description "HA Spark on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/65242C9B-C266-4F1D-A28B-98B1A2043A84/metadata '{"bullets":["20 GB of Disk","three worker nodes"],"displayName":"High Available" }'
$ETCDCTL set /servicebroker/openshift/catalog/0674A0E3-5EE4-425C-BE43-5A61EB3F52A6/plan/65242C9B-C266-4F1D-A28B-98B1A2043A84/free false




###创建服务 zookeeper
$ETCDCTL mkdir /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416 # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/name "ZooKeeper"
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/description "A Sample ZooKeeper (v3.4.8) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/tags 'zookeeper,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/metadata '{"displayName":"ZooKeeper","imageUrl":"https://zookeeper.apache.org/images/feather_small.gif","longDescription":"Managed, highly available zookeeper clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"https://zookeeper.apache.org/doc/trunk","supportUrl":"https://zookeeper.apache.org/"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan
###创建套餐1 (emptyDir)
$ETCDCTL mkdir /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/221F1338-96C1-4135-A865-9CDA4C12A185
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/221F1338-96C1-4135-A865-9CDA4C12A185/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/221F1338-96C1-4135-A865-9CDA4C12A185/description "HA ZooKeeper on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/221F1338-96C1-4135-A865-9CDA4C12A185/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/221F1338-96C1-4135-A865-9CDA4C12A185/free false
###创建套餐2 (pvc)
$ETCDCTL mkdir /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/dffc3555-eb18-4c7b-ac56-bd326b19dcd0
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/dffc3555-eb18-4c7b-ac56-bd326b19dcd0/name "volumes_standalone"
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/dffc3555-eb18-4c7b-ac56-bd326b19dcd0/description "HA ZooKeeper With Volumes on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/dffc3555-eb18-4c7b-ac56-bd326b19dcd0/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/FA780A47-4AB2-4035-8638-287538B13416/plan/dffc3555-eb18-4c7b-ac56-bd326b19dcd0/free false


###创建服务 RabbitMQ
$ETCDCTL mkdir /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/name "RabbitMQ"
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/description "A Sample RabbitMQ (v3.6.1) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/tags 'rabbitmq,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/metadata '{"displayName":"RabbitMQ","imageUrl":"https://www.rabbitmq.com/img/rabbitmq_logo_strap.png","longDescription":"Managed, highly available rabbitmq clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"https://www.rabbitmq.com/documentation.html","supportUrl":"https://www.rabbitmq.com/"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/plan/CC5ED8A2-1677-43A0-ADE5-27F713C6F51B
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/plan/CC5ED8A2-1677-43A0-ADE5-27F713C6F51B/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/plan/CC5ED8A2-1677-43A0-ADE5-27F713C6F51B/description "HA RabbitMQ on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/plan/CC5ED8A2-1677-43A0-ADE5-27F713C6F51B/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/86272DCB-86D5-4039-9E05-894436B8E71D/plan/CC5ED8A2-1677-43A0-ADE5-27F713C6F51B/free false



###创建服务 Redis
$ETCDCTL mkdir /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3 # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/name "Redis"
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/description "A Sample Redis (v2.8) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/tags 'redis,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/metadata '{"displayName":"Redis","imageUrl":"http://redis.io/images/redis-300dpi.png","longDescription":"Managed, highly available redis clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"http://redis.io/documentation","supportUrl":"http://redis.io"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan
###创建套餐1 (emptyDir)
$ETCDCTL mkdir /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/46ED475B-82A7-4C46-9767-0E3E806848F5
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/46ED475B-82A7-4C46-9767-0E3E806848F5/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/46ED475B-82A7-4C46-9767-0E3E806848F5/description "HA Redis on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/46ED475B-82A7-4C46-9767-0E3E806848F5/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/46ED475B-82A7-4C46-9767-0E3E806848F5/free false
###创建套餐2 (pvc)
$ETCDCTL mkdir /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/f8554827-4f67-4082-84af-6d35475c1ea8
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/f8554827-4f67-4082-84af-6d35475c1ea8/name "volumes_standalone"
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/f8554827-4f67-4082-84af-6d35475c1ea8/description "HA Redis With Volumes on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/f8554827-4f67-4082-84af-6d35475c1ea8/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/A54BC117-25E3-4E41-B8F7-14FC314D04D3/plan/f8554827-4f67-4082-84af-6d35475c1ea8/free false




###创建服务 Kafka
$ETCDCTL mkdir /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309 # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/name "Kafka"
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/description "A Sample Kafka (v0.9.0) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/tags 'kafka,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/metadata '{"displayName":"Kafka","imageUrl":"http://svn.apache.org/repos/asf/kafka/site/logos/kafka-logo-wide.png","longDescription":"Managed, highly available kafka clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"http://kafka.apache.org/documentation.html","supportUrl":"http://kafka.apache.org"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/plan/9A14FEF4-FB41-4C84-A175-A80492A50875
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/plan/9A14FEF4-FB41-4C84-A175-A80492A50875/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/plan/9A14FEF4-FB41-4C84-A175-A80492A50875/description "HA Kafka on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/plan/9A14FEF4-FB41-4C84-A175-A80492A50875/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/9972923D-0787-4271-839C-D000BD87E309/plan/9A14FEF4-FB41-4C84-A175-A80492A50875/free false



###创建服务 Cassandra
$ETCDCTL mkdir /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55 # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/name "Cassandra"
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/description "A Sample Cassandra (v3.4) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/tags 'cassandra,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/metadata '{"displayName":"Cassandra","imageUrl":"https://cassandra.apache.org/media/img/cassandra_logo.png","longDescription":"Managed, highly available cassandra clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"https://wiki.apache.org/cassandra/GettingStarted","supportUrl":"https://cassandra.apache.org/"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/plan/7B7EC041-2090-4ACB-AE0F-E8BDF315A778
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/plan/7B7EC041-2090-4ACB-AE0F-E8BDF315A778/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/plan/7B7EC041-2090-4ACB-AE0F-E8BDF315A778/description "HA Cassandra on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/plan/7B7EC041-2090-4ACB-AE0F-E8BDF315A778/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/3D7D7D07-D704-4B22-B492-EE5AE5301A55/plan/7B7EC041-2090-4ACB-AE0F-E8BDF315A778/free false



###创建服务 Storm
$ETCDCTL mkdir /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596 # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/name "Storm"
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/description "A Sample Storm (v0.9.2) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/tags 'storm,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/metadata '{"displayName":"Storm","imageUrl":"https://storm.apache.org/images/logo.png","longDescription":"Managed, highly available storm clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"https://storm.apache.org/releases/1.0.1/index.html","supportUrl":"https://storm.apache.org/"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/plan/D0B82741-770A-463C-818F-6E99862367DF
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/plan/D0B82741-770A-463C-818F-6E99862367DF/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/plan/D0B82741-770A-463C-818F-6E99862367DF/description "HA Storm on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/plan/D0B82741-770A-463C-818F-6E99862367DF/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/6555DBC1-E6BC-4F0D-8948-E1B5DF6BD596/plan/D0B82741-770A-463C-818F-6E99862367DF/free false



###创建服务 NiFi
$ETCDCTL mkdir /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/name "NiFi"
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/description "A Sample NiFi (v0.6.1) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/tags 'nifi,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/metadata '{"displayName":"NiFi","imageUrl":"https://nifi.apache.org/assets/images/nifiDrop.svg","longDescription":"Managed, highly available nifi clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"https://nifi.apache.org/docs.html","supportUrl":"https://nifi.apache.org"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/plan/4435A93C-6CC9-45F0-AFB0-EA070361DD6A
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/plan/4435A93C-6CC9-45F0-AFB0-EA070361DD6A/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/plan/4435A93C-6CC9-45F0-AFB0-EA070361DD6A/description "HA NiFi on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/plan/4435A93C-6CC9-45F0-AFB0-EA070361DD6A/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/BCC10E77-98B6-4EF0-8AFB-E5FD789F712E/plan/4435A93C-6CC9-45F0-AFB0-EA070361DD6A/free false



###创建服务 Kettle
$ETCDCTL mkdir /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/name "Kettle"
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/description "A Sample Kettle (v5.0.1) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/tags 'kettle,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/metadata '{"displayName":"Kettle","imageUrl":"http://www.pentaho.com/sites/all/themes/pentaho_resp/_media/logo-pentaho-n.png","longDescription":"Managed, highly available kettle clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"http://wiki.pentaho.com/display/EAI/Latest+Pentaho+Data+Integration+%28aka+Kettle%29+Documentation","supportUrl":"http://community.pentaho.com/projects/data-integration/"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/plan/31B428F8-AA3E-4CAC-85A2-7294C7CAA79D
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/plan/31B428F8-AA3E-4CAC-85A2-7294C7CAA79D/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/plan/31B428F8-AA3E-4CAC-85A2-7294C7CAA79D/description "HA Kettle on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/plan/31B428F8-AA3E-4CAC-85A2-7294C7CAA79D/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/51219A87-E37E-44F5-8E30-4767348D644D/plan/31B428F8-AA3E-4CAC-85A2-7294C7CAA79D/free false



###创建服务 TensorFlow
$ETCDCTL mkdir /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/name "TensorFlow"
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/description "A Sample TensorFlow (v0.8.0) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/tags 'tensorflow,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/metadata '{"displayName":"TensorFlow","imageUrl":"https://www.tensorflow.org/images/apple-touch-icon-180x180.png","longDescription":"Managed, highly available tensorflow clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"https://www.tensorflow.org/get_started","supportUrl":"https://www.tensorflow.org/"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/plan/BE1CAAF2-CAB7-4B56-AEB4-2A3111225D50
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/plan/BE1CAAF2-CAB7-4B56-AEB4-2A3111225D50/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/plan/BE1CAAF2-CAB7-4B56-AEB4-2A3111225D50/description "HA TensorFlow on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/plan/BE1CAAF2-CAB7-4B56-AEB4-2A3111225D50/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/2DD1363D-8768-4DAA-BDC3-FB29C10C4A8C/plan/BE1CAAF2-CAB7-4B56-AEB4-2A3111225D50/free false




###创建服务 PySpider
$ETCDCTL mkdir /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113 # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/name "PySpider"
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/description "A Sample PySpider (v0.3.7) cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/tags 'pyspider,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/metadata '{"displayName":"PySpider","imageUrl":"","longDescription":"A Powerful Spider(Web Crawler) System in Python.","providerDisplayName":"Asiainfo","documentationUrl":"https://docs.pyspider.org","supportUrl":"https://github.com/binux/pyspider"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/plan
###创建套餐1
$ETCDCTL mkdir /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/plan/1f195802-1642-47e9-be69-9082cc1ceaf5
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/plan/1f195802-1642-47e9-be69-9082cc1ceaf5/name "standalone"
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/plan/1f195802-1642-47e9-be69-9082cc1ceaf5/description "HA Spider on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/plan/1f195802-1642-47e9-be69-9082cc1ceaf5/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/c6ed3cb8-d486-4faa-8185-7262183a1113/plan/1f195802-1642-47e9-be69-9082cc1ceaf5/free false




###创建服务 MongoDB
$ETCDCTL mkdir /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f # 服务id

###创建服务级的配置
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/name "Mongo"
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/description "A Sample MongoDB cluster on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/bindable true
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/planupdatable false
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/tags 'mongodb,openshift'
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/metadata '{"displayName":"MongoDB","imageUrl":"https://webassets.mongodb.com/_com_assets/global/mongodb-logo-white.png","longDescription":"Managed, highly available MongoDB clusters in the cloud.","providerDisplayName":"Asiainfo","documentationUrl":"https://docs.mongodb.com/","supportUrl":"https://www.mongodb.com/"}'

###创建套餐目录
$ETCDCTL mkdir /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/plan
###创建套餐1 (pvc)
$ETCDCTL mkdir /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/plan/3b7fc05d-c630-4c0b-8dda-2cedb271ccb5
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/plan/3b7fc05d-c630-4c0b-8dda-2cedb271ccb5/name "volumes_standalone"
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/plan/3b7fc05d-c630-4c0b-8dda-2cedb271ccb5/description "HA MongoDB With Volumes on Openshift"
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/plan/3b7fc05d-c630-4c0b-8dda-2cedb271ccb5/metadata '{"bullets":["20 GB of Disk","20 connections"],"displayName":"Shared and Free" }'
$ETCDCTL set /servicebroker/openshift/catalog/eac6c8cf-2a63-4120-9e29-30c4b716e37f/plan/3b7fc05d-c630-4c0b-8dda-2cedb271ccb5/free false



# https://www.jcloud.com/help/db_qa.html#qa1
# 目前云数据库只支持通过子网访问


