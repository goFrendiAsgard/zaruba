from typing import Mapping, Any

from confluent_kafka.admin import AdminClient, NewTopic
import traceback
import sys

def create_kafka_topic(topic, config: Mapping[str, Any]):
    kafka_admin = AdminClient(config)
    topic_metadata = kafka_admin.list_topics()
    trial = 3
    while topic_metadata.topics.get(topic) is None and trial > 0:
        try:
            print({'action': 'create_kafka_topic', 'topic': topic})
            fs = kafka_admin.create_topics([NewTopic(topic, 1, 1)], request_timeout=15.0, validate_only=False)
            for _topic, f in fs.items():
                f.result()
            topic_metadata = kafka_admin.list_topics()
            trial -=1
        except:
            print(traceback.format_exc(), file=sys.stderr)
    if topic_metadata.topics.get(topic) is None and trial > 0:
        raise Exception('Cannot create topic: {}'.format(topic))