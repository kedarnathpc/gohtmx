def registry = 'https://miniproject.jfrog.io'
def imageName = 'miniproject.jfrog.io/miniproject-docker-local/gohtmx'
def version = '1.0.0'

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
            // environment {
            //     scannerHome = tool 'miniproject-sonar-scanner'
            // }

            // steps {
            //     withSonarQubeEnv('miniproject-sonarqube-server') {
            //         sh "${scannerHome}/bin/sonar-scanner"
            //     }
            // }
            steps {
                sh 'ls'
            }
        }
        

        stage("Publish to Artifactory") {
            steps {
                script {
                    echo '<--------------- GoLang Publish Started --------------->'

                    def server = Artifactory.newServer url: registry + "/artifactory", credentialsId: "artifact-cred"
                    def filePath = "/home/ubuntu/jenkins/workspace/movie-app/gohtmx"
                    def artifactLocation = "gohtmx"
                    def repositoryPath = "miniproject-go-local/"

                    def properties = "buildid=${env.BUILD_ID},commitid=${GIT_COMMIT}"
                    def uploadSpec = """{
                        "files": [
                            {
                                "pattern": "${filePath}",
                                "target": "${repositoryPath}/${artifactLocation}",
                                "props": "${properties}"
                            }
                        ]
                    }"""

                    def buildInfo = server.upload(uploadSpec)
                    server.publishBuildInfo(buildInfo)

                    echo '<--------------- GoLang Publish Ended --------------->'
                }
            }
        }

        stage('Give Docker Permissions'){
            steps {
                script {
                    echo '<---------Giving Docker Permissions--------->'
                    sh 'sudo chmod 0777 /var/run/docker.sock'
                    echo '<---------Docker Permissions Given--------->'
                }
            }
        }

        stage ('Remove existing containers') {
            steps {
                script {
                    echo '<---------Removing existing containers--------->'
                    sh 'sudo docker system prune -y'
                    echo '<---------Existing containers removed--------->'
                }
            }
        }

        stage('Docker Build') {
            steps {
                script {
                    echo '<---------Building Docker Image--------->'
                    app = docker.build(imageName + ':' + version)
                    echo '<---------Docker Image Built--------->'
                }
            }
        }

        stage('Docker Publish'){
            steps {
                script {
                    echo '<---------Publishing Docker Image--------->'
                    docker.withRegistry(registry, 'artifact-cred') {
                        app.push()
                    }
                    echo '<---------Docker Image Published--------->'
                }
            }
        }

        stage('Deploy Docker image') {
            steps {
                script {
                    echo '<---------Deploying Docker Image--------->'
                    sh 'docker run -p 8080:8080 -d miniproject.jfrog.io/miniproject-docker-local/gohtmx:1.0.0'
                    echo '<---------Docker Image Deployed--------->'
                
                }
            }
        }

        // stage ('Delete Previous Deployment') {
        //     steps {
        //         script {
        //             echo '<---------Deleting previous deployment--------->'
        //             sh './delete-deploy.sh'
        //             echo '<---------Previous deployment deleted--------->'
        //         }
        //     }
        // }

        // stage ('New Deploy') {
        //     steps {
        //         script {
        //             echo '<---------Deploying application--------->'
        //             sh './deploy.sh'
        //             echo '<---------Application deployed--------->'
        //         }
        //     }
        // }
    }
}
