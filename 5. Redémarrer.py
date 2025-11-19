import os

service_name = "Spooler"  # exemple : Print Spooler, change selon tes tests

print(f"Redémarrage du service {service_name}...")
os.system(f'net stop "{service_name}"')
os.system(f'net start "{service_name}"')
print("Service redémarré.")
