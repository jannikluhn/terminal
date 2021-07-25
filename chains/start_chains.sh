#!/usr/bin/env bash

# Run geth
set -eu

SESSION=chains
SINK_WINDOW=sink
SOURCE_WINDOW=source

# setup tmux session
tmux new-session -d -s ${SESSION} -n ${SINK_WINDOW}
tmux new-window -t ${SESSION}: -n ${SOURCE_WINDOW}

tmux send-keys -t ${SESSION}:${SINK_WINDOW} "SHROOT=. ./start_sink_chain.sh" C-m
tmux send-keys -t ${SESSION}:${SOURCE_WINDOW} "SHROOT=. ./start_source_chain.sh" C-m
tmux attach -t ${SESSION}
