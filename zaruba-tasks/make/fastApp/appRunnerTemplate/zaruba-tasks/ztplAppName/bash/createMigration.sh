alembic upgrade head
alembic revision --autogenerate -m "$("${ZARUBA_BIN}" str currentTime) $("${ZARUBA_BIN}" str toSnake $("${ZARUBA_BIN}" str newName))"