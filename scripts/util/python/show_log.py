from common_helper import get_argv
import csv, re, sys, traceback

# USAGE
# python create_service.py <log_file> <task>

def show_log(log_file: str, pattern: str):
    f = open(log_file, 'r')
    csv_reader = csv.reader(f, delimiter=',', quotechar='"')
    matched_dict = {}
    for row in csv_reader:
        if row[3] not in matched_dict:
            matched_dict[row[3]] = re.search(pattern, row[3], re.IGNORECASE)
        if not matched_dict[row[3]]:
            continue
        timestamp, task, log = row[0][:23], row[3], row[4]
        output_type_icon = "🔥" if row[1] == "ERR" else "  "
        command_type_icon = "🚀" if row[2] == "START" else "🔎"
        funky_name = row[5] if len(row) > 5 else task
        print('{timestamp} {output_type_icon} {command_type_icon} {funky_name} {log}'.format(
            timestamp=timestamp,
            output_type_icon=output_type_icon,
            command_type_icon=command_type_icon,
            funky_name=funky_name,
            log=log
        ))
    f.close()
    

if __name__ == '__main__':
    try:
        log_file = get_argv(1, 'log.zaruba.csv')
        pattern = get_argv(2, '.*')
        show_log(log_file, pattern)
    except Exception as e:
        print(e)
        traceback.print_exc()
        sys.exit(1)