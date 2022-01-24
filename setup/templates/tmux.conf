set -g @plugin 'tmux-plugins/tpm'
set -g @plugin 'tmux-plugins/tmux-sensible'
run '~/.tmux/plugins/tpm/tpm'

set-option -sg escape-time 10
set-option -g focus-events On
set-option -g default-terminal "screen-256color"
set-option -sa terminal-overrides ',xterm-256color:RGB'

bind c new-window -c "#{pane_current_path}"
bind '"' split-window -c "#{pane_current_path}"
bind % split-window -h -c "#{pane_current_path}"
