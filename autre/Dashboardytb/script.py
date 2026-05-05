from mutagen import File

audio = File("music/JACK UZI - JACK DE LA TOURETTE (Clip oficiel).mp3")

# Métadonnées basiques
print("Artiste :", audio.get("TPE1"))
print("Album :", audio.get("TALB"))
print("Titre :", audio.get("TIT2"))
