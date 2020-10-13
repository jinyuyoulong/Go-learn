[TOC]

## 为什么要了解Bazel

这段时间搞到了bilibili的一部分源码，那个事件你知道的，就不说了。怀着对大牛的向往，打算研究一下，之前听说主程做的这个架构很牛逼，就算是代码泄露也不会对服务器造成影响，而且模块化做的非常好，小弟们码好代码，主程一键做个合并打包发布就行。自动化做的很好。就想知道到底是怎么做的，于是我了解到了他——Bazel。

原文链接 https://filipnikolovski.com/managing-go-monorepo-with-bazel/

# 使用Bazel管理Go monorepo

在[InPlayer中](https://inplayer.com/)，我们有一个使用*微服务*架构风格构建的平台，它基本上将应用程序构建为许多不同服务的集合。在这篇文章中，我将讨论如何建构(structure)，构建(build)和部署(deploy)Go应用程序。

我们编写的每一段Go代码都驻留在一个Git存储库中 - 一个monorepo。由于每个库和服务都在一个项目中，因此它允许我们进行交叉更改，而无需使用某些外部包管理工具。基本上，代码不可能不同步，我们所做的每个更改都可以视为一个单元。

虽然好处很明显，但使用Go monorepo的挑战是如何有效地构建和测试每个包。答案 - [Bazel](https://bazel.build/)。

### Bazel是什么？

Bazel是一款速度极快的构建工具。它只重建必要的东西，它利用高级缓存机制和并行执行，使您的构建非常，非常快。除了这些功能外，它还可以管理您的代码依赖项，调用外部工具和插件，还可以从二进制可执行文件构建Docker镜像。它使用`go build`引擎盖，但它也可以支持许多不同的语言，而不仅仅是Go。您可以将它用于Java，C ++，Android，iOS和各种其他语言平台。您可以在三个主要操作系统上运行Bazel - Windows，macOS和Linux。

### 项目结构

在我们深入了解Bazel之前，首先让我们讨论一下我们的项目结构：

```
    platform
    |-- src
    |    |-- foo
    |    |   |--cmd
    |    |   |  `--bar
    |    |   |     |--BUILD
    |    |   |     `--main.go
    |    |   `--pkg
    |    |-- utils
    |    |-- vendor
    |    |-- Gopkg.lock
    |    |-- Gopkg.toml
    |    |-- BUILD
    |    `-- WORKSPACE
    |-- README.md
    `-- gitlab-ci.yml
```

该`platform`目录是我们的根本，一切从这里开始。在该文件夹中，我们有CI配置和`src`保存所有代码的目录。每个服务都是`src`文件夹中的子目录，在每个服务中我们都有两个顶级目录，即`cmd`和`pkg`文件夹。在下面`cmd`我们有我们的二进制文件（我们的主程序）的`pkg`目录，该目录用于我们的服务库。

Bazel从名为*workspace*的目录中组织的代码构建软件，该目录基本上是我们的src目录。在这里，我们的工作空间目录必须包含一个名为的文件`WORKSPACE`，该文件可能引用了构建输出所需的外部依赖关系以及构建规则。

这是一个示例WORKSPACE文件：

```
http_archive(
    name = "io_bazel_rules_go",
    url = "https://github.com/bazelbuild/rules_go/releases/download/0.9.0/rules_go-0.9.0.tar.gz",
    sha256 = "4d8d6244320dd751590f9100cf39fd7a4b75cd901e1f3ffdfd6f048328883695",
)
http_archive(
    name = "bazel_gazelle",
    url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.9/bazel-gazelle-0.9.tar.gz",
    sha256 = "0103991d994db55b3b5d7b06336f8ae355739635e0c2379dea16b8213ea5a223",
)
git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    tag = "v0.3.0",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()
```

在此文件中，有几个依赖项添加到工作区。我们特别声明我们将使用[rules_go](https://github.com/bazelbuild/rules_go)和[rules_docker](https://github.com/bazelbuild/rules_docker)依赖项以及[Gazelle](https://github.com/bazelbuild/bazel-gazelle)，这将帮助我们生成Bazel所需的一些文件。不要担心，如果您不熟悉这种语法，需要一些时间来适应它。

### BUILD文件

Bazel有一个关于*包*的概念，它被定义为相关文件的集合以及它们之间依赖关系的规范。如果Bazel工作空间内的目录包含名为的文件`BUILD`，则会将该目录视为包。包中包含其目录中的所有文件，以及其下的所有子目录，除了那些本身包含BUILD文件的文件。

BUILD文件包含构建规则，这些规则定义了我们应该如何构建包。您可以[在此处](https://docs.bazel.build/versions/master/build-ref.html)阅读有关概念和术语的更多信息。

在开始一个新项目时，我们需要做的第一件事是在根目录中添加一个BUILD文件，这将加载稍后用于运行[Gazelle](https://github.com/bazelbuild/bazel-gazelle) with Bazel 的瞪羚规则。

```
package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_docker//container:container.bzl")
load("@io_bazel_rules_go//go:def.bzl", "go_prefix", "gazelle")

go_prefix("github.com/example/project")
gazelle(
  prefix = "github.com/example/project/src",
  name = "gazelle",
  command = "fix",
  external = "vendored"
)
```

添加此文件后，我们可以使用以下命令运行Gazelle：

```
bazel run //:gazelle
```

这将根据项目中的go文件生成新的BUILD文件。稍后添加新程序和库时，应使用相同的命令更新现有的BUILD文件，否则构建可能会失败。

作为一个例子（基于我们之前显示的项目结构），gazelle将为我们的`bar`程序生成一个BUILD文件，该文件位于`foo`包中，如下所示：

```
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/example/project/src/foo/cmd/bar",
    visibility = ["//visibility:private"],
    deps = [
        #Any dependencies that our library has will be loaded here
    ],
)

go_binary(
    name = "bar",
    embed = [":go_default_library"],
    importpath = "github.com/example/project/src/foo/cmd/bar",
    visibility = ["//visibility:public"],
)
```

现在通过运行命令`bazel build //foo/...`bazel将构建我们的Go程序并将二进制文件保存在输出目录中。如果要构建整个项目，只需`bazel build //...`在根文件夹中运行即可。

如果您为您的库和程序（您应该）编写测试，gazelle将为`go_test`它们生成规则，然后您可以运行`bazel test //...`将运行所有测试。

Bazel的高级缓存，使运行`build`和`test`命令对整个工作区超级快，因为它只会建造或测试您已更改的文件，以及依赖于这些修改过的文件的文件。

⚠️注意：确保将CI服务器设置为缓存输出目录，否则运行bazel不会带来太多好处。

### Docker镜像

在我们想要将二进制文件构建和部署为docker图像的情况下，bazel有一套很好的规则可以做到这一点。更为重要的是，Bazel**并不**需要Docker拉取，构建或推送镜像。这意味着您可以使用这些规则在Windows / OSX上构建Docker镜像而无需使用`docker-machine`或者`boot2docker`也不需要在笔记本电脑上进行*root*访问。

我们`bar`程序的BUILD文件的完整示例如下所示：

```
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

# First we load the go_image and container_push rules
load("@io_bazel_rules_docker//go:image.bzl", "go_image")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/example/project/src/foo/cmd/bar",
    visibility = ["//visibility:private"],
    deps = [
        #Any dependencies that our library has will be loaded here
    ],
)

go_binary(
    name = "bar",
    embed = [":go_default_library"],
    importpath = "github.com/example/project/src/foo/cmd/bar",
    visibility = ["//visibility:public"],
)

go_image(
    name = "docker",
    binary = ":bar",
)
```

该`go_image`规则使用[distroless](https://github.com/GoogleCloudPlatform/distroless)镜像作为基础，只添加二进制文件作为要运行的命令。`container_push`如果要将映像推送到远程存储库，也可以使用该规则。

要将二进制文件作为docker镜像运行，只需键入`bazel run //foo/cmd/bar:docker`命令即可。您还可以构建一个tar包，然后可以使用以下命令手动将其加载到docker中：

- `bazel build //foo/cmd/bar:docker.tar`
- `docker load -i bazel-output/foo/cmd/bar/docker.tar`

您可以在[此处](https://github.com/bazelbuild/rules_docker)找到有关规则的更多信息。