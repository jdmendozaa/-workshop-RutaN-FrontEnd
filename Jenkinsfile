pipeline {
    agent any

    stages {
        stage('Build Frontend Image') {
            steps {
                sh 'docker build -t danieldi/front .'
            }
        }
        stage('Test front') {
            steps {
                echo 'testing...'
            }
        }
        stage('Docker Login') {
            steps {
                sh 'docker login -u danieldi -p Praxis20221*team7'
            }
        }
        stage('Push image') {
            steps {
                sh 'docker push danieldi/front'
            }
    }
}