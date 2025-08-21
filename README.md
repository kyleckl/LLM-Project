# [Local LLM Project]

This is a program that is designed to run locally, and hit an LLM either locally through a service like Ollama, or another online provider.

# TODO:
- Implement a Makefile that is able to build, run, and lint the program.
- Create an API client struct that the program can use to communicate with LLM APIs.
- Implement tests using Ginkgo/Gomega.
- Dockerfile/Kubernetes files to deploy in a cluster?
- Add lefthooks to automate pre-commit and pre-push