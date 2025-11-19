import shutil
import os

source_dir = r"C:\Users\Public\TP_Automatisation\test_zip"
os.makedirs(source_dir, exist_ok=True)

# créer des fichiers de test
for i in range(3):
    with open(os.path.join(source_dir, f"file{i}.txt"), "w") as f:
        f.write("test\n")

backup_file = r"C:\Users\Public\TP_Automatisation\backup"
shutil.make_archive(backup_file, "zip", source_dir)
print(f"Archive créée : {backup_file}.zip")
