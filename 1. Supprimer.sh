

LOG_DIR="/c/Users/Public/TP_Automatisation/test_logs"
mkdir -p "$LOG_DIR"
find "$LOG_DIR" -type f -mtime +7 -print -exec rm -f {} \;
