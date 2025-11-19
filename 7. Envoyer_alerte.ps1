$From = "aminnajjar243@gmail.com"
$To = "aminnajjar3e@gmail.com"
$Subject = "Alerte"
$Body = "Ceci est un test d'alerte"
$SMTP = "smtp.gmail.com"
$Port = 587
$Password = "qfpjnpuoetrappea"  # générer un mot de passe d'application Gmail

$SecurePassword = ConvertTo-SecureString $Password -AsPlainText -Force
$Cred = New-Object System.Management.Automation.PSCredential($From, $SecurePassword)

Send-MailMessage -From $From -To $To -Subject $Subject -Body $Body -SmtpServer $SMTP -Port $Port -UseSsl -Credential $Cred
Write-Host "Mail envoyé de $From à $To"
