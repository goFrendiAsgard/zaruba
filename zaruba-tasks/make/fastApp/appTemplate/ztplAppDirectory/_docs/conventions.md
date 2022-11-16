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


<!--startTocSubTopic-->
<!--endTocSubTopic-->