
SERVICE="Spooler"
echo "Redémarrage du service $SERVICE..."
powershell.exe -Command "Restart-Service -Name '$SERVICE' -Force"
echo "Service redémarré."
