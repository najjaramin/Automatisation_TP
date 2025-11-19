$Action = New-ScheduledTaskAction -Execute "powershell.exe" -Argument "-File C:\Users\Public\TP_Automatisation\1.Supprimer.ps1"
$Trigger = New-ScheduledTaskTrigger -Daily -At 2am
$TaskName = "TP_Automatisation_Supprimer"

Register-ScheduledTask -Action $Action -Trigger $Trigger -TaskName $TaskName -Description "Supprime les logs anciens" -Force
Write-Host "Tâche planifiée créée : $TaskName"
