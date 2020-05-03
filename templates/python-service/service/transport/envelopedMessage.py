class EnvelopedMessage:

    def __init__(self):
        self.correlation_id = ""
        self.message = {}
        self.error_message = ""

    def to_json(self):
        data = {"correlation_id": self.correlation_id}
