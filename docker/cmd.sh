#build
# GradingService
docker build -f docker/gradingservice.Dockerfile -t gradingservice:latest .

# LogService
docker build -f docker/logservice.Dockerfile -t logservice:latest .

# RegistryService
docker build -f docker/registryservice.Dockerfile -t registryservice:latest .


# run
# RegistryService
docker run -d -p 3000:3000 --name registryservice registryservice:latest

# GradingService
docker run -d -p 6000:6000 --name gradingservice gradingservice:latest

# LogService
docker run -d -p 4000:4000 --name logservice logservice:latest