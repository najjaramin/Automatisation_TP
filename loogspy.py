import os
import time

log_dir = "/var/log/myapp"
now = time.time()

for filename in os.listdir(log_dir):
    filepath = os.path.join(log_dir, filename)
    if os.path.isfile(filepath) and now - os.path.getmtime(filepath) > 7 * 86400:
        os.remove(filepath)