#!/bin/bash
powershell.exe -Command "schtasks /Create /SC DAILY /TN 'TestScript' /TR 'python C:\\Users\\Public\\TP_Automatisation\\1.supprime.py' /ST 02:00 /F"
