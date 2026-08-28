# Official Go image with the complete toolchain required by the evaluator.
FROM golang:{{GO_VERSION}}

WORKDIR /app

{{GO_DEPENDENCIES}}

# Keep the complete project, including any project-owned Dockerfile and BENZHI_README.md.
COPY . .

{{BUILD_STEPS}}

CMD ["bash"]
