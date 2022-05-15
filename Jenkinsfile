node {
    stage('Get a changes'){
        git url:'https://github.com/jsapuyesp/praxis-FE', branch:'main'
    }
    stage('Build Frontend Image') {
        sh 'docker build -t danieldi/front .'
    }
    stage('Test front') {
        echo 'testing...'
    }
    stage('Docker Login') {
        sh 'docker login -u danieldi -p Praxis20221*team7'
    }
    stage('Push image') {
        sh 'docker push danieldi/front'
    }
}