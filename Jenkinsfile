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
                    def server = Artifactory.server url: "https://miniproject3.jfrog.io/artifactory", credentialsId: "artifact-cred"
                    def buildInfo = Artifactory.newBuildInfo()

                    def filePath = "/home/ubuntu/jenkins/workspace/test2_main/gohtmx"
                    def artifactLocation = "gohtmx"
                    def repositoryPath = "miniproject-go-local/"

                    server.upload spec: [
                        // Specifying the file path(s) and destination directory in Artifactory
                        // The syntax is { source: destination }
                        // For multiple files, you can specify them individually or use a wildcard pattern
                        // Example: '/path/to/source/file.txt': 'repo-name/path/in/artifactory/file.txt'
                        // Wildcard example: '/path/to/source/*.txt': 'repo-name/path/in/artifactory/'
                        (filePath): repositoryPath + artifactLocation
                    ], buildInfo: buildInfo, failNoOp: true, recursive: true, flat: false

                    server.publishBuildInfo buildInfo
                }
            }
        }
    }
}
