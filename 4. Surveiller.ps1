$threshold = 10  # pourcentage

Get-PSDrive -PSProvider FileSystem | ForEach-Object {
    $freePercent = ($_.Free / ($_.Used + $_.Free)) * 100
    if ($freePercent -lt $threshold) {
        Write-Host "ALERTE : Espace disque faible sur $_.Name ($([math]::Round($freePercent,2))%)"
    } else {
        Write-Host "Espace disque sur $_.Name : $([math]::Round($freePercent,2))%"
    }
}
