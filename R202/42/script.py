import os
import pyzipper
from concurrent.futures import ThreadPoolExecutor, as_completed
import threading

SOURCE_ZIP = "42.zip"
DEST_DIR = "extrait"
PASSWORD = b"42"
MAX_WORKERS = 8

total_extracted = 0
lock = threading.Lock()

def format_size(b):
    for unit in ["o", "Ko", "Mo", "Go", "To"]:
        if b < 1024:
            return f"{b:.1f} {unit}"
        b /= 1024
    return f"{b:.1f} Po"

def extract_zip(zip_path, extract_to):
    global total_extracted
    extracted_here = 0
    files_count = 0
    try:
        with pyzipper.AESZipFile(zip_path) as z:
            z.pwd = PASSWORD
            for info in z.infolist():
                try:
                    z.extract(info, extract_to)
                    with lock:
                        total_extracted += info.file_size
                    extracted_here += info.file_size
                    files_count += 1
                except Exception as e:
                    print(f"[!] {info.filename} : {e}")
        print(f"  ✓ {os.path.basename(zip_path)} → {files_count} fichier(s), {format_size(extracted_here)}")
        return True
    except Exception as e:
        print(f"[!] Archive {zip_path} : {e}")
        return False

def find_zip_files(directory):
    return [
        os.path.join(root, f)
        for root, _, files in os.walk(directory)
        for f in files
        if f.lower().endswith(".zip")
    ]

def main():
    os.makedirs(DEST_DIR, exist_ok=True)
    loop = 0

    while True:
        # Recopie le zip source si plus rien à extraire
        zips = find_zip_files(DEST_DIR)
        if not zips:
            import shutil
            print(f"\n[→] Plus de zips, on repart depuis {SOURCE_ZIP}...")
            shutil.copy(SOURCE_ZIP, os.path.join(DEST_DIR, os.path.basename(SOURCE_ZIP)))
            zips = find_zip_files(DEST_DIR)
            if not zips:
                print("Source introuvable, arrêt.")
                break

        loop += 1
        print(f"\n[Loop {loop}] {len(zips)} zip(s) — Cumul : {format_size(total_extracted)}")
        loop_start = total_extracted

        with ThreadPoolExecutor(max_workers=MAX_WORKERS) as executor:
            futures = {executor.submit(extract_zip, z, DEST_DIR): z for z in zips}
            for future in as_completed(futures):
                z = futures[future]
                try:
                    os.remove(z)
                except Exception as e:
                    print(f"[!] Suppression {z} : {e}")

        loop_delta = total_extracted - loop_start
        print(f"  → Ce loop : +{format_size(loop_delta)} | Cumul : {format_size(total_extracted)}")

if __name__ == "__main__":
    main()