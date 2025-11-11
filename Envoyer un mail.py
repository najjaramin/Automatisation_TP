import smtplib
from email.message import EmailMessage

msg = EmailMessage()
msg.set_content("Message d'alerte personnalisé")
msg["Subject"] = "Alerte système"
msg["From"] = "expediteur@example.com"
msg["To"] = "aminnajjar3@gmail.com"

with smtplib.SMTP("localhost") as s:
    s.send_message(msg)