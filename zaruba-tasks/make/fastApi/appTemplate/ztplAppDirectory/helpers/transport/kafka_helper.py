from typing import Mapping, Any

from confluent_kafka.admin import AdminClient, NewTopic
import traceback

def create_kafka_topic(topic, config: Mapping[str, Any]):
    try:
        kafka_admin = AdminClient(config)
        topic_metadata = kafka_admin.list_topics()
        if topic_metadata.topics.get(topic) is None:
            print({'action': 'create_kafka_topic', 'topic': topic})
            kafka_admin.create_topics([NewTopic(topic, 1, 1)])
    except:
        print(traceback.format_exc())