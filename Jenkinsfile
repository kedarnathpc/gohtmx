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
        
        stage('Clone-code') {
            steps {
                echo '<---------Cloning code--------->'
                git branch: 'main', url: 'https://github.com/kedarnathpc/gohtmx.git'
                echo '<---------Code cloned--------->'
            }
        } 
        
        stage('Install Dependencies') {
            steps {
                echo '<---------Installing dependencies--------->'
                sh 'go mod download'
                echo '<---------Dependencies installed--------->'
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
