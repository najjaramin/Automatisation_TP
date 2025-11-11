import subprocess

cron_job = "0 2 * * * /path/to/script.sh"
subprocess.run(f'(crontab -l; echo "{cron_job}") | crontab -', shell=True)