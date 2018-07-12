echo "Hello world"
Start-Process docker -ArgumentList "run", "-d", "micdenny/redis-windows" -NoNewWindow
Start-Sleep -s 3600