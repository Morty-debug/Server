
### usando docker
```bash
docker compose up
```

### usando solo golang
```bash
mkdir subidos
go run main.go
```

### mostrar en internet mediante un tunel (Linux)
```bash
sudo chmod a+x ngrok.bin
./ngrok.bin config add-authtoken 21sjOuIGRhjIdgEDjQ7yp0HEBgt_2d34Sft292xQybQfQKEeq
./ngrok.bin http 80
```

### mostrar en internet mediante un tunel (Windows)
```bash
ngrok.exe config add-authtoken 21sjOuIGRhjIdgEDjQ7yp0HEBgt_2d34Sft292xQybQfQKEeq
ngrok.exe http 80
```
