# Creating Fast API CRUD

CRUD (create-read-update-delete) is a common use case for any database application. CRUD is easy, yet tedious. Zaruba helps you creating CRUD feature so that you don't need to code everything by yourself.

To create a Fast API CRUD, you can invoke `zaruba please makeFastApiCrud -i`

Under the hood, zaruba use [SQLAlchemy](https://www.sqlalchemy.org/), [message bus](fast-api-message-bus.md), [route](creating-fast-api-route.md), and [RPC handler](creating-fast-api-rpc-handler.md).

Also You need to make sure that you already have [Fast Api service](creating-fast-api-service.md) and [Fast Api module](creating-fast-api-module.md)

When you make a CRUD, zaruba will add/modify several files:

* `schemas/<entity>.py`: Data structure definition for CRUD-related communication.
* `repos/<entity>.py`: Repo interface definition.
* `repos/db<Entity>.py`: Repo interface implementation using SQLAlchemy.
* `<module>/handle<Entity>Event.py`: RPC handler for CRUD related action.
* `<module>/handle<Entity>Route.py`: Route handler for CRUD related action.
* `<module>/controller.py`: Module controller, the one that load event/route handler.
* `main.py`: Main file that define all components, including database repositories.

Suppose you have `myModule` as module name, `book` as entity, then zaruba will add/modify these files:

## schemas/book.py

```python
from pydantic import BaseModel
import datetime

class BookData(BaseModel):
    title: str
    author: str
    synopsis: str


class Book(BookData):
    id: str
    created_at: datetime.datetime
    updated_at: datetime.datetime
    class Config:
        orm_mode = True
```

`BookData` defines the data being used for communication (i.e: when create/update), while `Book` defines the data returned by the service (thus there are two additional fields, `created_at` and `updated_at`).

These schema is going to be used by `RPC handler`, `route handler`, and `repo`.

## repos/book.py

```python
from typing import List
from schemas.book import Book, BookData

import abc

class BookRepo(abc.ABC):

    @abc.abstractmethod
    def find_by_id(self, id: str) -> Book:
        pass

    @abc.abstractmethod
    def find(self, keyword: str, limit: int, offset: int) -> List[Book]:
        pass

    @abc.abstractmethod
    def insert(self, book_data: BookData) -> Book:
        pass

    @abc.abstractmethod
    def update(self, id: str, book_data: BookData) -> Book:
        pass

    @abc.abstractmethod
    def delete(self, id: str) -> Book:
        pass
```

`BookRepo` is an interface (technically it is Abstract Base Class) to define how our entity repo should behave. In this case, we define that any repository that deal with book entity should have 5 methods: `find_by_id`, `find`, `insert`, `update`, and `delete`.

Defining an interface let you define your own implementation in case of you don't want to use SQLAlchemy. As long as your implementation comply with the interface, then you are good to go.

## repos/dbBook.py

```python
from typing import List
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from sqlalchemy import Boolean, Column, DateTime, ForeignKey, Integer, String
from schemas.book import Book, BookData
from repos.book import BookRepo

import uuid
import datetime

Base = declarative_base()

class DBBookEntity(Base):
    __tablename__ = "book"
    id = Column(String(36), primary_key=True, index=True)
    title = Column(String(20), index=True)
    author = Column(String(20), index=True)
    synopsis = Column(String(20), index=True)
    created_at = Column(DateTime, default=datetime.datetime.utcnow)
    updated_at = Column(DateTime, default=datetime.datetime.utcnow)


class DBBookRepo(BookRepo):

    def __init__(self, engine: Engine, create_all: bool):
        self.engine = engine
        if create_all:
            Base.metadata.create_all(bind=engine)


    def find_by_id(self, id: str) -> Book:
        db = Session(self.engine)
        result: Book
        try:
            db_result = db.query(DBBookEntity).filter(DBBookEntity.id == id).first()
            if db_result is None:
                return None
            result = Book.from_orm(db_result)
        finally:
            db.close()
        return result

    
    def find(self, keyword: str, limit: int, offset: int) -> List[Book]:
        db = Session(self.engine)
        results: List[Book] = []
        try:
            keyword = '%{}%'.format(keyword) if keyword != '' else '%'
            db_results = db.query(DBBookEntity).filter(DBBookEntity.title.like(keyword)).offset(offset).limit(limit).all()
            results = [Book.from_orm(db_result) for db_result in db_results]
        finally:
            db.close()
        return results

    
    def insert(self, book_data: BookData) -> Book:
        db = Session(self.engine)
        result: Book
        try:
            db_entity = DBBookEntity(
                id=str(uuid.uuid4()),
                title=book_data.title, 
                author=book_data.author, 
                synopsis=book_data.synopsis, 
                created_at=datetime.datetime.utcnow()
            )
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity) 
            result = Book.from_orm(db_entity)
        finally:
            db.close()
        return result
    
    def update(self, id: str, book_data: BookData) -> Book:
        db = Session(self.engine)
        result: Book
        try:
            db_entity = db.query(DBBookEntity).filter(DBBookEntity.id == id).first()
            if db_entity is None:
                return None
            db_entity.title = book_data.title
            db_entity.author = book_data.author
            db_entity.synopsis = book_data.synopsis
            db_entity.updated_at = datetime.datetime.utcnow()
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity) 
            result = Book.from_orm(db_entity)
        finally:
            db.close()
        return result

 
    def delete(self, id: str) -> Book:
        db = Session(self.engine)
        result: Book
        try:
            db_entity = db.query(DBBookEntity).filter(DBBookEntity.id == id).first()
            if db_entity is None:
                return None
            db.delete(db_entity)
            db.commit()
            result = Book.from_orm(db_entity)
        finally:
            db.close()
        return result
```

`DBBook` is an SQLAlchemy implementation of `BookRepo`. Since SQLAlchemy able to deal with multiple SQL dialect, usually this repository implementation is enough.

## myModule/handleBookEvent.py

```python
from typing import Any, List, Mapping
from helpers.transport import MessageBus
from schemas.book import Book, BookData
from repos.book import BookRepo

def handle_event(mb: MessageBus, book_repo: BookRepo):

    @mb.handle_rpc('find_book')
    def find_book(keyword: str, limit: int, offset: int) -> List[Mapping[str, Any]]:
        results = book_repo.find(keyword, limit, offset)
        return [result.dict() for result in results]


    @mb.handle_rpc('find_book_by_id')
    def find_book_by_id(id: str) -> Mapping[str, Any]:
        result = book_repo.find_by_id(id)
        if result is None:
            return None
        return result.dict()


    @mb.handle_rpc('insert_book')
    def insert_book(data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = book_repo.insert(BookData.parse_obj(data))
        if result is None:
            return None
        return result.dict()


    @mb.handle_rpc('update_book')
    def update_book(id: str, data: Mapping[str, Any]) -> Mapping[str, Any]:
        result = book_repo.update(id, BookData.parse_obj(data))
        if result is None:
            return None
        return result.dict()


    @mb.handle_rpc('delete_book')
    def delete_book(id: str) -> Mapping[str, Any]:
        result = book_repo.delete(id)
        if result is None:
            return None
        return result.dict()
    

    print('Handle event for book')
```

This is an RPC handler to deal with book CRUD. This is going to be useful if you decide to deal with microservices that need to talk to each other.

## myModule/handleBookRoute.py

```python
from typing import Any, List, Mapping
from helpers.transport import MessageBus
from fastapi import FastAPI, HTTPException
from schemas.book import Book, BookData

import traceback

def handle_route(app: FastAPI, mb: MessageBus):

    @app.get('/book/', response_model=List[Book])
    def find_book(keyword: str='', limit: int=100, offset: int=0):
        try:
            results = mb.call_rpc('find_book', keyword, limit, offset)
            return [Book.parse_obj(result) for result in results]
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')


    @app.get('/book/{id}', response_model=Book)
    def find_book_by_id(id: str):
        try:
            result = mb.call_rpc('find_book_by_id', id)
            if result is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return Book.parse_obj(result)
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')


    @app.post('/book/', response_model=Book)
    def insert_book(data: BookData):
        try:
            result = mb.call_rpc('insert_book', data.dict())
            if result is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return Book.parse_obj(result)
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')


    @app.put('/book/{id}', response_model=Book)
    def update_book(id: str, data: BookData):
        try:
            result = mb.call_rpc('update_book', id, data.dict())
            if result is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return Book.parse_obj(result)
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')


    @app.delete('/book/{id}')
    def delete_book(id: str):
        try:
            result = mb.call_rpc('delete_book', id)
            if result is None:
                raise HTTPException(status_code=404, detail='Not Found')
            return Book.parse_obj(result)
        except HTTPException as error:
            raise error
        except Exception as error:
            print(traceback.format_exc()) 
            raise HTTPException(status_code=500, detail='Internal Server Error')


    print('Handle route for book')
```

This is our CRUD related route handler.

## myModule/controller.py

```python
from typing import Mapping, List, Any
from fastapi import FastAPI, HTTPException
from helpers.transport import MessageBus
from repos.book import BookRepo
from myModule.handleBookRoute import handle_route as handle_book_route
from myModule.handleBookEvent import handle_event as handle_book_event

import traceback

class Controller():

    def __init__(self, app: FastAPI, mb: MessageBus, enable_route: bool, enable_event: bool, book_repo: BookRepo):
        self.app = app
        self.mb = mb
        self.enable_route = enable_route
        self.enable_event = enable_event
        self.book_repo = book_repo


    def start(self):
        if self.enable_event:
            handle_book_event(self.mb, self.book_repo)
            self.handle_event()
        if self.enable_route:
            handle_book_route(self.app, self.mb)
            self.handle_route()
    

    def handle_event(self):
        print('Handle events for myModule')
    

    def handle_route(self):
        print('Handle routes for myModule')
```

You have probably notice that now our controller's constructor need `BookRepo` parameter. Also, `handle_book_event` and `handle_book_route` is now added to controller's `start`.

## main.py

```python
from fastapi import FastAPI
from sqlalchemy import create_engine
from helpers.transport import MessageBus, RMQMessageBus, RMQEventMap, LocalMessageBus
from myModule.controller import Controller as MyModuleController
from repos.dbBook import DBBookRepo


import os

def create_message_bus(mb_type: str) -> MessageBus:
    if mb_type == 'rmq':
        rmq_host = os.getenv('MY_SERVICE_RABBITMQ_HOST', 'localhost')
        rmq_user = os.getenv('MY_SERVICE_RABBITMQ_USER', 'root')
        rmq_pass = os.getenv('MY_SERVICE_RABBITMQ_PASS', 'toor')
        rmq_vhost = os.getenv('MY_SERVICE_RABBITMQ_VHOST', '/')
        rmq_event_map = RMQEventMap({})
        return RMQMessageBus(rmq_host, rmq_user, rmq_pass, rmq_vhost, rmq_event_map)
    return LocalMessageBus()

db_url = os.getenv('MY_SERVICE_SQLALCHEMY_DATABASE_URL', 'sqlite://')
mb_type = os.getenv('MY_SERVICE_MESSAGE_BUS_TYPE', 'local')
enable_route = os.getenv('MY_SERVICE_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_event = os.getenv('MY_SERVICE_ENABLE_EVENT_HANDLER', '1') != '0'

engine = create_engine(db_url, echo=True)
app = FastAPI()
mb = create_message_bus(mb_type)

@app.on_event('shutdown')
def on_shutdown():
    mb.shutdown()

book_repo = DBBookRepo(engine=engine, create_all=True)
my_module_controller = MyModuleController(app=app, mb=mb, enable_route=enable_route, enable_event=enable_event, book_repo=book_repo)
my_module_controller.start()
```

Finally, we initiate `BookRepo` and inject it to our controller. Now everything is ready and connected.


# Example


```sh
# run interactively
zaruba please makeFastApiCrud -i

# run with paramter
zaruba please makeFastApiCrud generator.fastApi.service.name=myService generator.fastApi.module.name=myModule generator.fastApi.crud.entity=book generator.fastApi.crud.fields=title,author
```

# Involved tasks

* [makeFastApiCrud](tasks/makeFastApiCrud.md)


# What's next

* [Creating Fast API service task](creating-fast-api-service-task.md)
* [Creating Fast API route](creating-fast-api-route.md)
* [Creating Fast API event handler](creating-fast-api-event-handler.md)
* [Creating Fast API RPC handler](creating-fast-api-rpc-handler.md)