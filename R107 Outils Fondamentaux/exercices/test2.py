import smtplib
from email.mime.text import MIMEText
import time

def send_email(to_addr, subject, body):
    smtp_server = "smtp.gmail.com"
    smtp_port = 22
    
    # REMPLACEZ CECI PAR VOS NOUVEAUX IDENTIFIANTS
    username = "karistoalabeat@gmail.com"
    password = "kgdmofmcqofxfivj" 

    msg = MIMEText(body, "plain", "utf-8")
    msg["Subject"] = subject
    msg["From"] = username
    msg["To"] = to_addr

    try:
        print(f"Tentative de connexion au serveur {smtp_server}...")
        with smtplib.SMTP(smtp_server, smtp_port) as server:
            server.set_debuglevel(1)  # Affiche les détails de la connexion dans la console
            server.starttls()  # Sécurise la connexion
            server.login(username, password)
            server.send_message(msg)
            print(f"✅ Email envoyé avec succès à {to_addr}")
            
    except smtplib.SMTPAuthenticationError:
        print("❌ Erreur d'authentification : Vérifiez votre email et votre mot de passe d'application.")
    except Exception as e:
        print(f"❌ Une erreur s'est produite : {e}")

# Test direct (sans la boucle infinie pour commencer)
print("Démarrage du script...")
send_email("noammsakni@gmail.com", "Test Rapport", "Ceci est un test de connexion.")

# Une fois que le test ci-dessus fonctionne, vous pouvez décommenter la boucle :
# while True:
#     send_email("noammsakni@gmail.com", "Rapport", "Contenu du rapport…")
#     print("Pause d'une heure...")
#     time.sleep(3600)