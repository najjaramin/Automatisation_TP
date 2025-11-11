free_percent=$(df / | awk 'NR==2 {print 100 - $5}')
if [ "$free_percent" -lt 10 ]; then
    echo "⚠️ Alerte : espace disque libre inférieur à 10%"
else
    echo "Espace libre : $free_percent%"
fi