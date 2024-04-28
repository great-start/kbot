pipeline {
    agent any
    environment {
        REPO = "https://github.com/great-start/kbot.git"
        BRANCH = "main"
    }
    parameters {

        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        
        choice(name: 'arch', choices: ['amd64', 'aarch64', 'ARMv7', 'i386'], description: 'Pick arch')

    }
    stages {
        stage('clone') {
            steps {
                echo "Clone repo"
                git branch: "${BRANCH}", url: "${REPO}" 
            }
        }

        stage('test') {
            steps {
                echo "Test"
                bash "make test"
            }
        }
        
        stage('build') {
            steps {
                echo "Build for platform ${params.OS}"
                echo "Build for arch: ${params.ARCH}"
                bash "make TARGETOS=${params.OS} TARGETARCH=${params.ARCH} build"
            }
        }
    }
}