# Kafka Practice Repo

### Installing and Running Kafka + Zookeeper locally (Mac)

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

## Interact with Kafka by CLI

### Create a Kafka Topic

change directories, where the kafka binarys are installed. On mac that is typically here:
`cd /usr/local/bin`

## Create Topics

```
./kafka-topics --bootstrap-server 127.0.0.1:9092 --topic first_topic --create
```

or with more handling

```
~/usr/local/bin/kafka-topics --bootstrap-server 127.0.0.1:9092 --topic second_topic --create --partitions 3 --replication-factor 1
```

Should see an output similar to this:
`Created topic second_topic.`

```
./kafka-console-consumer --bootstrap-server 127.0.0.1:9092 --topic first_topic --from-beginning
```

## Produce Messages to Topic

```
./kafka-console-producer --bootstrap-server 127.0.0.1:9092 --topic first_topic
```

Submit text messages like below

`> First Message`
`> Second Message`
`> Third Message`

Then cancel out of the cli.

## Consume Messages from a topic

```
./kafka-console-consumer --bootstrap-server 127.0.0.1:9092 --topic first_topic --from-beginning
```

`First Message`
`Second Message`
`Third Message`

## Interact with Kafka + Zookeeper within a Container (Docker)

1. Ensure Docker is installed and running. To download see [here](https://docs.docker.com/get-docker/). Confirm that docker is running before continuing

2. Execute `docker-compose up -d` within the root directory of this application. This command will run docker-compose to create 2 containers, 1 running kafka and 1 running zookeeper.

3. Confirm the containers were successfully created by running `docker ps`. There should be two containers logs that look similar to the output below:

```
6ba38109e236   wurstmeister/kafka:latest   "start-kafka.sh"         12 minutes ago   Up 12 minutes   0.0.0.0:9092->9092/tcp  kafka
8a5f9ce02046   zookeeper:latest            "/docker-entrypoint.…"   17 minutes ago   Up 17 minutes   2888/tcp, 3888/tcp, 0.0.0.0:2181->2181/tcp, 8080/tcp  zookeeper
```

4. Enter the bash shell within the kafka container by executing

```
docker exec -it kafka /bin/bash
```

Change directories into the directory that contains the kafka scripts and binaries. The file path is `/opt/kafka/bin`.

Execute the new commands to create a topic, to produce messages to the topic, and to consume the topic.
`./kafka-topics.sh --bootstrap-server 127.0.0.1:9092 --topic first_topic --create`
`./kafka-console-producer.sh --bootstrap-server 127.0.0.1:9092 --topic first_topic`
`./kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic first_topic`

Exit bash shell with `Ctrl + D`

5. Tear down containers with `docker-compose down`

## Interact with Kafka via Go + Docker

1. Ensure Docker is installed and running. To download see [here](https://docs.docker.com/get-docker/). Confirm that docker is running before continuing.

2. Execute `docker-compose up -d` within the root directory of this application. This command will run docker-compose to create 2 containers, 1 running kafka and 1 running zookeeper.

3. Confirm the containers were successfully created by running `docker ps`. There should be two containers logs that look similar to the output below:

```
6ba38109e236   wurstmeister/kafka:latest   "start-kafka.sh"         12 minutes ago   Up 12 minutes   0.0.0.0:9092->9092/tcp  kafka
8a5f9ce02046   zookeeper:latest            "/docker-entrypoint.…"   17 minutes ago   Up 17 minutes   2888/tcp, 3888/tcp, 0.0.0.0:2181->2181/tcp, 8080/tcp  zookeeper
```

4. Run the go program, first compile the program into a binary and then execute the binary by executing the follow commands.

```
go build .
./kafka-practice
```

The Go program connects to the running kafka instance, initializes a producer and consumer, creates a topic, and submit messages to that topic. In the command line, message producing/consuming will be logged to show the output of each step. The program should remain running and will continue to push a message from the producer to the topic every second.

In the go program, the consumer is ran on a separate go routine, so that in theory our program will mimic a consumer consuming topic messages as the produces publishes them. In Step 5, we'll see it in action!

5. In a new window, enter a bash shell within the kafka docker container by executing

```
docker exec -it kafka /bin/bash
```

change directories to the kafka binaries by executing, found in the file path here `/opt/kafka/bin/`

Consume the messages, starting from the beginning by executing...

```
./kafka-console-consumer.sh --bootstrap-server 127.0.0.1:9092 --topic myTopic --from-beginning
```

Now as new messages are published from the go program, this consumer will output those messages when they are logged. Exit the consumer script with `Ctrl + C` and the container with `Ctrl + D`. End the go program with `Ctrl + C`.

6. Tear down the docker containers with

```
docker-compose down
```
