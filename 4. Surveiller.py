import shutil

total, used, free = shutil.disk_usage("/")
percent_free = free / total * 100

if percent_free < 10:
    print("⚠️ Alerte : espace disque libre inférieur à 10%")
else:
    print(f"Espace libre : {percent_free:.2f}%")
