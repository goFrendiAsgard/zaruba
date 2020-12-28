from typing import Any, Callable
from sqlalchemy.orm import sessionmaker


def handle(DBSession: sessionmaker) -> Callable[..., Any]:
    def decorator(handler: Callable[..., Any]):
        def decorated_handler(*args: Any) -> Any:
            db = DBSession()
            try:
                result = handler(db, *args)
            finally:
                db.close()
            return result
        return decorated_handler
    return decorator
