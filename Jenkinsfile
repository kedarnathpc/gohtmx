def registry = "https://miniproject3.jfrog.io"

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

        stage('SonarQube analysis') {
            environment {
                scannerHome = tool 'miniproject-sonar-scanner'
            }

            steps {
                withSonarQubeEnv('miniproject-sonarqube-server') {
                    sh "${scannerHome}/bin/sonar-scanner"
                }
            }
        }
        
        stage('Push to Artifactory') {
            steps {
                script {
                    def server = Artifactory.server url: registry+"/artifactory", credentialsId: "artifact-cred"
                    def buildInfo = Artifactory.newBuildInfo()

                    def filePath = "/home/ubuntu/jenkins/workspace/test2_main/gohtmx"
                    def artifactLocation = "gohtmx"
                    def repositoryPath = "miniproject-go-local"

                    server.upload spec: {
                        '/home/ubuntu/jenkins/workspace/test2_main/gohtmx': repositoryPath + artifactLocation
                    }, buildInfo: buildInfo, failNoOp: true, recursive: true, flat: false

                    server.publishBuildInfo buildInfo
                }
            }
        }
    }
}
