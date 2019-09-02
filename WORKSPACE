load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_jar")

http_archive(
        name = "io_bazel_rules_go",
        urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.18.3/rules_go-0.18.3.tar.gz"],
        sha256 = "86ae934bd4c43b99893fc64be9d9fc684b81461581df7ea8fc291c816f5ee8c5",
        )

http_archive(
        name = "bazel_gazelle",
        urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
        sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
        )

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")


go_repository(
        name = "com_github_antlr_antlr4",
        importpath = "github.com/antlr/antlr4",
        commit = "38a95da3976344019d2cdf2d4c106966166a781d",
        )

go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        commit = "6f77996f0c42f7b84e5a2b252227263f93432e9b",
        )

gazelle_dependencies()

http_jar(
        name = "antlr4-complete",
        url = "https://www.antlr.org/download/antlr-4.7.2-complete.jar",
        sha256 = "6852386d7975eff29171dae002cc223251510d35f291ae277948f381a7b380b4",
        )
