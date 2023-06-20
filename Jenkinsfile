pipeline {
    agent { 
        kubernetes {
               
            yaml '''
apiVersion: v1
kind: Pod
metadata:
  name: kaniko-test
  namespace: jenkins
spec:
  containers:
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug
    ports:
    - containerPort: 50000
    command: 
    - sleep
    args: 
    - infinity
    volumeMounts:
      - name: kaniko-secret
        mountPath: /kaniko/.docker
  volumes:
    - name: kaniko-secret
      secret:
        secretName: kaniko-test
        items:
          - key: .dockerconfigjson
            path: config.json
'''
        }

    }
    stages {
        stage('Main') {
            steps {
                container(name: 'kaniko') {
                sh 'executor --dockerfile=go-test/Dockerfile \
                    --context=git://github.com/coldpaper2/nks-cicd.git \
                    --destination=mhkim1560/jenkins-test:test${BUILD_NUMBER}'
            }
         }   
        }
    }
}
