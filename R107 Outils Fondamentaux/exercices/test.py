import smtplib
from email.mime.text import MIMEText

def send_email(to_addr, subject, body):
    smtp_server = "smtp.gmail.com"
    smtp_port = 587
    username = "karistoalabeat@gmail.com"
    password = "kgdmofmcqofxfivj"  # mot de passe d’application

    msg = MIMEText(body, "plain", "utf-8")
    msg["Subject"] = subject
    msg["From"] = username
    msg["To"] = to_addr

    with smtplib.SMTP(smtp_server, smtp_port) as server:
        server.starttls()
        server.login(username, password)
        server.send_message(msg)

import time

while True:
    send_email("noammsakni@gmail.com", "Rapport", "Contenu du rapport…")
    #send_email("luzrobertrand@gmail.com", "Rapport", "Contenu du rapport…")
    time.sleep(10)  # une fois par heure

