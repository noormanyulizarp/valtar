#!/bin/bash

export_history() {
  # Export Fish history
  cp .local/share/fish/fish_history fish_history_backup
  echo "Fish history exported to fish_history_backup"
}

import_history() {
  # Import Fish history
  cp fish_history_backup .local/share/fish/fish_history
  echo "Fish history imported"
}

case "$1" in
  "create")
    export_history
    ;;
  "import")
    import_history
    ;;
  *)
    echo "Usage: $0 {create|import}"
    exit 1
    ;;
esac
