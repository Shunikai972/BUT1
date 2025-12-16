from fpdf import FPDF

class PDF(FPDF):
    def header(self):
        self.set_font('Arial', 'B', 14)
        self.cell(0, 10, 'Fiche de Revision COMPLETE R1.04 (TP1-4)', 0, 1, 'C')
        self.ln(5)

    def footer(self):
        self.set_y(-15)
        self.set_font('Arial', 'I', 8)
        self.cell(0, 10, f'Page {self.page_no()}', 0, 0, 'C')

    def section_title(self, title):
        self.set_font('Arial', 'B', 12)
        self.set_fill_color(230, 230, 230)
        self.cell(0, 8, title, 0, 1, 'L', 1)
        self.ln(2)

    def subsection(self, title):
        self.set_font('Arial', 'B', 10)
        self.cell(0, 6, title, 0, 1, 'L')

    def command(self, cmd, desc):
        self.set_font('Courier', 'B', 9)
        self.cell(70, 5, cmd, 0, 0)
        self.set_font('Arial', '', 9)
        self.multi_cell(0, 5, desc)
        self.ln(1)

pdf = FPDF()
pdf.add_page()
pdf.set_auto_page_break(auto=True, margin=15)

# --- TP1 ---
pdf.section_title('TP 1 : Shell & Système de Fichiers')

pdf.subsection('Navigation & Aide')
pdf.command('cd -', 'Retour au répertoire précédent')
pdf.command('cd', 'Retour au répertoire personnel (home)')
pdf.command('cd ~ / cd ~user', 'Aller au home courant / home d\'un autre user')
pdf.command('pwd', 'Afficher chemin absolu courant')
pdf.command('type [cmd]', 'Connaitre le type d\'une commande (alias, interne...)')
pdf.command('date --help', 'Aide interne (Epoch = 01/01/1970)')
pdf.command('date "+%d %B %Y"', 'Date formatée (ex: 25 décembre 2022)')

pdf.subsection('Manipulation de Fichiers')
pdf.command('mkdir -p A/B/C', 'Créer une arborescence complète (parents inclus)')
pdf.command('touch [file]', 'Créer fichier vide ou mettre à jour l\'horodatage')
pdf.command('cp [src] [dst]', 'Copier fichier')
pdf.command('mv [src] [dst]', 'Déplacer ou Renommer')
pdf.command('rm -rf [dir]', 'Supprimer récursivement et sans confirmation')
pdf.command('cat -n [file]', 'Afficher contenu avec numéros de ligne')
pdf.command('head -2 [file]', 'Afficher les 2 premières lignes')
pdf.command('tail -2 [file]', 'Afficher les 2 dernières lignes')
pdf.command('tail +2 [file]', 'Afficher tout SAUF la 1ère ligne (commence à la 2e)')
pdf.command('less [file]', 'Visualiseur (sortir avec q). Recherche: /motif ou &motif')

pdf.subsection('Droits (chmod/chown/chgrp)')
pdf.command('chmod u+rx [file]', 'Ajouter Lecture/Exéc au propriétaire (User)')
pdf.command('chmod a+r [file]', 'Lecture pour tous (All)')
pdf.command('chmod g+rw [file]', 'Lecture/Écriture pour le groupe')
pdf.command('chmod 755 [dir]', 'Octal: rwx(7) pour user, rx(5) pour les autres')
pdf.command('chmod 644 [file]', 'Octal: rw(6) user, r(4) autres')
pdf.command('chgrp [grp] [file]', 'Changer le groupe propriétaire')

pdf.subsection('Archivage (tar) & Compression (gzip)')
pdf.command('gzip [file] / zcat', 'Compresser (.gz) / Lire un fichier compressé')
pdf.command('tar -cvf f.tar [dossier]', 'Créer une archive non compressée')
pdf.command('tar -xvf f.tar', 'Extraire une archive (v = verbeux)')
pdf.command('tar -tvf f.tar', 'Lister le contenu sans extraire')
pdf.command('tar -zcvf f.tgz [d]', 'Créer archive compressée GZIP')
pdf.command('tar -jcvf f.tar.bz2 [d]', 'Créer archive compressée BZIP2')
pdf.command('tar ... --directory=Dir', 'Extraire dans un dossier spécifique')

# --- TP2 ---
pdf.section_title('TP 2 : Shell Avancé (Historique, Variables)')

pdf.subsection('Historique & Édition')
pdf.command('history -w [file]', 'Sauvegarder l\'historique dans un fichier')
pdf.command('history -c', 'Effacer l\'historique courant')
pdf.command('history -r [file]', 'Lire/Restaurer un historique')
pdf.command('!!', 'Relancer la dernière commande')
pdf.command('!72', 'Relancer la commande n°72')
pdf.command('!-2', 'Relancer l\'avant-dernière commande')
pdf.command('Ctrl+R', 'Rechercher interactivement dans l\'historique')

pdf.subsection('Jokers (Globbing)')
pdf.command('*', 'N\'importe quelle chaine')
pdf.command('?', 'Un seul caractère')
pdf.command('[abc] / [!abc]', 'Un caractère parmi liste / ou hors liste')
pdf.command('@(pat1|pat2)', 'Globbing étendu (ex: commence par pat1 OU pat2)')

