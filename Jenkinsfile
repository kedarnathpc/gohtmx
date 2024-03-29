pipeline {
    agent {
        node {
            label 'slave'
        }
    }
    
    stages {
        stage('Checkout') {
            steps {
                echo '<---------Checking out code--------->'
                checkout scm
                echo '<---------Code checked out--------->'
            }
        }
        
        stage('Build') {
            steps {
                echo '<---------Building code--------->'
                sh 'go build .'
                echo '<---------Code built--------->'
            }
        }
        
        stage('Test') {
            steps {
                echo '<---------Testing code--------->'
                sh 'go test ./...'
                echo '<---------Code tested--------->'
            }
        }
    }
}
