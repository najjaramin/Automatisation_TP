#!/bin/bash
SOURCE_DIR="/c/Users/Public/TP_Automatisation/test_zip"
mkdir -p "$SOURCE_DIR"
# créer fichiers test
for i in {0..2}; do echo "test" > "$SOURCE_DIR/file$i.txt"; done
zip -r "/c/Users/Public/TP_Automatisation/backup.zip" "$SOURCE_DIR"
echo "Archive créée : /c/Users/Public/TP_Automatisation/backup.zip"
