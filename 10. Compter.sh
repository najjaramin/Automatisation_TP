read -p "Entrez le nom de l'utilisateur : " user
count=$(ps -u "$user" --no-headers | wc -l)
echo "Nombre de processus pour l'utilisateur $user : $count"