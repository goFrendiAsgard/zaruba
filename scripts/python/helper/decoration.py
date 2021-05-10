import random

normal="\033[0m"
bold="\033[1m"
faint="\033[2m"
italic="\033[3m"
underline="\033[4m"
blinkSlow="\033[5m"
blinkRapid="\033[6m"
inverse="\033[7m"
conceal="\033[8m"
crossedOut="\033[9m"
black="\033[30m"
red="\033[31m"
green="\033[32m"
yellow="\033[33m"
blue="\033[34m"
magenta="\033[35m"
cyan="\033[36m"
white="\033[37m"
bgBlack="\033[40m"
bgRed="\033[41m"
bgGreen="\033[42m"
bgYellow="\033[43m"
bgBlue="\033[44m"
bgMagenta="\033[45m"
bgCyan="\033[46m"
bgWhite="\033[47m"
noStyle="\033[0m"
noUnderline="\033[24m"
noInverse="\033[27m"
noColor="\033[39m"


def generate_icon() -> str:
    icon_list = ['🥜', '🍄', '🌰', '🍞', '🥐', '🥖', '🥞', '🧀', '🍖', '🍗', '🥓', '🍔', '🍟', '🍕', '🌭', '🌮', '🌯', '🥙', '🍲', '🥗', '🍿']
    index = random.randrange(0, len(icon_list))
    return icon_list[index]
