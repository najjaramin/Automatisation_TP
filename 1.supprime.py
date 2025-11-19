import os
import time

log_dir = r"C:\Users\Public\TP_Automatisation\test_logs"

days = 7
now = time.time()
limit = days * 24 * 60 * 60

os.makedirs(log_dir, exist_ok=True)  

for filename in os.listdir(log_dir):
    file_path = os.path.join(log_dir, filename)
    if os.path.isfile(file_path):
        file_age = now - os.path.getmtime(file_path)
        if file_age > limit:
            print(f"Suppression de : {file_path}")
            os.remove(file_path)
