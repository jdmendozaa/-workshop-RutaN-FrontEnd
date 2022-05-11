pipeline {
    agent any

    stages {
        stage('Build Frontend') {
            steps {
                sh "docker build -t front ."
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}