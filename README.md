
### usando docker
```bash
docker compose up
```

### usando solo golang
```bash
mkdir subidos
go run main.go
```


### mostrar en internet mediante un tunel (Windows)
```batch
ngrok.exe config add-authtoken 21sjOuIGRhjIdgEDjQ7yp0HEBgt_2d34Sft292xQybQfQKEeq
ngrok.exe http 80
```

### mostrar en internet mediante un tunel (Linux)
```bash
sudo chmod a+x ngrok.bin
./ngrok.bin config add-authtoken 21sjOuIGRhjIdgEDjQ7yp0HEBgt_2d34Sft292xQybQfQKEeq
./ngrok.bin http 80
```

### mostrar en internet mediante un tunel (BSD)
```bash
sudo chmod a+x ngrok.bsd
./ngrok.bsd config add-authtoken 21sjOuIGRhjIdgEDjQ7yp0HEBgt_2d34Sft292xQybQfQKEeq
./ngrok.bsd http 80
```

### mostrar en internet mediante un tunel (MacOS)
```bash
sudo chmod a+x ngrok.mac
./ngrok.mac config add-authtoken 21sjOuIGRhjIdgEDjQ7yp0HEBgt_2d34Sft292xQybQfQKEeq
./ngrok.mac http 80
```
