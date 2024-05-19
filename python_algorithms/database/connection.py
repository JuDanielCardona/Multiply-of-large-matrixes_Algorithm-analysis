import os
from sqlalchemy import create_engine, Table, MetaData, Column, Integer, String, Float
from sqlalchemy.orm import sessionmaker
from sqlalchemy.exc import SQLAlchemyError

# Cargar variables de entorno para las credenciales
DATABASE_URL = os.getenv('DATABASE_URL', 'postgresql://admin:12345@localhost:5432/executiontimes_db')

# Crear motor de base de datos
engine = create_engine(DATABASE_URL)

# Crear metadatos
metadata = MetaData()

# Definir la tabla
execution_times = Table(
    'execution_times', metadata,
    autoload_with=engine
)

# Crear sesión
Session = sessionmaker(bind=engine)
session = Session()

def register_info(test, language, algorithm, time):
    try:
        with session.begin():
            insert_statement = execution_times.insert().values(test=test, language=language, algorithm=algorithm, time=time)
            session.execute(insert_statement)
        return True
    except SQLAlchemyError as e:
        print(f"Error al crear un nuevo registro en la base de datos: {e}")
        return False
