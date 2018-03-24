node('jnlp-slave') {

    stage ('checkout') {
        git credentialsId: 'nevis', url: 'http://gitlab.dpool.sina.com.cn/nevis.io/test.git'
        sh "pwd"
        sh "hostname"
        sh "ls /home/jenkins"
        sh "ifconfig"
        echo "make build"
    }

    stage ('自定义名称-代码构建') {
        // 代码编译，开发语言：PHP、Python、Golang、C、Java
        echo "make build"

        // 单元测试
        echo "make ut"

        // 功能测试
        echo "make ft"

        // 其他
        echo "other script"

    }

    stage ('自定义名称-构建镜像') {
        // 镜像构建
        // URL: 镜像仓库地址，registry.dpool.sina.com.cn/nevis.io
        // NAME：镜像名称
        // TAG：镜像版本，包括自定义、CommitID、时间戳
        // DOCKERFILE_PATH: Dockerfile路径，默认值为/Dockerfile
        echo "docker build -t URL/NAME:TAG . -f ./DOCKERFILE_PATH" 

    }

    stage ('自定义名称-灰度发布') {
        // 执行部署脚本
        // CLUSTER: 所属集群
        // NAMESPACE：组织
        // SERVICE: 服务
        echo "deploy CLUSTER NAMESPACE SERVICE" 

    }

    input message: "是否继续?"

    stage ('自定义名称-全量发布') {
        // 执行部署脚本
        // NAMESPACE：组织
        // SERVICE: 服务
        echo "deploy all NAMESPACE SERVICE" 

    }

    echo "Pipeline has been done successfully."
}
