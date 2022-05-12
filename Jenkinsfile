pipeline {
    agent any

    stages {
        stage('Build Frontend') {
            steps {
                echo 'docker pull danieldi/front'
            }
        }
        stage('Test front') {
            steps {
                echo 'testing...'
            }
        }
        stage('Deploy Front') {
            steps {
                echo 'docker run --name font-end --network="my-net" --ip 122.22.0.32 -p 4200:4200 -d danieldi/front'
            }
        }
    }
}