import sys
import re
import csv

# USAGE
# python create_service.py <log_file> <task>

def show_log(log_file: str, pattern: str):
    with open(log_file, 'r') as f:
        csv_reader = csv.reader(f, delimiter=',', quotechar='"')
        for row in csv_reader:
            if re.match(pattern, row[3]):
                timestamp = row[0]
                output_type = row[1]
                command_type = row[2]
                task = row[3]
                log = row[4]
                funkyName = row[5] if len(row) > 4 else task
                timestamp_parts = timestamp.split(' ')
                timestamp = ' '.join([timestamp_parts[0], timestamp_parts[1]]).ljust(27)
                print('  '.join([timestamp, output_type, command_type, funkyName, log]))


if __name__ == '__main__':
    log_file = sys.argv[1] if len(sys.argv) > 1 else 'log.zaruba.csv'
    pattern = sys.argv[2] if len(sys.argv) > 2 else '.*'
    show_log(log_file, pattern)


