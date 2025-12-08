import socket

def obtenir_ip_locale():
    """
    Tente de connecter un socket à une adresse non-routable 
    (mais non connectée) pour déterminer l'interface réseau locale utilisée.
    """
    try:
        # Crée un socket de type DGRAM (UDP)
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        
        # Le socket ne se connecte pas réellement, mais utilise la table
        # de routage pour trouver l'adresse de l'interface par défaut.
        s.connect(("8.8.8.8", 80)) # Utilise un serveur DNS public comme cible arbitraire
        ip_locale = s.getsockname()[0]
        s.close()
        return ip_locale
    except Exception as e:
        return f"Erreur lors de la récupération de l'IP : {e}"

# Exécution
mon_ip = obtenir_ip_locale()
print(f"L'adresse IP locale de votre machine est : {mon_ip}")