pdf.subsection('Variables & Environnement')
pdf.command('VAR=val', 'Définir variable locale (pas d\'espaces autour du =)')
pdf.command('export VAR', 'Rendre la variable globale (processus fils)')
pdf.command('echo $VAR', 'Accéder au contenu')
pdf.command('$PWD / $OLDPWD', 'Rép. courant / Rép. précédent')
pdf.command('$PATH', 'Chemins de recherche des exécutables')
pdf.command('$PS1', 'Invite de commande (Prompt). Ex: \\u@\\h \\w $')

pdf.subsection('Protection (Quoting)')
pdf.command('\\', 'Échappe le caractère suivant')
pdf.command('\'...\'', 'Protection FORTE : rien n\'est interprété (ni $ ni *)')
pdf.command('\"...\"', 'Protection FAIBLE : $ et ` ` interprétés, * non')
pdf.command('`cmd` ou $(cmd)', 'Substitution : remplace par le résultat de cmd')

# --- TP3 ---
pdf.section_title('TP 3 : Processus, Flux & Filtres')

pdf.subsection('Processus')
pdf.command('ps -e / ps -A', 'Tous les processus')
pdf.command('ps -l', 'Format long (PID, PPID, UID, PRI, CMD...)')
pdf.command('ps -u', 'Processus d\'un utilisateur')
pdf.command('Ctrl+Z', 'Suspendre (Stop) le processus premier plan')
pdf.command('Ctrl+C', 'Interrompre (Kill) le processus')
pdf.command('cmd &', 'Lancer en arrière-plan')
pdf.command('jobs', 'Lister les tâches (Jobs)')
pdf.command('fg %1 / bg %1', 'Mettre job 1 en avant-plan / arrière-plan')

pdf.subsection('Signaux (kill)')
pdf.command('kill [PID]', 'Envoie signal 15 (SIGTERM) par défaut')
pdf.command('kill -9 [PID]', 'SIGKILL (Arrêt forcé immédiat)')
pdf.command('kill -l', 'Lister les signaux (1=HUP, 2=INT, 9=KILL, 15=TERM...)')

pdf.subsection('Redirections E/S')
pdf.command('> fichier', 'Sortie standard vers fichier (Écrase)')
pdf.command('>> fichier', 'Sortie standard vers fichier (Ajoute)')
pdf.command('2> fichier', 'Sortie Erreur vers fichier')
pdf.command('2>> fichier', 'Sortie Erreur ajoutée au fichier')
pdf.command('>& fichier', 'Sortie Standard ET Erreur dans le même fichier')
pdf.command('< fichier', 'Entrée standard depuis un fichier')
pdf.command('echo $?', 'Code retour dernière commande (0=OK, >0=Erreur)')

pdf.subsection('Filtres & Pipes (|)')
pdf.command('wc -l', 'Compter les lignes')
pdf.command('sort -k3,3', 'Trier sur colonne 3 seulement')
pdf.command('sort -n', 'Tri numérique')
pdf.command('sort -t ","', 'Définir le séparateur (ici virgule)')
pdf.command('cut -d "," -f 2-3', 'Garder champs 2 à 3 (séparateur virgule)')
pdf.command('uniq -c', 'Compter les doublons consécutifs (trier avant !)')
pdf.command('join -t: -1 4 -2 3', 'Fusionner fichiers sur champ 4 du f1 et 3 du f2')

# --- TP4 ---
pdf.section_title('TP 4 : Encodage & Recherche')

pdf.subsection('Encodage')
pdf.command('xxd -p', 'Dump hexadécimal brut')
pdf.command('xxd -r -p', 'Reconvertir hexa vers binaire/texte')
pdf.command('iconv -f S -t D', 'Convertir encodage (ex: -f UTF-8 -t ISO-8859-1)')
pdf.command('ISO-8859-1', 'Encodage 1 octet (Latin-1)')
pdf.command('UTF-8', 'Encodage taille variable (compatible ASCII)')
pdf.command('UCS-2 / UCS-4', 'Encodages taille fixe (2 ou 4 octets)')

pdf.subsection('Recherche (find)')
pdf.command('find . -name "*.mp3"', 'Chercher par nom')
pdf.command('find . -type f / -type d', 'Chercher fichiers / dossiers')
pdf.command('find . -size +1k', 'Taille supérieure à 1ko')
pdf.command('find . -not -name "*t"', 'Négation')
pdf.command('find . -exec cmd {} \\;', 'Exécuter commande sur chaque résultat ({})')
pdf.command('-printf "%TF %f"', 'Afficher date et nom formatés')

pdf.subsection('Recherche Texte (grep)')
pdf.command('grep "motif" [f]', 'Lignes contenant le motif')
pdf.command('grep -v "motif"', 'Lignes NE contenant PAS le motif')
pdf.command('grep -n', 'Afficher numéros de ligne')
pdf.command('grep -i', 'Ignorer la casse (majuscule/minuscule)')
pdf.command('grep -e A -e B', 'Contenant A OU B')
raise('')

pdf.output('Fiche_Revision_R104_Complete.pdf')