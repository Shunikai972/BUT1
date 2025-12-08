import socket
import threading
import pickle
import time
import random
import pygame

# --- Configuration Réseau et Jeu ---
PORT = 22
MAX_CLIENTS = 4

# Configuration Pygame pour l'affichage
WIDTH, HEIGHT = 600, 600
GRID_SIZE = 20
GRID_WIDTH = WIDTH // GRID_SIZE
GRID_HEIGHT = HEIGHT // GRID_SIZE

# Couleurs
BLACK = (0, 0, 0)
GREEN_SELF = (0, 255, 0)
BLUE_OTHER = (0, 0, 255)
RED_FOOD = (255, 0, 0)

# --- État Global du Jeu (Géré par le Serveur) ---
game_state = {
    'snakes': {},        # {id: [(x1, y1), ...], ...}
    'directions': {},    # {id: 'RIGHT', ...}
    'food': (10, 10),
    'scores': {}
}
clients = {} # {id: conn_socket, ...}
current_id = 0
server_lock = threading.Lock() # Pour sécuriser l'accès à game_state
game_active = False

# --- Fonctions Utilitaires ---

def get_local_ip():
    """Récupère l'IP locale de la machine pour l'Host."""
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    try:
        # Tente de se connecter (sans envoyer de données) pour obtenir l'IP de l'interface
        s.connect(('10.255.255.255', 1))
        IP = s.getsockname()[0]
    except Exception:
        IP = '127.0.0.1'
    finally:
        s.close()
    return IP

def generate_new_food(snakes):
    """Génère une nouvelle position de nourriture qui n'est pas sur un serpent."""
    occupied = set()
    for snake in snakes.values():
        for segment in snake:
            occupied.add(segment)
            
    while True:
        x = random.randint(0, GRID_WIDTH - 1)
        y = random.randint(0, GRID_HEIGHT - 1)
        new_pos = (x, y)
        if new_pos not in occupied:
            return new_pos

# --- Logique Serveur (Host) ---

def apply_move(snake, direction):
    """Calcule la nouvelle tête du serpent en fonction de la direction."""
    head_x, head_y = snake[0]
    
    if direction == 'UP':
        new_head = (head_x, head_y - 1)
    elif direction == 'DOWN':
        new_head = (head_x, head_y + 1)
    elif direction == 'LEFT':
        new_head = (head_x - 1, head_y)
    elif direction == 'RIGHT':
        new_head = (head_x + 1, head_y)
    else: # Ne devrait pas arriver
        return snake[0] 
        
    # Gestion du wrapping (passage à travers les murs)
    new_x = new_head[0] % GRID_WIDTH
    new_y = new_head[1] % GRID_HEIGHT
    
    return (new_x, new_y)

def update_game_state():
    """Logique principale du jeu, exécutée par le Serveur."""
    global game_state
    
    with server_lock:
        
        snakes = game_state['snakes']
        directions = game_state['directions']
        food_pos = game_state['food']
        
        # 1. Faire avancer tous les serpents
        for id, snake in list(snakes.items()):
            if id not in directions:
                continue # Serpent sans direction connue (ne bouge pas encore)
                
            direction = directions[id]
            new_head = apply_move(snake, direction)
            
            # Vérification des collisions (mur et auto-collision gérées par wrapping/head-tail)
            
            # 2. Vérification de la nourriture
            if new_head == food_pos:
                # Mange la nourriture (pas de retrait du dernier segment)
                snakes[id].insert(0, new_head)
                game_state['scores'][id] = game_state['scores'].get(id, 0) + 1
                
                # Génère une nouvelle nourriture
                game_state['food'] = generate_new_food(snakes)
            else:
                # Mouvement normal (ajoute la tête, retire la queue)
                snakes[id].insert(0, new_head)
                snakes[id].pop()
            
            # --- Collision avec les autres serpents ou auto-collision ---
            # Si le serpent entre en collision avec lui-même ou un autre serpent, il perd une partie de sa longueur
            all_segments = set()
            for seg_id, seg_snake in snakes.items():
                if seg_id != id:
                    all_segments.update(seg_snake)
                else:
                    # Auto-collision (la tête est déjà le premier segment)
                    for segment in seg_snake[1:]:
                        all_segments.add(segment)
            
            if new_head in all_segments:
                 # Simplification: le serpent meurt et redémarre
                snakes[id] = [(random.randint(5, GRID_WIDTH-5), random.randint(5, GRID_HEIGHT-5))]
                game_state['directions'][id] = 'RIGHT' # Direction par défaut
                game_state['scores'][id] = max(0, game_state['scores'].get(id, 0) - 1)


