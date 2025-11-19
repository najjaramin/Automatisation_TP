import shutil

total, used, free = shutil.disk_usage("C:/")
percent_free = free / total * 100

if percent_free < 10:
    print(f"Alerte ! Espace disque faible : {percent_free:.2f}% libre")
else:
    print(f"Espace disque OK : {percent_free:.2f}% libre")
