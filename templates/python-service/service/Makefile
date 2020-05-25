test:
	(conda env create --file=environment.yml || conda env update --file=environment.yml  --prune || echo "Failed. Keep going.") && python -m pytest
run:
	(conda env create --file=environment.yml || conda env update --file=environment.yml  --prune || echo "Failed. Keep going.") && python main.py
