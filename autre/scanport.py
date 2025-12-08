import socket
import threading
from queue import Queue
import time

# --- Configuration ---
TARGET = '127.0.0.1'  # Scanner toujours ta machine locale par défaut (localhost)
START_PORT = 1
END_PORT = 1024       # On ne scanne que les ports bien connus (0 à 1023) pour la rapidité
MAX_THREADS = 20      # Nombre de tentatives de connexion simultanées
TIMEOUT = 1           # Temps max en secondes pour attendre une réponse

# --- File d'attente des ports à scanner ---
q = Queue()
open_ports = []
print_lock = threading.Lock()

def port_scan(port):
    """Tente de se connecter au port spécifié."""
    try:
        # Création d'un socket TCP/IP
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.settimeout(TIMEOUT)
        
        # Tenter la connexion
        result = s.connect_ex((TARGET, port))
        
        # Le code 0 signifie que la connexion est réussie (port ouvert)
        if result == 0:
            # Récupérer le nom du service si possible
            try:
                service = socket.getservbyport(port, "tcp")
            except OSError:
                service = "inconnu"
            
            with print_lock:
                print(f"✅ Port {port} ouvert ({service})")
            open_ports.append(port)
            
        s.close()
    except Exception as e:
        # print(f"Erreur lors du scan du port {port}: {e}")
        pass # Ignorer les erreurs de connexion

def worker():
    """Fonction exécutée par chaque thread."""
    while True:
        # Récupérer un port de la file d'attente
        port = q.get()
        port_scan(port)
        q.task_done()

def run_scanner():
    """Initialise les threads et la file d'attente."""
    print(f"Démarrage du scan sur {TARGET} de {START_PORT} à {END_PORT}...")
    start_time = time.time()

    # 1. Créer et démarrer les threads
    for _ in range(MAX_THREADS):
        t = threading.Thread(target=worker, daemon=True)
        t.start()

    # 2. Remplir la file d'attente avec les ports
    for port in range(START_PORT, END_PORT + 1):
        q.put(port)

    # 3. Attendre que tous les ports aient été traités
    q.join()
    end_time = time.time()

    # --- Résultat ---
    print("\n--- Scan Terminé ---")
    if open_ports:
        print(f"Ports ouverts trouvés: {sorted(open_ports)}")
    else:
        print("Aucun port ouvert trouvé dans la plage spécifiée.")
    print(f"Durée totale: {end_time - start_time:.2f} secondes.")


if __name__ == "__main__":
    run_scanner()