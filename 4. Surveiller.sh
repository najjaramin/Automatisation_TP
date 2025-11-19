#!/bin/bash
FREE=$(df -h /c | tail -1 | awk '{print $5}' | sed 's/%//')
if [ "$FREE" -lt 10 ]; then
  echo "Alerte ! Espace disque faible : $FREE%"
else
  echo "Espace disque OK : $FREE%"
fi
