import os

files = [(f, os.path.getsize(f)) for f in os.listdir(".") if os.path.isfile(f)]
files.sort(key=lambda x: x[1], reverse=True)

for f, size in files[:5]:
    print(f"{f}: {size / (1024*1024):.2f} MB")