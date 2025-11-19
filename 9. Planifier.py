import os
task_name = "TestScript"
script_path = r"C:\Users\Public\TP_Automatisation\1.supprime.py"
os.system(f'schtasks /Create /SC DAILY /TN {task_name} /TR "python {script_path}" /ST 02:00 /F')
print("Tâche planifiée à 2h du matin")
