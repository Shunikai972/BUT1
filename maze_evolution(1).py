"""
=============================================================
  MAZE CAR EVOLUTION — Algorithme génétique en pur Python
  Dépendances : pip install pygame numpy
  Lancer      : python maze_evolution.py
=============================================================
"""

import pygame
import numpy as np
import random
import math
import sys
import time

# ─────────────────────────────────────────────
#  CONFIGURATION GLOBALE
# ─────────────────────────────────────────────
SCREEN_W, SCREEN_H = 1100, 700
MAZE_W,   MAZE_H   = 700,  700
PANEL_W             = SCREEN_W - MAZE_W

CELL       = 35          # Taille d'une cellule du labyrinthe
COLS       = MAZE_W  // CELL
ROWS       = MAZE_H  // CELL

POP_SIZE       = 40      # Voitures par génération
MAX_STEPS      = 1200    # Durée max d'une run
N_ELITE        = 3       # Élites préservées sans mutation
K_TOURNAMENT   = 6       # Taille du tournoi de sélection
SIGMA_INIT     = 0.12    # Amplitude de mutation initiale
MUTATION_RATE  = 0.18    # Probabilité de muter un poids
N_RAYS         = 7       # Rayons de détection
RAY_LEN        = 130     # Longueur max des rayons (pixels)
CAR_SPEED      = 2.8     # Vitesse de déplacement
CAR_TURN       = 3.5     # Vitesse de rotation (degrés)
CAR_SIZE       = 8       # Rayon de la voiture

# Réseau neuronal : entrées → cachée → sorties
NN_LAYERS = [N_RAYS + 1, 10, 2]   # +1 pour la vitesse actuelle

FPS = 60

# ─────────────────────────────────────────────
#  COULEURS
# ─────────────────────────────────────────────
BG          = (18,  18,  24)
WALL_COL    = (55,  60,  80)
FLOOR_COL   = (28,  30,  42)
START_COL   = (40, 180, 100)
END_COL     = (220, 80,  60)
CHKP_COL    = (60, 130, 220)
CHKP_DONE   = (60, 200, 120)
CAR_ALIVE   = (255, 200,  60)
CAR_ELITE   = (255, 100, 100)
CAR_DEAD    = (80,   80,  90)
RAY_COL     = (80,  200, 255, 90)
PANEL_BG    = (22,  22,  30)
TEXT_COL    = (200, 200, 220)
TEXT_DIM    = (120, 120, 140)
ACCENT      = (100, 160, 255)

