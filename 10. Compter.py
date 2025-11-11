import subprocess

utilisateur = "amin"  # Remplace par le nom d'utilisateur souhaité
commande = ["ps", "-u", utilisateur]
resultat = subprocess.check_output(commande).decode()
nb_processus = len(resultat.strip().split("\n")) - 1  # Enlever l'en-tête

print(f"Nombre de processus pour l'utilisateur {utilisateur} : {nb_processus}")
