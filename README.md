export ZOOK=10.246.88.2:2181
./bin/kafka-topics.sh --create --zookeeper $ZOOK --replication-factor 1 --partitions 1 --topic test

./bin/kafka-topics.sh --list --zookeeper $ZOOK

./bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
This is a message
This is another message

./bin/kafka-console-consumer.sh --zookeeper $ZOOK --topic test --from-beginning
