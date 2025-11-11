import subprocess

output = subprocess.check_output(["who"]).decode()
print("Utilisateurs connectés :\n", output)