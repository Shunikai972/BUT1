import os
import re
import time
import requests
from bs4 import BeautifulSoup

def main():
    url = "https://pokemondb.net/sprites"
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
    }

    print("Téléchargement de la page principale...")
    response = requests.get(url, headers=headers)
    response.raise_for_status()
    soup = BeautifulSoup(response.text, "html.parser")

    # Trouver toutes les sections de génération
    # Les ancres sont #gen1, #gen2, etc. avec des <h2> comme titres
    gen_sections = soup.find_all("h2")

    current_gen = 0
    pokemon_list = []  # liste de tuples (generation, nom, url_image)

    # Parcourir tous les éléments pour associer chaque pokémon à sa génération
    # Stratégie : trouver les h2 "Generation X" puis les liens qui suivent
    # On va plutôt parcourir les panels/divs

    # Approche : itérer sur tous les éléments frères après chaque h2
    content = soup.find("main") or soup.find("div", class_="main-content") or soup

    # Trouver tous les h2 et les infocard-list qui suivent
    all_h2 = content.find_all("h2")

    for h2 in all_h2:
        # Extraire le numéro de génération
        title = h2.get_text(strip=True)
        match = re.search(r"Generation\s+(\d+)", title)
        if not match:
            continue
        gen_num = int(match.group(1))
        gen_folder = f"generation{gen_num}"
        os.makedirs(gen_folder, exist_ok=True)

        print(f"\n=== {title} ===")

        # Trouver le conteneur de pokémons qui suit ce h2
        sibling = h2.find_next_sibling()
        while sibling:
            # Si on tombe sur un autre h2 ou hr, on arrête
            if sibling.name == "h2":
                break

            # Chercher tous les liens de pokémons dans ce bloc
            pokemon_links = sibling.find_all("a", href=re.compile(r"^/sprites/"))
            for link in pokemon_links:
                # Extraire le nom depuis le href: /sprites/bulbasaur -> bulbasaur
                href = link.get("href", "")
                poke_name = href.replace("/sprites/", "").strip("/")
                if not poke_name:
                    continue

                # Trouver l'image dans ce lien
                img = link.find("img")
                if img:
                    img_url = img.get("src", "")
                else:
                    continue

                if not img_url:
                    continue

                # Le nom du fichier
                # Nettoyer le nom (certains ont des caractères spéciaux comme nidoran-f)
                safe_name = poke_name.lower().replace(" ", "-")
                filename = os.path.join(gen_folder, f"{safe_name}.jpg")

                if os.path.exists(filename):
                    print(f"  [SKIP] {safe_name} (déjà téléchargé)")
                    continue

                print(f"  Téléchargement de {safe_name}...")
                try:
                    img_response = requests.get(img_url, headers=headers)
                    img_response.raise_for_status()
                    with open(filename, "wb") as f:
                        f.write(img_response.content)
                    # Petit délai pour ne pas surcharger le serveur
                    time.sleep(0.1)
                except Exception as e:
                    print(f"  [ERREUR] {safe_name}: {e}")

            sibling = sibling.find_next_sibling()

    print("\n✅ Téléchargement terminé !")


if __name__ == "__main__":
    main()
