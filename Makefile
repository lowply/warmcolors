clean:
	bundle exec jekyll clean

build: clean
	bundle exec jekyll build

watch: clean
	bundle exec jekyll server --watch --config _config.yml
