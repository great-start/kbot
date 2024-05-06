pipeline {
    agent any
    environment {
        GITHUB_TOKEN=credentials('great-start')
        REPO = "https://github.com/great-start/kbot.git"
        BRANCH = "develop"
    }

    parameters {
        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'ios', 'android'], description: 'Pick OS')
        choice(name: 'arch', choices: ['amd64', 'arm', '386'], description: 'Pick arch')

    }

    stages {
        stage('clone') {
            steps {
                echo "Clone repo"
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }

        stage('format') {
            steps {
                echo "Test"
                sh "make test"
            }
        }

        stage('image') {
            steps {
                echo "Building image for platform ${params.OS} on ${params.ARCH} started"
                sh "make image TARGETOS=${params.OS} TARGETARCH=${params.ARCH}"
            }
        }

        stage('login to ghcr.io') {
            steps {
                sh "echo $GITHUB_TOKEN_PSW | docker login ghcr.io -u $GITHUB_TOKEN_USR --password-stdin"
            }
        }

        stage('push') {
            steps {
                echo "Push image"
                sh "make push TARGETOS=${params.OS} TARGETARCH=${params.ARCH}"
            }
        }
    }
}