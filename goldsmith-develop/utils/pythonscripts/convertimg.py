from PIL import Image
import os

# chemin dossier
input_folder = "/home/noam/goldsmith/utils/pythonscripts/img"
output_folder = os.path.join(input_folder, "webp_output")

# crée dossier sortie s'il est pas la
os.makedirs(output_folder, exist_ok=True)

#extension dispor
image_extensions = (".jpg", ".jpeg", ".png", ".bmp", ".tiff", ".gif")

# récup et converti 
for filename in os.listdir(input_folder):
    if filename.lower().endswith(image_extensions):
        file_path = os.path.join(input_folder, filename)
        with Image.open(file_path) as img:
            # conversion
            output_path = os.path.join(output_folder, os.path.splitext(filename)[0] + ".webp")
            img.save(output_path, "WEBP")
        print(f"Converti : {filename} -> {os.path.basename(output_path)}")

print("Conversion terminée !")
