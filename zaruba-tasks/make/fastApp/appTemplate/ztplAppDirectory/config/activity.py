from typing import List
import os
import json

# List of activity events
activity_events: List[str] = json.loads(
    os.getenv('APP_ACTIVITY_EVENTS', '[]')
)
