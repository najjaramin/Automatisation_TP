
$logDir = "C:\Users\Public\TP_Automatisation\logs"  # change selon ton dossier
$days = 7

if (-Not (Test-Path $logDir)) {
    Write-Host "Le dossier $logDir n'existe pas"
    exit
}

Get-ChildItem $logDir -File | Where-Object { $_.LastWriteTime -lt (Get-Date).AddDays(-$days) } | ForEach-Object {
    Remove-Item $_.FullName -Force
    Write-Host "Supprimé: $($_.FullName)"
}
