$user = "AMIN"  # nom de l'utilisateur Windows

$count = Get-Process | Where-Object { $_.StartInfo.Username -eq $user } | Measure-Object | Select-Object -ExpandProperty Count
Write-Host "Nombre de processus pour $user : $count"
