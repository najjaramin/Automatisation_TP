import smtplib
from email.mime.text import MIMEText

smtp_server = "smtp.gmail.com"
port = 587
sender = "aminnajjar243@gmail.com"
password = "qfpjnpuoetrappea"  
receiver = "aminnajjar3@gmail.com"
subject = "Alerte"
message = "Ceci est un test d'alerte"

msg = MIMEText(message)
msg["Subject"] = subject
msg["From"] = sender
msg["To"] = receiver

try:
    server = smtplib.SMTP(smtp_server, port)
    server.starttls()
    server.login(sender, password)
    server.sendmail(sender, receiver, msg.as_string())
    server.quit()
    print("Mail envoyé")
except Exception as e:
    print("Erreur :", e)
