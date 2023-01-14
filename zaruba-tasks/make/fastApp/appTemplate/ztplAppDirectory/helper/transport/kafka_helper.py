from typing import Mapping, Any
from confluent_kafka.admin import AdminClient, NewTopic
import logging


def create_kafka_topic(topic, config: Mapping[str, Any]):
    kafka_admin = AdminClient(config)
    topic_metadata = kafka_admin.list_topics()
    attempt, limit_attempt = 1, 3
    while attempt <= limit_attempt:
        try:
            logging.info('Creating kafka topic (attempt {} of {}): {}'.format(
                attempt, limit_attempt, topic
            ))
            fs = kafka_admin.create_topics(
                [NewTopic(topic, 1, 1)],
                request_timeout=15.0,
                validate_only=False
            )
            for _, f in fs.items():
                f.result()
            topic_metadata = kafka_admin.list_topics()
        except Exception:
            logging.error(
                'Error creating kafka topic (attempt {} of {}): {}'.format(
                    attempt, limit_attempt, topic
                ),
                exc_info=True
            )
        attempt += 1
        if topic_metadata.topics.get(topic) is not None:
            break
    if topic_metadata.topics.get(topic) is None:
        raise Exception('Cannot create topic: {}'.format(topic))
