$folder = "C:\Users\Public\TP_Automatisation"

Get-ChildItem -Path $folder -File | Sort-Object Length -Descending | Select-Object -First 5 | ForEach-Object {
    Write-Host "$($_.Name) - $([math]::Round($_.Length/1MB,2)) MB"
}
