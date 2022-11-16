from typing import List
import os
import json

activity_events: List[str] = json.loads(os.getenv('APP_ACTIVITY_EVENTS', '[]'))