def main_game_loop():
    """Boucle principale du serveur, gère la synchronisation du jeu."""
    global game_active
    # Le jeu se met à jour environ 8 fois par seconde
    TICK_RATE = 0.125 
    
    while game_active:
        time.sleep(TICK_RATE) 
        
        # 1. Mise à jour de l'état du jeu (mouvement, collisions, nourriture)
        update_game_state()
        
        # 2. Envoi de l'état complet du jeu à tous les clients
        data_to_send = pickle.dumps(game_state)
        
        with server_lock:
            for client_id, conn in list(clients.items()): 
                try:
                    conn.sendall(data_to_send)
                except:
                    # Si l'envoi échoue, le client est probablement déconnecté
                    pass 


def handle_client(conn, client_id):
    """Thread dédié pour recevoir les commandes d'un client."""
    global game_state
    
    # 1. Envoi de l'ID au client
    conn.send(str(client_id).encode())
    
    # 2. Boucle de réception des mouvements
    while True:
        try:
            # Reçoit la direction du client (ex: 'UP')
            data = conn.recv(2048).decode()
            if not data:
                break 
            
            # Met à jour la direction souhaitée par le client
            with server_lock:
                 # Le serveur vérifie qu'on ne fait pas un demi-tour immédiat
                if data == 'UP' and game_state['directions'].get(client_id) != 'DOWN':
                    game_state['directions'][client_id] = data
                elif data == 'DOWN' and game_state['directions'].get(client_id) != 'UP':
                    game_state['directions'][client_id] = data
                elif data == 'LEFT' and game_state['directions'].get(client_id) != 'RIGHT':
                    game_state['directions'][client_id] = data
                elif data == 'RIGHT' and game_state['directions'].get(client_id) != 'LEFT':
                    game_state['directions'][client_id] = data
                
        except:
            break

    # Nettoyage si le client se déconnecte
    print(f"Client {client_id} déconnecté.")
    with server_lock:
        if client_id in clients: del clients[client_id]
        if client_id in game_state['snakes']: del game_state['snakes'][client_id]
        if client_id in game_state['directions']: del game_state['directions'][client_id]
        if client_id in game_state['scores']: del game_state['scores'][client_id]
    conn.close()


def start_server(host_ip):
    """Démarre le socket du serveur et attend les clients."""
    global current_id, game_active
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    
    try:
        s.bind((host_ip, PORT))
        s.listen(MAX_CLIENTS)
    except Exception as e:
        print(f"Erreur de liaison: {e}. Vérifiez que le port {PORT} n'est pas déjà utilisé.")
        return

    print(f"\n--- SERVEUR DÉMARRÉ ---")
    print(f"Hôte: {host_ip}:{PORT}. En attente de connexions (Max {MAX_CLIENTS})...")
    
    game_active = True
    # Démarrer la boucle de jeu dans un thread séparé
    threading.Thread(target=main_game_loop, daemon=True).start()
    
    try:
        while True:
            conn, addr = s.accept()
            print(f"Client connecté: {addr}")
            
            # Initialise le nouveau serpent dans l'état du jeu
            initial_pos = (random.randint(5, GRID_WIDTH-5), random.randint(5, GRID_HEIGHT-5))
            with server_lock:
                game_state['snakes'][current_id] = [initial_pos]
                game_state['directions'][current_id] = 'RIGHT'
                game_state['scores'][current_id] = 0

            # Stocke la connexion et démarre un thread pour la gérer
            clients[current_id] = conn
            threading.Thread(target=handle_client, args=(conn, current_id), daemon=True).start()
            current_id += 1
    except KeyboardInterrupt:
        print("\nServeur arrêté.")
    finally:
        game_active = False
        s.close()
        
# --- Logique Client (Joueur) ---

my_id = None
client_socket = None
# Ces variables sont mises à jour par le thread de réception
client_snakes = {}
client_food = (10, 10)
client_scores = {}

