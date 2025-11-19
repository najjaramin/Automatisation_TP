#!/bin/bash
USER=$(whoami)
COUNT=$(powershell.exe -Command "(Get-Process | Where-Object {$_.StartInfo.EnvironmentVariables['USERNAME'] -eq '$USER'}).Count")
echo "$COUNT processus pour $USER"
