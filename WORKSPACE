load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_jar")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8663604808d2738dc615a2c3eb70eba54a9a982089dd09f6ffe5d0e75771bc4f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.6/rules_go-v0.23.6.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.6/rules_go-v0.23.6.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_antlr_antlr4",
    commit = "38a95da3976344019d2cdf2d4c106966166a781d",
    importpath = "github.com/antlr/antlr4",
)

go_repository(
    name = "com_github_google_go_cmp",
    commit = "6f77996f0c42f7b84e5a2b252227263f93432e9b",
    importpath = "github.com/google/go-cmp",
)

gazelle_dependencies()

http_jar(
    name = "antlr4-complete",
    sha256 = "6852386d7975eff29171dae002cc223251510d35f291ae277948f381a7b380b4",
    url = "https://www.antlr.org/download/antlr-4.7.2-complete.jar",
)
