import sys
import re
import csv
import traceback

# USAGE
# python create_service.py <log_file> <task>

def show_log(log_file: str, pattern: str):
    try:
        f = open(log_file, 'r')
        csv_reader = csv.reader(f, delimiter=',', quotechar='"')
        for row in csv_reader:
            if not re.match(pattern, row[3]):
                continue
            timestamp = row[0]
            output_type_icon = "🔥" if row[1] == "ERR" else "  "
            command_type_icon = "🚀" if row[2] == "START" else "🔎"
            task = row[3]
            log = row[4]
            funkyName = row[5] if len(row) > 5 else task
            timestamp_parts = timestamp.split(' ')
            timestamp = ' '.join([timestamp_parts[0], timestamp_parts[1]]).ljust(27)
            print(' '.join([timestamp, output_type_icon, command_type_icon, funkyName, log]))
        f.close()
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)


if __name__ == '__main__':
    log_file = sys.argv[1] if len(sys.argv) > 1 and sys.argv[1] != '' else 'log.zaruba.csv'
    pattern = sys.argv[2] if len(sys.argv) > 2 and sys.argv[2] != '' else '.*'
    show_log(log_file, pattern)


