import os

folder = r"C:\Users\Public\TP_Automatisation\test_logs"
files = [(f, os.path.getsize(os.path.join(folder, f))) for f in os.listdir(folder) if os.path.isfile(os.path.join(folder, f))]
files.sort(key=lambda x: x[1], reverse=True) 
for f, size in files[:5]:
    print(f"{f} : {size} octets")
