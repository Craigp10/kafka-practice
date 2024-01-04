# Kafka Practice Repo

## Installing and Running Kafka + Zookeeper (Mac)

1. Install kafka - If on mac I reocmmend using homebrew:
   `$ brew install kafka`
   Or use the offical Kafka docs page [here](https://kafka.apache.org/quickstart).

   Kafka should now be found within the directory...
   `/usr/local/bin` if on mac.

2. Install Zookeeper - If on mac I recommend using homebrew:
   `$ brew install zookeeper`
   or use the offcial docs page [here](https://zookeeper.apache.org/releases.html).
   <b>Why Zookeeper?</b> - ZooKeeper is a centralized service for maintaining configuration information, naming, providing distributed synchronization, and providing group services.More information on the program can be found [here](https://zookeeper.apache.org/).

   Can be found here
   `~/Downloads/apache-zookeeper-3.8.3-bin/bin`

3. run Zookeeper
   if installed with homebrew:
   `$ brew services start zookeeper`
   onfirm it is running with
   `$ brew services`

4. run Kafka
   change directories to where kafka binary is stored and run the below command.
   `$ kafka-server-start /usr/local/etc/kafka/server.properties`
   Or if installed with homebrew.
   `brew services start kafka`
   \*Note that special configs aren't as straight forward with this approach.

### Create a Kafka Topic

change directories, where the kafka binarys are installed. On mac that is typically here:
`cd /usr/local/bin`

## Create Topics

`$ ./kafka-topics --bootstrap-server 127.0.0.1:9092 --topic first_topic --create`
or with more handling
`$ ~/usr/local/bin/kafka-topics --bootstrap-server 127.0.0.1:9092 --topic second_topic --create --partitions 3 --replication-factor 1`

Should see an output similar to this:
`Created topic second_topic.`

`./kafka-console-consumer --bootstrap-server 127.0.0.1:9092 --topic first_topic --from-beginning`

## Produce Messages to Topic

`./kafka-console-producer --bootstrap-server 127.0.0.1:9092 --topic first_topic`

Submit text messages like below

```
> First Message
> Second Message
> Third Message
```

Then cancel out of the cli.

## Consume Messages from a topic

`./kafka-console-consumer --bootstrap-server 127.0.0.1:9092 --topic first_topic --from-beginning`

```
First Message
Second Message
Third Message
```
