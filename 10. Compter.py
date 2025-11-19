import os
user = os.getlogin()
count = os.popen(f'tasklist /FI "USERNAME eq {user}"').read().count("\n") - 3  # enlever en-tête
print(f"{count} processus en cours pour {user}")
