docker build --no-cache -t ms-game-engine:latest .;

docker tag ms-game-engine:latest emailtovamos/ms-game-engine:latest;

docker push emailtovamos/ms-game-engine:latest;

kubectl delete deployment ms-game-engine;

kubectl apply -f devops/deployment.yaml;

#chmod +x command.sh is the command to make a script file executable in Mac.