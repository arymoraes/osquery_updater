#!/bin/bash

# Define file paths based on OS
declare -A softwareUpdateDataFiles=(
  ["windows"]="$HOME/AppData/Local/osquerything/data/SoftwareUpdate.sqlite"
  ["darwin"]="$HOME/Library/Application Support/osquerything/Data/SoftwareUpdate.sqlite"
  ["linux"]="/opt/osquerything/SoftwareUpdate.sqlite"
)

# Determine the OS
OS="$(uname -s)"

# Set file path based on OS
case "$OS" in
  Linux*)     filePath="${softwareUpdateDataFiles["linux"]}" ;;
  Darwin*)    filePath="${softwareUpdateDataFiles["darwin"]}" ;;
  CYGWIN*|MINGW32*|MSYS*|MINGW*) filePath="${softwareUpdateDataFiles["windows"]}" ;;
  *)          echo "Unsupported OS: $OS"; exit 1 ;;
esac

echo "Selected path: $filePath"

# Create directory if it doesn't exist
directory=$(dirname "$filePath")
if [ ! -d "$directory" ]; then
  sudo mkdir -p "$directory"
  echo "Created directory: $directory"
fi

# Create the SQLite database file if it doesn't exist
if [ ! -f "$filePath" ]; then
  sudo touch "$filePath"
  echo "Created file: $filePath"
else
  echo "File already exists: $filePath"
fi

# Create the software_update table in the SQLite database
create_table_sql="CREATE TABLE IF NOT EXISTS software_updates (
  version TEXT,
  last_updated_at TEXT
);"

# Run the SQL command to create the table
echo "$create_table_sql" | sudo sqlite3 "$filePath"

echo "Table 'software_updates' created (if it didn't exist already) in $filePath."
