# Go-FastText

Golang bindings to the fasttext library.

## Usage

To perform a prediction on a model, use the following command

```
go run main.go prediction -m [model_path] [query]
```

For example

```
go run main.go predict -m ~/Downloads/ag_news.bin chicken
```
