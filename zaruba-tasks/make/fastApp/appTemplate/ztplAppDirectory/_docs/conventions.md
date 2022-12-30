<!--startTocHeader-->
[🏠](README.md)
# Conventions
<!--endTocHeader-->

# Naming convention

Clear naming convention will make your code more readable, and expectable.

## ✔️ Do

- Give a proper and meaningful names.
- Follow Python naming conventions.
    - Contant names should be `UPPER_CASE_AND_SEPARATED_BY_UNDERSCORE`
    - Class names should be `PascalCase`
    - Variable names should be `snake_case`
    - File/module names should be `lowercase` or `snake_case`


```python
DEFAULT_USER_NAME = 'root'

class UserSeeder():

    def __init__(self, user_name: str):
        self.user_name = user_name


user_seeder = UserSeeder(user_name=DEFAULT_USER_NAME)
```

## ❌ Don't

- Give non meaningful name
- Break Python naming convention

```python
CONSTANT = 'root' # not meaningful

class My_Seeder(): # not meaningful, break Python convention

    def __init__(self, a: str): # a is not meaningful
        self.a = a


x = My_Seeder(user_name=CONSTANT) # x is not meaningful

```

# Dependency Injection

Dependency injection allows you to change the implementation without doing too much modification in your code.

Prefer dependency injection whenever possible.

## ✔️ Do

- Use interface
- Pass dependencies as function/constructor arguments
- Initiate everything in `main.py`

### Define interfaces for dependencies

```python

class BookRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Optional[Book]:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[Book]:
        pass
```

### Make implementation for the interfaces

```python
class DBBookRepo(BookRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)


    def find_by_id(self, id: str) -> Optional[Book]:
        db = Session(self.engine, expire_on_commit=False)
        book: Book
        try:
            db_book = db.query(DBBookEntity).filter(DBBookEntity.id == id).first()
            if db_book is None:
                return None
            book = Book.from_orm(db_book)
        finally:
            db.close()
        return book


    def find(self, keyword: str, limit: int, offset: int) -> List[Book]:
        db = Session(self.engine, expire_on_commit=False)
        books: List[Book] = []
        try:
            keyword_filter = self._get_keyword_filter(keyword)
            db_books = db.query(DBBookEntity).filter(DBBookEntity.title.like(keyword_filter)).offset(offset).limit(limit).all()
            books = [Book.from_orm(db_result) for db_result in db_books]
        finally:
            db.close()
        return books


```

### Use Interfaces as dependencies

```python
class BookService():

    def __init__(self, mb: AppMessageBus, rpc: AppRPC, book_repo: BookRepo):
        self.mb = mb
        self.rpc = rpc
        self.book_repo = book_repo
```

### Initiate everything in main.py

```python
mb = LocalMessageBus()
rpc = LocalRPC()
book_repo = DBBookRepo(engine=engine, create_all=db_create_all)
book_service = BookService(mb, rpc, book_repo)
```

Now, you will be able to change `book_repo` implementation without touching `book_service` implementation.

## ❌ Don't

- Initiate dependencies inside the implementation


```python
class BookService():

    def __init__(self):
        self.mb = LocalMessageBus()
        self.rpc = LocalRPC()
        self.book_repo = DBBookRepo(engine=engine, create_all=db_create_all)
```

Suppose you have multiple services (not just `BookService`), and you want to change `mb` or `rpc` implementation. Then you have to modify all services.

# Inter module communication

Always use RPC/events if you need to do inter module communication.

If you deploy your application as monolith, you can use `LocalRPC` and `LocalMessagebus`. Later if you want to deploy your application as microservices, those interfaces can be changed without changing any other implementation.

## ✔️ Do

- Always Use RPC/messagebus.

For example, you want to notify `ActivityService` in `log` module about book-insertion event.

First you need to create an `event handler` in `log` module as follow:

```python

def register_activity_event(mb: MessageBus, rpc: RPC, auth_service: AuthService, activity_service: ActivityService):

    @mb.handle('new_activity')
    def insert_activity(activity_data: Mapping[str, Any]) -> Optional[Mapping[str, Any]]:
        activity = ActivityData.parse_obj(activity_data) 
        new_activity = activity_service.insert(activity)
        return None if new_activity is None else new_activity.dict()

```

```python
class BookService():

    def __init__(self, mb: AppMessageBus, rpc: AppRPC, book_repo: BookRepo):
        self.mb = mb
        self.rpc = rpc
        self.book_repo = book_repo


    def insert(self, book_data: BookData, current_user: User) -> Optional[Book]:
        book_data.created_by = current_user.id
        book_data.updated_by = current_user.id
        book_data = self._validate_data(book_data)
        new_book = self.book_repo.insert(book_data)
        self.mb.publish('new_activity', ActivityData(
            user_id = current_user.id,
            activity = 'insert',
            object = 'book',
            row = new_book.dict(),
            row_id = new_book.id
        ))
        return new_book
```

## ❌ Don't

- Construct other module service
- Call other module service


```python
from modules.other_module.action import ActionService

class BookService():

    def __init__(self, mb: AppMessageBus, rpc: AppRPC, book_repo: BookRepo, action_service: ActionService):
        self.mb = mb
        self.rpc = rpc
        self.book_repo = book_repo
        # When you deploy the app as microservices, log module might has it's own database
        # If you use action_service here, you limit the microservices to use the same database
        self.action_service = action_service


    def insert(self, book_data: BookData, current_user: User) -> Optional[Book]:
        book_data.created_by = current_user.id
        book_data.updated_by = current_user.id
        book_data = self._validate_data(book_data)
        new_book = self.book_repo.insert(book_data)
        self.action_service.insert(ActivityData(
            user_id = current_user.id,
            activity = 'insert',
            object = 'book',
            row = new_book.dict(),
            row_id = new_book.id
        ))
        return new_book
```


# Subtopics
<!--startTocSubtopic-->
<!--endTocSubtopic-->