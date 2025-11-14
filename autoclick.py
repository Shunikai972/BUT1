#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Autoclicker robuste + debug info.
Toggle: touche '*' (essaye aussi le pavé numérique), fallback: F8.
"""

import os
import sys
import threading
import time

# Affiche infos d'environnement utiles pour le debug
print("=== ENV INFO ===")
print("python:", sys.executable, sys.version.splitlines()[0])
print("DISPLAY:", os.environ.get("DISPLAY"))
print("WAYLAND_DISPLAY:", os.environ.get("WAYLAND_DISPLAY"))
print("================\n")

# Dépendances
try:
    from pynput.mouse import Controller, Button
    from pynput import keyboard
except Exception as e:
    print("Erreur import pynput:", e)
    print("Installe pynput pour cet interpréteur :")
    print(f"{sys.executable} -m pip install --user pynput")
    sys.exit(1)

mouse = Controller()
clicking = False
click_delay = 0.05  # secondes entre clics

def click_loop():
    """Boucle de clics"""
    global clicking
    while True:
        if clicking:
            try:
                mouse.click(Button.left)
            except Exception as e:
                print("Erreur clic:", e)
                # si clic échoue, on arrête pour éviter boucle rapide d'erreurs
                clicking = False
            time.sleep(click_delay)
        else:
            time.sleep(0.12)

def is_star_key(key):
    """
    Détecte la touche '*' de plusieurs façons :
    - key.char == '*'
    - key.vk == 106 (pavé numérique sur certaines plateformes)
    - repr(key) contient '*' (fallback)
    """
    try:
        if getattr(key, "char", None) == '*':
            return True
    except Exception:
        pass
    # certains environnements fournissent un attribut vk (ex: Windows/Linux)
    vk = getattr(key, "vk", None)
    if vk is not None:
        if vk in (106,):  # 106 est souvent le VK du pavé numérique '*'
            return True
    # fallback: representation textuelle
    if "*" in repr(key):
        return True
    return False

def on_press(key):
    global clicking
    # Affiche info pour debug à chaque touche (utile pour savoir ce que ton environnement envoie)
    print("Key pressed:", repr(key))
    # Toggle si '*' détecté
    try:
        if is_star_key(key):
            clicking = not clicking
            print(f"[{'ON' if clicking else 'OFF'}] AutoClicker (toggle via *)")
            return
    except Exception as e:
        print("Erreur detection '*':", e)

    # Fallback utile si '*' ne marche pas dans ton environnement : F8 pour toggle
    try:
        if key == keyboard.Key.f8:
            clicking = not clicking
            print(f"[{'ON' if clicking else 'OFF'}] AutoClicker (toggle via F8)")
            return
    except Exception:
        pass

    # Quitter proprement via ESC
    try:
        if key == keyboard.Key.esc:
            print("ESC détecté -> sortie.")
            return False  # stop listener
    except Exception:
        pass

def on_release(key):
    # noop mais utile si on veut plus tard
    pass

# Lancer thread clic
thread = threading.Thread(target=click_loop, daemon=True)
thread.start()

print("Instructions :")
print("- Appuie sur '*' (ou F8) pour activer/désactiver.")
print("- Appuie ESC pour quitter.")
print("- Observe les lignes 'Key pressed: ...' pour voir ce que ton système envoie.")
print("Lancement du listener...\n")

with keyboard.Listener(on_press=on_press, on_release=on_release) as listener:
    listener.join()

print("Programme terminé.")
