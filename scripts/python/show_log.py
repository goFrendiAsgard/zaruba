from helper import cli
import csv, re

@cli
def show_log(log_file: str = 'log.zaruba.csv', pattern: str = '.*'):
    f = open(log_file, 'r')
    csv_reader = csv.reader(f, delimiter=',', quotechar='"')
    matched_dict = {}
    for row in csv_reader:
        if row[3] not in matched_dict:
            matched_dict[row[3]] = re.search(pattern, row[3], re.IGNORECASE)
        if not matched_dict[row[3]]:
            continue
        timestamp, task_name, log = row[0][:23], row[3], row[4]
        output_type_icon = "🔥" if row[1] == "ERR" else "  "
        command_type_icon = "🚀" if row[2] == "START" else "🔎"
        print('{output_type_icon} {command_type_icon} {task_name} \033[2m{timestamp}\033[0m {log}'.format(
            output_type_icon=output_type_icon,
            command_type_icon=command_type_icon,
            task_name=task_name.ljust(17),
            timestamp=timestamp,
            log=log
        ))
    f.close()
    

if __name__ == '__main__':
    show_log()