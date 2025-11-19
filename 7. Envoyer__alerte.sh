#!/bin/bash
powershell.exe -Command "Send-MailMessage -From 'tonmail@gmail.com' -To 'destinataire@gmail.com' -Subject 'Alerte' -Body 'Ceci est un test' -SmtpServer 'smtp.gmail.com' -Port 587"
