import webview
import os

def main():
    # Configurez le proxy
    proxy_address = '45.77.56.114'  # Remplacez par l'adresse du proxy public
    proxy_port = '80'  # Remplacez par le port du proxy public
    os.environ['HTTP_PROXY'] = f'http://{proxy_address}:{proxy_port}'
    os.environ['HTTPS_PROXY'] = f'http://{proxy_address}:{proxy_port}'

    # Créez une fenêtre de navigateur avec le backend disponible
    try:
        webview.create_window('Mon navigateur avec proxy', 'https://www.example.com')
        webview.start()
    except Exception as e:
        print(f"Erreur : {e}")

if __name__ == '__main__':
    main()
