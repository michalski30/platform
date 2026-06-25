buildcdev:
	cd backend && docker build -t platform_backend_dev:0.0.1 -f ./Dockerfile.dev .
	cd frontend && docker build -t platform_frontend_dev:0.0.1 -f ./Dockerfile.dev .
	
runcdev:
	docker run -d -p 8080:8080 --name platform_backend_dev --env-file .env platform_backend_dev:0.0.1
	docker run -d -p 5173:5173 --name platform_frontend_dev --env-file .env platform_frontend_dev:0.0.1

buildcprod:
	cd frontend && docker build -t michalski30/platform_frontend_prod:0.0.1 -f ./Dockerfile.prod .
	cd backend && docker build -t michalski30/platform_backend_prod:0.0.1 -f ./Dockerfile.prod .

runwatch:
	docker compose -f docker-compose.yaml -f docker-compose.dev.yaml watch