def connect_to_server(server_ip):
    """Établit la connexion et reçoit l'ID unique."""
    global my_id, client_socket
    try:
        client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        client_socket.connect((server_ip, PORT))
        
        # Réception de notre ID
        my_id = int(client_socket.recv(1024).decode())
        print(f"Connecté au Serveur. Mon ID de joueur: {my_id}")
        return client_socket
    except Exception as e:
        print(f"Impossible de se connecter au serveur {server_ip}:{PORT}. Erreur: {e}")
        return None

def send_move(move):
    """Envoie la direction de ce client au serveur."""
    if client_socket:
        try:
            client_socket.send(move.encode())
        except:
            pass # Le thread de réception gérera la déconnexion

def receive_game_state():
    """Thread pour recevoir en continu l'état du jeu du serveur."""
    global client_snakes, client_food, client_scores
    
    if not client_socket: return

    while True:
        try:
            data = client_socket.recv(4096) 
            if data:
                # Décrypter l'état du jeu
                game_state_received = pickle.loads(data)
                client_snakes = game_state_received.get('snakes', {})
                client_food = game_state_received.get('food', (10, 10))
                client_scores = game_state_received.get('scores', {})
            else:
                break # Le serveur s'est déconnecté
        except:
            break
            
    print("Déconnexion du serveur.")
    pygame.quit()


def draw_game(screen):
    """Dessine tous les éléments sur l'écran (Serpents, Nourriture, Score)."""
    screen.fill(BLACK) 
    
    # Dessine la nourriture
    fx, fy = client_food
    food_rect = pygame.Rect(fx * GRID_SIZE, fy * GRID_SIZE, GRID_SIZE, GRID_SIZE)
    pygame.draw.rect(screen, RED_FOOD, food_rect)

    # Dessine tous les serpents
    for id, snake in client_snakes.items():
        color = GREEN_SELF if id == my_id else BLUE_OTHER 
        # Dessin des segments
        for x, y in snake:
            snake_rect = pygame.Rect(x * GRID_SIZE, y * GRID_SIZE, GRID_SIZE, GRID_SIZE)
            pygame.draw.rect(screen, color, snake_rect)

    # Affichage du score
    font = pygame.font.Font(None, 24)
    y_offset = 10
    score_text = font.render(f"Mon Score: {client_scores.get(my_id, 0)}", True, (255, 255, 255))
    screen.blit(score_text, (WIDTH - 150, y_offset))
    
    pygame.display.update()


def main_client(server_ip):
    """Démarre le jeu en mode Client."""
    pygame.init()
    screen = pygame.display.set_mode((WIDTH, HEIGHT))
    pygame.display.set_caption("Multiplayer Snake LAN")
    
    if not connect_to_server(server_ip):
        return

    # Démarrer le thread de réception de l'état du jeu
    threading.Thread(target=receive_game_state, daemon=True).start()
    
    running = True
    clock = pygame.time.Clock()

    while running:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            
            # Gestion des touches pour envoyer la direction au serveur
            if event.type == pygame.KEYDOWN:
                move = None
                if event.key == pygame.K_UP:
                    move = 'UP'
                elif event.key == pygame.K_DOWN:
                    move = 'DOWN'
                elif event.key == pygame.K_LEFT:
                    move = 'LEFT'
                elif event.key == pygame.K_RIGHT:
                    move = 'RIGHT'
                
                if move:
                    send_move(move) 

        draw_game(screen)
        clock.tick(60) # Limite le FPS d'affichage

    if client_socket:
        client_socket.close()
    pygame.quit()


# --- Fonction de Démarrage ---

def main():
    print("\n--- Jeu Python Snake Multijoueur LAN ---")
    
    local_ip = get_local_ip()
    print(f"⚠️ Votre IP locale (LAN) est: **{local_ip}** (à communiquer si vous êtes Hôte)")

    # Demander le rôle
    role = input("\nVoulez-vous (H)éberger le jeu ou (J)oindre une partie ? [H/J] : ").strip().upper()

    if role == 'H':
        print("\n**Mode HÔTE (Serveur) sélectionné.**")
        print("Lancement du Serveur...")
        start_server(local_ip)

    elif role == 'J':
        print("\n**Mode JOUEUR (Client) sélectionné.**")
        
        server_ip = input("Entrez l'adresse IP du Serveur (Host) : ").strip()
        
        if not server_ip:
            print("Erreur: Adresse IP non valide.")
            return

        print(f"Tentative de connexion à {server_ip}...")
        main_client(server_ip)

    else:
        print("Choix non valide. Fermeture du programme.")

if __name__ == "__main__":
    main()