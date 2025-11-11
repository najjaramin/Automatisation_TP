import subprocess

service = "nginx"
subprocess.run(["sudo", "systemctl", "restart", service])