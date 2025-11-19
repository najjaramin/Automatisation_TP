$source = "C:\Windows\System32"  # exemple de dossier à sauvegarder
$destination = "C:\Users\Public\TP_Automatisation\backup.zip"

if (-Not (Test-Path $source)) {
    Write-Host "Le dossier $source n'existe pas"
    exit
}

Compress-Archive -Path $source -DestinationPath $destination -Force
Write-Host "Archive créée : $destination"
