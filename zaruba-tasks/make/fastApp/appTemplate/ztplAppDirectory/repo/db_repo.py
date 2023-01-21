from typing import Any, Generic, List, Mapping, Optional, TypeVar, Type
from pydantic import BaseModel
from sqlalchemy import or_
from sqlalchemy.engine import Engine
from sqlalchemy.orm import Session
from repo.base import Base
from sqlalchemy.orm.attributes import InstrumentedAttribute

import uuid
import datetime
import logging

SchemaData = TypeVar('SchemaData', bound=BaseModel)
Schema = TypeVar('Schema', bound=BaseModel)
DBEntity = TypeVar('DBEntity', bound=Base)


class DBRepo(Generic[DBEntity, Schema, SchemaData]):
    '''
    Database repository

    Usage:

    ```
    from typing import Optional
    from repo import Base, BaseMixin, DBRepo
    from sqlalchemy import create_engine, Column, DateTime, String
    from pydantic import BaseModel

    import datetime

    # SQLAlchemy model
    class DBBookEntity(Base, BaseMixin):
        title = Column(String(255), index=True, nullable=False)

    # Pydantic schema
    class BookData(BaseModel):
        title: str
        created_at: Optional[datetime.datetime]
        created_by: Optional[str]
        updated_at: Optional[datetime.datetime]
        updated_by: Optional[str]

    class Book(BookData):
        id: str
        class Config:
            orm_mode = True

    # DBBookRepo definition
    class DBBookRepo(DBRepo[DBBookEntity, Book, BookData]):
        schema_class = Book
        db_entity_class = DBBookEntity

    engine = create_engine('sqlite:///database.db', echo=True)
    book_repo = DBBookRepo(engine=engine, create_all=True)

    new_book = book_repo.insert(BookData(title='Doraemon'))
    print(new_book)
    books = book_repo.find()
    print(books)
    ```
    '''

    schema_class: Type[Schema]
    db_entity_class: Type[DBEntity]

    def __init__(
        self,
        engine: Engine,
        create_all: bool = True
    ):
        self.engine = engine
        self.db_entity_attribute_names: List[str] = dir(self.db_entity_class)
        self.schema_attribute_names: List[str] = dir(self.schema_class)
        if create_all:
            Base.metadata.create_all(bind=engine)

    def get_keyword_fields(self) -> List[InstrumentedAttribute]:
        '''
        Return list of fields for keyword filtering
        '''
        return [
            getattr(self.db_entity_class, field)
            for field in self.db_entity_attribute_names
            if type(
                getattr(self.db_entity_class, field)
            ) == InstrumentedAttribute
        ]

    def from_schema_data_to_db_entity_dict(
        self, schema_data: SchemaData
    ) -> Mapping[str, Any]:
        '''
        Convert entity_data into dictionary
        The result of this convertion is used for inserting/updating db_entity.
        '''
        entity_dict = schema_data.dict()
        return {
            field: value
            for field, value in entity_dict.items()
            if field in self.db_entity_attribute_names
        }

    def from_db_entity_to_schema(self, db_entity: DBEntity) -> Schema:
        '''
        Convert db_entity into schema.
        The result of this convertion is usually returned to external layer (e.g., service).
        '''
        return self.schema_class.from_orm(db_entity)

    def get_search_filter(self, db: Session, keyword: str) -> Any:
        '''
        Return keyword filtering.
        The result is usually used to invoke find/count.
        '''
        like_keyword = '%{}%'.format(keyword) if keyword != '' else '%'
        keyword_filter = [
            keyword_field.like(like_keyword)
            for keyword_field in self.get_keyword_fields()
        ]
        return or_(*keyword_filter)
    
    def create_db_sesion(self) -> Session:
        '''
        Return a db session.
        '''
        return Session(self.engine, expire_on_commit=False)

    def find_by_id(self, id: str) -> Optional[Schema]:
        '''
        Find a record by id.
        '''
        db = self.create_db_sesion()
        try:
            search_filter = self.db_entity_class.id == id
            return self.fetch_one_by_filter(db, search_filter)
        finally:
            db.close()

    def find(
        self, keyword: str, limit: int = 100, offset: int = 0
    ) -> List[Schema]:
        '''
        Find multiple records by keyword with limit and offset.
        '''
        db = self.create_db_sesion()
        try:
            search_filter = self.get_search_filter(db, keyword)
            return self.fetch_by_filter(db, search_filter, limit, offset)
        finally:
            db.close()

    def count(self, keyword: str) -> int:
        '''
        Count records by keyword.
        '''
        db = self.create_db_sesion()
        try:
            search_filter = self.get_search_filter(db, keyword)
            return self.count_by_filter(db, search_filter)
        finally:
            db.close()

    def insert(self, schema_data: SchemaData) -> Optional[Schema]:
        '''
        Insert a new record.
        '''
        db = self.create_db_sesion()
        try:
            db_entity = self.db_entity_class(
                ** self.from_schema_data_to_db_entity_dict(schema_data),
            )
            if 'id' in self.db_entity_attribute_names:
                new_id = str(uuid.uuid4())
                db_entity.id = new_id
            if 'created_at' in self.db_entity_attribute_names:
                db_entity.created_at = datetime.datetime.utcnow()
            if 'updated_at' in self.db_entity_attribute_names:
                db_entity.updated_at = datetime.datetime.utcnow()
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity)
            return self.from_db_entity_to_schema(db_entity)
        except Exception:
            self._log_error(
                'Error while invoking insert {}'.format(
                    self.db_entity_class
                ),
                'schema_data={}'.format(schema_data)
            )
            raise
        finally:
            db.close()

    def update(
        self, id: str, schema_data: SchemaData
    ) -> Optional[Schema]:
        '''
        Update a record.
        '''
        db = self.create_db_sesion()
        try:
            db_entity = db.query(self.db_entity_class).filter(
                self.db_entity_class.id == id
            ).first()
            if db_entity is None:
                return None
            db_entity_dict = self.from_schema_data_to_db_entity_dict(
                schema_data
            )
            for field, value in db_entity_dict.items():
                if field == 'created_at' or field == 'created_by':
                    continue
                setattr(db_entity, field, value)
            if 'updated_at' in self.db_entity_attribute_names:
                db_entity.updated_at = datetime.datetime.utcnow()
            db.add(db_entity)
            db.commit()
            db.refresh(db_entity)
            return self.from_db_entity_to_schema(db_entity)
        except Exception:
            self._log_error(
                'Error while invoking update {}'.format(
                    self.db_entity_class
                ),
                'id={}'.format(id),
                'schema_data={}'.format(schema_data)
            )
            raise
        finally:
            db.close()

    def delete(self, id: str) -> Optional[Schema]:
        '''
        Delete a record.
        '''
        db = self.create_db_sesion()
        try:
            db_entity = db.query(self.db_entity_class).filter(
                self.db_entity_class.id == id
            ).first()
            if db_entity is None:
                return None
            db.delete(db_entity)
            db.commit()
            return self.from_db_entity_to_schema(db_entity)
        except Exception:
            self._log_error(
                'Error while invoking delete {}'.format(
                    self.db_entity_class
                ),
                'id={}'.format(id),
            )
            raise
        finally:
            db.close()

    def fetch_by_filter(
        self,
        db: Session,
        search_filter: Any,
        limit: int = 100,
        offset: int = 0
    ) -> List[Schema]:
        try:
            db_query = db.query(self.db_entity_class).filter(
                search_filter
            )
            if 'created_at' in self.db_entity_attribute_names:
                db_query = db_query.order_by(
                    self.db_entity_class.created_at.desc()
                )
            db_entities = db_query.offset(offset).limit(limit).all()
            return [
                self.from_db_entity_to_schema(db_entity)
                for db_entity in db_entities
            ]
        except Exception:
            self._log_error(
                'Error while invoking find_by_filter {}'.format(
                    self.db_entity_class
                ),
                'search_filter={}'.format(search_filter),
                'limit={}'.format(limit),
                'offset={}'.format(offset)
            )
            raise

    def count_by_filter(
        self,
        db: Session,
        search_filter: Any,
        limit: int = 100,
        offset: int = 0
    ) -> List[Schema]:
        try:
            return db.query(self.db_entity_class).filter(
                search_filter
            ).count()
        except Exception:
            self._log_error(
                'Error while invoking count {}'.format(
                    self.db_entity_class
                ),
                'search_filter={}'.format(search_filter)
            )
            raise

    def fetch_one_by_filter(
        self, db: Session, search_filter: Any
    ) -> Optional[Schema]:
        try:
            db_entity = db.query(self.db_entity_class).filter(
                search_filter
            ).first()
            if db_entity is None:
                return None
            return self.from_db_entity_to_schema(db_entity)
        except Exception:
            self._log_error(
                'Error while invoking find_by_id {}'.format(
                    self.db_entity_class
                ),
                'search_filter={}'.format(search_filter)
            )
            raise
    
    def _log_error(*messages: str):
        logging.error(
            ' '.join(messages),
            exc_info=True
        )

