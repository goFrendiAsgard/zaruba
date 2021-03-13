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
        funky_name = row[5] if len(row) > 5 else task_name
        print('{timestamp} {output_type_icon} {command_type_icon} {funky_name} {log}'.format(
            timestamp=timestamp,
            output_type_icon=output_type_icon,
            command_type_icon=command_type_icon,
            funky_name=funky_name,
            log=log
        ))
    f.close()
    

if __name__ == '__main__':
    show_log()