# ─────────────────────────────────────────────
#  LABYRINTHE (généré par DFS récursif)
# ─────────────────────────────────────────────
class Maze:
    def __init__(self):
        # walls[r][c] = {'N':bool, 'S':bool, 'E':bool, 'W':bool}
        self.grid = [
            [{'N': True, 'S': True, 'E': True, 'W': True}
             for _ in range(COLS)]
            for _ in range(ROWS)
        ]
        self.visited = [[False]*COLS for _ in range(ROWS)]
        self._generate(0, 0)
        self.walls = self._build_wall_rects()
        self.checkpoints = self._place_checkpoints()

    def _generate(self, r, c):
        self.visited[r][c] = True
        dirs = [('N', -1, 0), ('S', 1, 0), ('E', 0, 1), ('W', 0, -1)]
        random.shuffle(dirs)
        for d, dr, dc in dirs:
            nr, nc = r + dr, c + dc
            if 0 <= nr < ROWS and 0 <= nc < COLS and not self.visited[nr][nc]:
                # Casser le mur
                self.grid[r][c][d] = False
                opposite = {'N': 'S', 'S': 'N', 'E': 'W', 'W': 'E'}[d]
                self.grid[nr][nc][opposite] = False
                self._generate(nr, nc)

    def _build_wall_rects(self):
        rects = []
        for r in range(ROWS):
            for c in range(COLS):
                x, y = c * CELL, r * CELL
                if self.grid[r][c]['N']:
                    rects.append(pygame.Rect(x, y, CELL, 2))
                if self.grid[r][c]['W']:
                    rects.append(pygame.Rect(x, y, 2, CELL))
                if self.grid[r][c]['S']:
                    rects.append(pygame.Rect(x, y + CELL - 2, CELL, 2))
                if self.grid[r][c]['E']:
                    rects.append(pygame.Rect(x + CELL - 2, y, 2, CELL))
        return rects

    def _place_checkpoints(self):
        """Checkpoints le long d'un chemin DFS depuis (0,0) vers le coin opposé."""
        path = self._find_path(0, 0, ROWS - 1, COLS - 1)
        # Prendre 1 checkpoint tous les ~5 cellules
        step = max(1, len(path) // 12)
        cps = []
        for i in range(step, len(path), step):
            r, c = path[i]
            cx = c * CELL + CELL // 2
            cy = r * CELL + CELL // 2
            cps.append((cx, cy))
        return cps

    def _find_path(self, sr, sc, er, ec):
        """BFS pour trouver un chemin dans le labyrinthe généré."""
        from collections import deque
        queue = deque([(sr, sc, [(sr, sc)])])
        seen = set()
        seen.add((sr, sc))
        dir_map = {'N': (-1, 0), 'S': (1, 0), 'E': (0, 1), 'W': (0, -1)}
        while queue:
            r, c, path = queue.popleft()
            if r == er and c == ec:
                return path
            for d, (dr, dc) in dir_map.items():
                if not self.grid[r][c][d]:
                    nr, nc = r + dr, c + dc
                    if (nr, nc) not in seen:
                        seen.add((nr, nc))
                        queue.append((nr, nc, path + [(nr, nc)]))
        return [(sr, sc)]

    def draw(self, surface):
        surface.fill(FLOOR_COL)
        for rect in self.walls:
            pygame.draw.rect(surface, WALL_COL, rect)
        # Start
        pygame.draw.rect(surface, START_COL,
                         (2, 2, CELL - 4, CELL - 4), 2)
        # End
        er, ec = ROWS - 1, COLS - 1
        pygame.draw.rect(surface, END_COL,
                         (ec*CELL + 2, er*CELL + 2, CELL-4, CELL-4), 2)

    def draw_checkpoints(self, surface, reached_set):
        for i, (cx, cy) in enumerate(self.checkpoints):
            col = CHKP_DONE if i in reached_set else CHKP_COL
            pygame.draw.circle(surface, col, (cx, cy), 6)
            pygame.draw.circle(surface, (255,255,255), (cx, cy), 6, 1)

# ─────────────────────────────────────────────
#  RÉSEAU DE NEURONES (feedforward)
# ─────────────────────────────────────────────
class NeuralNet:
    def __init__(self, layers, weights=None):
        self.layers = layers
        if weights is not None:
            self.weights, self.biases = self._unpack(weights)
        else:
            self.weights, self.biases = self._random_init()

    def _random_init(self):
        ws, bs = [], []
        for i in range(len(self.layers) - 1):
            ws.append(np.random.randn(self.layers[i], self.layers[i+1]) * 0.5)
            bs.append(np.zeros(self.layers[i+1]))
        return ws, bs

    def genome_size(self):
        s = 0
        for i in range(len(self.layers) - 1):
            s += self.layers[i] * self.layers[i+1] + self.layers[i+1]
        return s

    def _unpack(self, flat):
        ws, bs = [], []
        idx = 0
        for i in range(len(self.layers) - 1):
            n_w = self.layers[i] * self.layers[i+1]
            n_b = self.layers[i+1]
            ws.append(flat[idx:idx+n_w].reshape(self.layers[i], self.layers[i+1]))
            idx += n_w
            bs.append(flat[idx:idx+n_b])
            idx += n_b
        return ws, bs

    def get_genome(self):
        parts = []
        for w, b in zip(self.weights, self.biases):
            parts.append(w.flatten())
            parts.append(b)
        return np.concatenate(parts)

    def forward(self, x):
        out = np.array(x, dtype=float)
        for i, (w, b) in enumerate(zip(self.weights, self.biases)):
            out = out @ w + b
            if i < len(self.weights) - 1:
                out = np.tanh(out)
            else:
                out = np.tanh(out)  # Sorties dans [-1, 1]
        return out

# ─────────────────────────────────────────────
#  VOITURE
# ─────────────────────────────────────────────
class Car:
    def __init__(self, genome=None):
        self.brain     = NeuralNet(NN_LAYERS, genome)
        self.reset()
        self.is_elite  = False

    def reset(self):
        # Départ au centre de la cellule (0,0)
        self.x        = CELL / 2
        self.y        = CELL / 2
        self.angle    = 0.0     # degrés, 0 = droite
        self.speed    = 0.0
        self.alive    = True
        self.finished = False
        self.score    = 0.0
        self.steps    = 0
        self.checkpoints_reached = set()
        self.collisions = 0
        self.dist_traveled = 0.0
        self.last_x   = self.x
        self.last_y   = self.y
        self.ray_distances = [1.0] * N_RAYS

    def cast_rays(self, maze_surface, walls):
        """Lance N_RAYS rayons et retourne distances normalisées [0,1]."""
        distances = []
        spread    = 160   # Angle total couvert (°)
        start_ang = self.angle - spread / 2

        for i in range(N_RAYS):
            ray_ang = math.radians(start_ang + i * spread / (N_RAYS - 1))
            hit_dist = RAY_LEN
            dx = math.cos(ray_ang)
            dy = math.sin(ray_ang)

            # Avancer pixel par pixel (simplifié, assez rapide pour des runs courtes)
            for t in range(1, RAY_LEN, 3):
                rx = self.x + dx * t
                ry = self.y + dy * t
                if rx < 0 or rx >= MAZE_W or ry < 0 or ry >= MAZE_H:
                    hit_dist = t
                    break
                # Vérifier collision mur
                pt = pygame.Rect(rx - 1, ry - 1, 2, 2)
                if any(pt.colliderect(w) for w in walls):
                    hit_dist = t
                    break

            distances.append(hit_dist / RAY_LEN)

        self.ray_distances = distances
        return distances

    def update(self, walls, checkpoints):
        if not self.alive or self.finished:
            return

        self.steps += 1

        # Rayons → réseau → commandes
        inputs    = self.ray_distances + [self.speed / CAR_SPEED]
        output    = self.brain.forward(inputs)
        accel     = float(output[0])   # [-1, 1]
        steering  = float(output[1])   # [-1, 1]

        # Physique
        self.speed  = max(0.2, min(CAR_SPEED, self.speed + accel * 0.3))
        self.angle += steering * CAR_TURN

        rad = math.radians(self.angle)
        nx  = self.x + math.cos(rad) * self.speed
        ny  = self.y + math.sin(rad) * self.speed

        # Collision murs
        car_rect = pygame.Rect(nx - CAR_SIZE, ny - CAR_SIZE,
                               CAR_SIZE*2, CAR_SIZE*2)
        if any(car_rect.colliderect(w) for w in walls):
            self.collisions += 1
            if self.collisions > 5:
                self.alive = False
                return
            # Ne pas avancer
        else:
            self.dist_traveled += math.hypot(nx - self.x, ny - self.y)
            self.x, self.y = nx, ny

        # Checkpoints
        for i, (cx, cy) in enumerate(checkpoints):
            if i not in self.checkpoints_reached:
                if math.hypot(self.x - cx, self.y - cy) < CELL * 0.6:
                    self.checkpoints_reached.add(i)

        # Arrivée (coin bas-droit)
        end_x = (COLS - 0.5) * CELL
        end_y = (ROWS - 0.5) * CELL
        if math.hypot(self.x - end_x, self.y - end_y) < CELL * 0.8:
            self.finished = True
            self.alive    = False

        # Timeout
        if self.steps >= MAX_STEPS:
            self.alive = False

    def compute_score(self, n_checkpoints):
        s  = len(self.checkpoints_reached) * 120
        s += self.dist_traveled * 0.08
        s -= self.collisions * 25
        if self.finished:
            s += 2000 + max(0, (MAX_STEPS - self.steps) * 2)
        return s

    def draw(self, surface, show_rays=False):
        if show_rays and self.alive:
            spread    = 160
            start_ang = self.angle - spread / 2
            for i, dist in enumerate(self.ray_distances):
                ray_ang = math.radians(start_ang + i * spread / (N_RAYS - 1))
                end_x   = self.x + math.cos(ray_ang) * dist * RAY_LEN
                end_y   = self.y + math.sin(ray_ang) * dist * RAY_LEN
                alpha_s = pygame.Surface((MAZE_W, MAZE_H), pygame.SRCALPHA)
                pygame.draw.line(alpha_s, (*ACCENT[:3], 40),
                                 (int(self.x), int(self.y)),
                                 (int(end_x), int(end_y)), 1)
                surface.blit(alpha_s, (0, 0))

        if not self.alive and not self.finished:
            col = CAR_DEAD
        elif self.is_elite:
            col = CAR_ELITE
        else:
            col = CAR_ALIVE

        # Dessiner la voiture comme un triangle orienté
        rad = math.radians(self.angle)
        tip  = (self.x + math.cos(rad)*CAR_SIZE*1.6,
                self.y + math.sin(rad)*CAR_SIZE*1.6)
        lft  = (self.x + math.cos(rad+2.4)*CAR_SIZE,
                self.y + math.sin(rad+2.4)*CAR_SIZE)
        rgt  = (self.x + math.cos(rad-2.4)*CAR_SIZE,
                self.y + math.sin(rad-2.4)*CAR_SIZE)
        pygame.draw.polygon(surface, col, [tip, lft, rgt])

# ─────────────────────────────────────────────
#  ALGORITHME GÉNÉTIQUE
# ─────────────────────────────────────────────
class Evolution:
    def __init__(self):
        self.generation    = 0
        self.sigma         = SIGMA_INIT
        self.best_scores   = []
        self.avg_scores    = []
        self.stagnation    = 0
        self.genome_size   = NeuralNet(NN_LAYERS).genome_size()
        self.population    = [Car() for _ in range(POP_SIZE)]

    def evaluate(self, n_checkpoints):
        for car in self.population:
            car.score = car.compute_score(n_checkpoints)

    def _tournament(self, ranked):
        candidates = random.sample(ranked, min(K_TOURNAMENT, len(ranked)))
        return max(candidates, key=lambda c: c.score)

    def _mutate(self, genome):
        g = genome.copy()
        for i in range(len(g)):
            if random.random() < MUTATION_RATE:
                g[i] += np.random.normal(0, self.sigma)
        return g

    def _crossover(self, g1, g2):
        mask = np.random.rand(len(g1)) < 0.5
        return np.where(mask, g1, g2)

    def next_generation(self):
        ranked = sorted(self.population, key=lambda c: c.score, reverse=True)

        best = ranked[0].score
        avg  = np.mean([c.score for c in self.population])
        self.best_scores.append(best)
        self.avg_scores.append(avg)

        # Sigma adaptatif
        if len(self.best_scores) >= 5:
            improvement = self.best_scores[-1] - self.best_scores[-5]
            if improvement < 15:
                self.stagnation += 1
                self.sigma = min(0.35, self.sigma * 1.4)
            else:
                self.stagnation = 0
                self.sigma = max(0.02, self.sigma * 0.92)

        # Injection de diversité si longue stagnation
        inject = []
        if self.stagnation > 8:
            n_inject = POP_SIZE // 5
            inject   = [Car() for _ in range(n_inject)]
            self.stagnation = 0

        # Nouvelle génération
        new_pop = []

        # Élites intactes
        for i in range(N_ELITE):
            c = Car(genome=ranked[i].brain.get_genome())
            c.is_elite = True
            new_pop.append(c)

        # Remplir le reste
        while len(new_pop) + len(inject) < POP_SIZE:
            if random.random() < 0.3 and len(ranked) >= 2:
                # Crossover entre 2 parents
                p1 = self._tournament(ranked)
                p2 = self._tournament(ranked)
                child_g = self._crossover(p1.brain.get_genome(),
                                          p2.brain.get_genome())
                child_g = self._mutate(child_g)
            else:
                parent  = self._tournament(ranked)
                child_g = self._mutate(parent.brain.get_genome())
            new_pop.append(Car(genome=child_g))

        self.population = new_pop + inject
        self.generation += 1

    def best_car(self):
        return max(self.population, key=lambda c: c.score)

# ─────────────────────────────────────────────
#  PANNEAU D'INFORMATION (droite)
# ─────────────────────────────────────────────
def draw_panel(surface, evo, maze, offset_x):
    panel = pygame.Surface((PANEL_W, SCREEN_H))
    panel.fill(PANEL_BG)

    f_big  = pygame.font.SysFont("monospace", 18, bold=True)
    f_med  = pygame.font.SysFont("monospace", 14)
    f_sml  = pygame.font.SysFont("monospace", 12)

    def txt(text, x, y, font=f_med, color=TEXT_COL):
        surface.blit(font.render(text, True, color), (offset_x + x, y))

    # Titre
    txt("▶ MAZE EVOLUTION", 10, 12, f_big, ACCENT)
    pygame.draw.line(surface, (50,55,75), (offset_x, 38), (SCREEN_W, 38), 1)

    # Stats génération
    txt(f"Génération   : {evo.generation}", 10, 50)
    txt(f"Population   : {POP_SIZE}", 10, 70)
    txt(f"Sigma (σ)    : {evo.sigma:.4f}", 10, 90)
    txt(f"Stagnation   : {evo.stagnation}", 10, 110,
        color=(255, 120, 80) if evo.stagnation > 5 else TEXT_COL)

    pygame.draw.line(surface, (50,55,75), (offset_x, 135), (SCREEN_W, 135), 1)

    # Meilleur score
    alive = [c for c in evo.population if c.alive]
    dead  = [c for c in evo.population if not c.alive]
    finished = [c for c in evo.population if c.finished]

    if evo.best_scores:
        txt(f"Meilleur ever: {max(evo.best_scores):.0f}", 10, 145, color=CAR_ELITE)
    if dead:
        best_dead = max(dead, key=lambda c: c.score)
        txt(f"Meilleur run : {best_dead.score:.0f}", 10, 165, color=CAR_ALIVE)

    txt(f"En vie      : {len(alive)}", 10, 190, color=START_COL)
    txt(f"Éliminées   : {len(dead)}", 10, 210, color=CAR_DEAD)
    txt(f"Arrivées    : {len(finished)}", 10, 230, color=CHKP_DONE)

    pygame.draw.line(surface, (50,55,75), (offset_x, 255), (SCREEN_W, 255), 1)

    # Checkpoints info
    if alive:
        best_alive = max(alive, key=lambda c: len(c.checkpoints_reached))
        cp_count   = len(maze.checkpoints)
        txt("Checkpoints :", 10, 265, color=TEXT_DIM)
        bar_w = PANEL_W - 30
        reached = len(best_alive.checkpoints_reached)
        pygame.draw.rect(surface, (50,55,75), (offset_x+10, 285, bar_w, 10), border_radius=3)
        fill_w = int(bar_w * reached / max(1, cp_count))
        pygame.draw.rect(surface, CHKP_COL, (offset_x+10, 285, fill_w, 10), border_radius=3)
        txt(f"  {reached}/{cp_count}", 10, 300, f_sml, CHKP_COL)

    pygame.draw.line(surface, (50,55,75), (offset_x, 320), (SCREEN_W, 320), 1)

    # Historique graphique
    txt("Historique scores :", 10, 328, color=TEXT_DIM)
    if len(evo.best_scores) > 1:
        gh = 140
        gw = PANEL_W - 30
        gx = offset_x + 10
        gy = 350
        pygame.draw.rect(surface, (30,32,45), (gx, gy, gw, gh), border_radius=4)

        max_s = max(evo.best_scores) or 1
        n     = len(evo.best_scores)

        # Courbe best
        pts_best = []
        pts_avg  = []
        for i, (bs, av) in enumerate(zip(evo.best_scores, evo.avg_scores)):
            px = gx + int(i / max(1, n-1) * gw)
            py_b = gy + gh - int(bs / max_s * gh)
            py_a = gy + gh - int(av / max_s * gh)
            pts_best.append((px, py_b))
            pts_avg.append((px, py_a))

        if len(pts_best) >= 2:
            pygame.draw.lines(surface, (100, 150, 255), False, pts_best, 2)
            pygame.draw.lines(surface, (200, 200, 80),  False, pts_avg,  1)

        txt("— best  — avg", gx - offset_x, gy + gh + 5, f_sml, TEXT_DIM)

    pygame.draw.line(surface, (50,55,75), (offset_x, 510), (SCREEN_W, 510), 1)

    # Légende
    txt("LÉGENDE", 10, 518, f_sml, TEXT_DIM)
    pygame.draw.circle(surface, CAR_ELITE, (offset_x+20, 540), 6)
    txt("Élite (non mutée)", 32, 535, f_sml)
    pygame.draw.circle(surface, CAR_ALIVE, (offset_x+20, 560), 6)
    txt("Voiture vivante",   32, 555, f_sml)
    pygame.draw.circle(surface, CAR_DEAD,  (offset_x+20, 580), 6)
    txt("Voiture éliminée",  32, 575, f_sml)
    pygame.draw.circle(surface, CHKP_COL,  (offset_x+20, 600), 6)
    txt("Checkpoint",        32, 595, f_sml)

    # Touches
    txt("[ESPACE] pause   [R] reset   [Q] quitter", 10, 630, f_sml, TEXT_DIM)
    txt("[V] rayons on/off", 10, 648, f_sml, TEXT_DIM)

    surface.blit(panel, (0, 0))  # déjà dessiné directement sur surface

# ─────────────────────────────────────────────
#  BOUCLE PRINCIPALE
# ─────────────────────────────────────────────
def main():
    pygame.init()
    screen = pygame.display.set_mode((SCREEN_W, SCREEN_H))
    pygame.display.set_caption("Maze Car Evolution")
    clock  = pygame.time.Clock()

    # Surfaces
    maze_surf  = pygame.Surface((MAZE_W, MAZE_H))

    # Init labyrinthe et évolution
    random.seed(42)
    np.random.seed(42)
    maze = Maze()
    evo  = Evolution()

    # Cast rayons une première fois (vide)
    for car in evo.population:
        car.cast_rays(maze_surf, maze.walls)

    paused    = False
    show_rays = False
    running   = True
    step      = 0

    while running:
        # ── Événements ────────────────────────────
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                running = False
            if event.type == pygame.KEYDOWN:
                if event.key == pygame.K_q:
                    running = False
                if event.key == pygame.K_SPACE:
                    paused = not paused
                if event.key == pygame.K_r:
                    maze = Maze()
                    evo  = Evolution()
                    for car in evo.population:
                        car.cast_rays(maze_surf, maze.walls)
                    step = 0
                if event.key == pygame.K_v:
                    show_rays = not show_rays

        if not paused:
            # ── Mise à jour voitures ──────────────────
            all_done = True
            for car in evo.population:
                if car.alive:
                    all_done = False
                    car.cast_rays(maze_surf, maze.walls)
                    car.update(maze.walls, maze.checkpoints)

            step += 1

            # ── Fin de génération ─────────────────────
            if all_done or step >= MAX_STEPS:
                evo.evaluate(len(maze.checkpoints))
                evo.next_generation()
                for car in evo.population:
                    car.cast_rays(maze_surf, maze.walls)
                step = 0

        # ── Rendu ─────────────────────────────────
        screen.fill(BG)

        # Labyrinthe
        maze.draw(maze_surf)
        maze.draw_checkpoints(maze_surf,
            max(evo.population, key=lambda c: len(c.checkpoints_reached)).checkpoints_reached
            if evo.population else set())

        # Voitures (les mortes d'abord, puis les vivantes)
        dead_cars  = [c for c in evo.population if not c.alive]
        alive_cars = [c for c in evo.population if c.alive]
        for car in dead_cars:
            car.draw(maze_surf, show_rays=False)
        for car in alive_cars:
            car.draw(maze_surf, show_rays=show_rays)

        screen.blit(maze_surf, (0, 0))

        # Séparateur
        pygame.draw.line(screen, (50, 55, 75), (MAZE_W, 0), (MAZE_W, SCREEN_H), 1)

        # Panneau info
        draw_panel(screen, evo, maze, offset_x=MAZE_W)

        # Pause overlay
        if paused:
            s = pygame.Surface((MAZE_W, SCREEN_H), pygame.SRCALPHA)
            s.fill((0, 0, 0, 100))
            screen.blit(s, (0, 0))
            f = pygame.font.SysFont("monospace", 32, bold=True)
            lbl = f.render("⏸  PAUSE", True, (255, 255, 255))
            screen.blit(lbl, (MAZE_W//2 - lbl.get_width()//2, SCREEN_H//2))

        pygame.display.flip()
        clock.tick(FPS)

    pygame.quit()
    sys.exit()

if __name__ == "__main__":
    main()
