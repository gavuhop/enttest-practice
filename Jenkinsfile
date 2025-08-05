pipeline {
    agent any
    
    environment {
        DOCKER_IMAGE = 'your-app'
        DOCKER_TAG = "${env.BUILD_NUMBER}"
        DB_HOST = 'your-db-host'
        DB_PORT = '3306'
        DB_NAME = 'dbname'
        DB_USER = 'user'
        DB_PASSWORD = 'password'
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    // Build production Docker image
                    docker.build("${DOCKER_IMAGE}:${DOCKER_TAG}", "-f Dockerfile.prod .")
                    
                    // Tag as latest
                    sh "docker tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_IMAGE}:latest"
                }
            }
        }
        
        stage('Run Tests') {
            steps {
                script {
                    // Run tests in container
                    sh """
                        docker run --rm \
                            -e DB_HOST=${DB_HOST} \
                            -e DB_PORT=${DB_PORT} \
                            -e DB_NAME=${DB_NAME} \
                            -e DB_USER=${DB_USER} \
                            -e DB_PASSWORD=${DB_PASSWORD} \
                            ${DOCKER_IMAGE}:${DOCKER_TAG} \
                            go test ./...
                    """
                }
            }
        }
        
        stage('Database Migration') {
            steps {
                script {
                    // Run Atlas migration
                    sh """
                        docker run --rm \
                            -e ATLAS_DB_URL="mysql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" \
                            ${DOCKER_IMAGE}:${DOCKER_TAG} \
                            atlas migrate apply --env prod
                    """
                }
            }
        }
        
        stage('Deploy to Production') {
            when {
                branch 'main'
            }
            steps {
                script {
                    // Deploy to production
                    sh """
                        # Stop existing container
                        docker stop prod-app || true
                        docker rm prod-app || true
                        
                        # Run new container
                        docker run -d \
                            --name prod-app \
                            --restart unless-stopped \
                            -p 8080:8080 \
                            -e DB_HOST=${DB_HOST} \
                            -e DB_PORT=${DB_PORT} \
                            -e DB_NAME=${DB_NAME} \
                            -e DB_USER=${DB_USER} \
                            -e DB_PASSWORD=${DB_PASSWORD} \
                            ${DOCKER_IMAGE}:${DOCKER_TAG}
                    """
                }
            }
        }
        
        stage('Health Check') {
            steps {
                script {
                    // Wait for app to start
                    sleep 10
                    
                    // Health check
                    sh """
                        curl -f http://localhost:8080/health || exit 1
                    """
                }
            }
        }
        
        stage('Cleanup') {
            always {
                script {
                    // Clean up old images
                    sh """
                        docker image prune -f
                        docker system prune -f
                    """
                }
            }
        }
    }
    
    post {
        success {
            echo 'Deployment successful!'
        }
        failure {
            echo 'Deployment failed!'
            // Rollback if needed
            sh """
                docker stop prod-app || true
                docker rm prod-app || true
                docker run -d \
                    --name prod-app \
                    --restart unless-stopped \
                    -p 8080:8080 \
                    -e DB_HOST=${DB_HOST} \
                    -e DB_PORT=${DB_PORT} \
                    -e DB_NAME=${DB_NAME} \
                    -e DB_USER=${DB_USER} \
                    -e DB_PASSWORD=${DB_PASSWORD} \
                    ${DOCKER_IMAGE}:latest
            """
        }
    }
} 