$serviceName = "wuauserv"  # exemple : service Windows Update
try {
    Restart-Service -Name $serviceName -Force
    Write-Host "Service $serviceName redémarré"
} catch {
    Write-Host "Erreur lors du redémarrage de $serviceName : $_"
}
