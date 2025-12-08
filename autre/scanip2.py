from scapy.all import Ether, ARP, srp
import sys

def scan_reseau(ip_plage):
    """
    Scanne le réseau local en envoyant des requêtes ARP
    à toutes les adresses IP de la plage spécifiée.
    """
    print(f"[*] Balayage du réseau : {ip_plage}...")

    # 1. Création du paquet ARP (couche 2)
    # pdst est la cible (toutes les adresses dans la plage)
    arp_request = ARP(pdst=ip_plage)
    
    # 2. Création du paquet Ethernet (couche 1)
    # L'adresse MAC de destination (dst) est "ff:ff:ff:ff:ff:ff" (broadcast)
    # pour s'assurer que toutes les machines reçoivent la requête.
    ether_frame = Ether(dst="ff:ff:ff:ff:ff:ff")
    
    # 3. Combinaison des deux couches pour former le paquet final
    packet = ether_frame / arp_request

    # 4. Envoi et réception des paquets
    # timeout: attendre 1 seconde maximum pour les réponses
    # verbose=False: n'affiche pas les paquets envoyés/reçus par défaut
    resultat = srp(packet, timeout=1, verbose=False)[0]

    machines_actives = []
    
    # 5. Extraction des résultats
    for sent, received in resultat:
        # received.psrc est l'adresse IP de la machine qui a répondu
        # received.hwsrc est l'adresse MAC de la machine qui a répondu
        machines_actives.append({'ip': received.psrc, 'mac': received.hwsrc})
        
    return machines_actives

# --- Configuration ---
# REMPLACEZ cette plage par celle de votre réseau. 
# Si votre IP est 192.168.1.X, utilisez 192.168.1.0/24
# Si votre IP est 10.0.0.X, utilisez 10.0.0.0/24
PLAGE_RESEAU = "192.168.1.0/24" 
# ---------------------

try:
    hosts = scan_reseau(PLAGE_RESEAU)
    
    print("-" * 30)
    print("Machines Actives Trouvées :")
    print("-" * 30)
    print("{:<16} {:<17}".format("IP", "MAC"))
    print("-" * 30)
    for host in hosts:
        print("{:<16} {:<17}".format(host['ip'], host['mac']))
    print("-" * 30)

except PermissionError:
    print("\n[!] Erreur de permission. Sur Linux/macOS, vous pourriez avoir besoin de l'exécuter avec 'sudo'.")
    print("    Exemple : sudo python votre_script.py")
except Exception as e:
    print(f"\n[!] Une erreur inattendue est survenue : {e}")